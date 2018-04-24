package merkle_dir

import (
	"crypto"
	_ "crypto/sha512"
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/uepoch/distributed-experiments/lib/utils"
	"go.uber.org/zap"
	"bytes"
)

type Hashable interface {
	Hash(hashType crypto.Hash) []byte
}

type MerklePath struct {
	sync.Mutex
	Path   string
	Sum    []byte
	Hasher crypto.Hash
	Parent *MerklePath
	Tree   map[string]*MerklePath
	Leaf   bool
}

var globalMerkle = MerklePath{Hasher: crypto.SHA512, Tree: map[string]*MerklePath{}}

func SanitizePath(Path string) string {
	if Path == "/" {
		return Path
	}
	ret := path.Clean(strings.TrimPrefix(Path, "/"))
	if strings.HasSuffix(Path, "/") {
		ret += "/"
	}
	return ret
}

func NewMerkleRoot(hashType crypto.Hash) (*MerklePath, error) {
	if !hashType.Available() {
		return nil, fmt.Errorf("Hash type provided is not supported in Go")
	}
	return &MerklePath{
		Path:   "/",
		Hasher: hashType,
		Parent: nil,
		Tree:   map[string]*MerklePath{},
	}, nil
}

func NewMerklePath(path string, hashType crypto.Hash, parent *MerklePath, leaf bool) *MerklePath {
	var tree map[string]*MerklePath

	if !leaf {
		tree = map[string]*MerklePath{}
	}

	return &MerklePath{
		Path:   path,
		Hasher: hashType,
		Parent: parent,
		Leaf:   leaf,
		Tree:   tree,
	}
}

func Get(Path string) (*MerklePath, error) {
	return globalMerkle.Get(Path)
}

func (m *MerklePath) Get(Path string) (*MerklePath, error) {
	Path = SanitizePath(Path)
	if Path == "/" {
		return m, nil
	}
	ret := m.get(Path)
	if ret == nil {
		return ret, fmt.Errorf("node was not found for path: %s", Path)
	}
	return ret, nil
}

func (m *MerklePath) get(Path string) *MerklePath {
	for subPath, nodePtr := range m.Tree {

		if subPath == Path {
			return nodePtr
		} else if strings.HasPrefix(Path, subPath) && strings.HasSuffix(subPath, "/") {
			return nodePtr.get(strings.TrimPrefix(Path, subPath))
		}
	}
	return nil
}

func (m *MerklePath) hashUpdate() error {
	if !m.Leaf {
		h := m.Hasher.New()
		for _, v := range m.Tree {
			h.Write([]byte(v.Path))
			h.Write(v.Sum)
		}
		m.Sum = h.Sum(nil)
	}
	if m.Parent != nil {
		return m.Parent.hashUpdate()
	}
	return nil
}

func Update(AbsPath string, hashable Hashable) (*MerklePath, error) {
	return globalMerkle.Update(AbsPath, hashable)
}

func (m *MerklePath) Update(AbsPath string, hashable Hashable) (*MerklePath, error) {
	defer utils.TimeTrack(zap.L().With(zap.String("path", AbsPath)), time.Now(), "Update finished")

	if strings.HasSuffix(AbsPath, "/") {
		return nil, fmt.Errorf("path %s has a trailing space, you can't update a directory", AbsPath)
	}
	subPath := SanitizePath(AbsPath)
	// In some KV, foo/bar and foo/bar/ are not the same

	return m.updateIter(subPath, hashable), nil
}

func (m *MerklePath) updateIter(Path string, hashable Hashable) *MerklePath {
	currNode := m
	//var prevNode *MerklePath
	prevNode := currNode
	var relPath string
	for leafFound := false; !leafFound; prevNode = currNode {
		currNode.Lock()
		defer currNode.Unlock()
		nextSlash := strings.IndexByte(Path, '/')
		if nextSlash > -1 {
			relPath = Path[:nextSlash+1]
		} else {
			relPath = Path
			leafFound = true
		}
		//zap.S().With("relPath", relPath, "fullPath", currNode.Name()).Info("Entering loop")
		if node, ok := currNode.Tree[relPath]; ok {
			currNode = node
		} else {
			currNode = NewMerklePath(relPath, m.Hasher, prevNode, leafFound)
			prevNode.Tree[relPath] = currNode
		}
		Path = Path[nextSlash+1:]
	}
	currNode.Sum = hashable.Hash(m.Hasher)
	if err := currNode.hashUpdate(); err != nil {
		panic(fmt.Errorf("error when updating the tree hash: %s", err))
	}
	return currNode
}

func (m *MerklePath) updateRecursive(path string, hashable Hashable) *MerklePath {

	m.Lock()
	defer m.Unlock()
	var subNode *MerklePath

	splited := strings.SplitN(path, "/", 3)
	nodePath := path
	leafNode := true

	if l := len(splited); l > 1 || (l == 1 && strings.HasSuffix(path, "/")) {
		// This is a directory and we need to dig in
		nodePath = splited[0] + "/"
		leafNode = false
	}
	if node, ok := m.Tree[nodePath]; ok {
		subNode = node
	} else {
		subNode = NewMerklePath(nodePath, m.Hasher, m, leafNode)
		m.Tree[nodePath] = subNode
	}

	if !leafNode {
		return subNode.updateRecursive(strings.TrimPrefix(path, nodePath), hashable)
	}
	// End of recursion
	subNode.Sum = hashable.Hash(m.Hasher)
	if err := subNode.hashUpdate(); err != nil {
		panic(fmt.Errorf("Fck it: %s", err))
	}
	return subNode

}

func (m *MerklePath) Name() string {
	ret := ""

	for m != nil {
		ret = m.Path + ret
		m = m.Parent
	}
	return ret
}
func Compare(other *MerklePath) (Changes, error) {
	return globalMerkle.Compare(other)
}


func (m *MerklePath) Compare(other *MerklePath) (changes Changes, error error) {
	defer utils.TimeTrack(zap.L(), time.Now(), "Compare finished")
	if m.Hasher != other.Hasher {
		return nil, fmt.Errorf("can't compare tree with different hashers")
	}
	return m.compare(other)
}

func (m *MerklePath) compare(other *MerklePath) (Changes, error) {
	m.Lock()
	defer m.Unlock()
	other.Lock()
	defer other.Unlock()
	changes := make(Changes, 0, len(other.Tree))
	if !bytes.Equal(m.Sum, other.Sum) {
		if !m.Leaf {
			for k, v := range m.Tree {
				if node, ok := other.Tree[k]; !ok {
					//Not present in the other Tree
					changes = append(changes, v.toChange(ActionDelete))
				} else {
					// Present
					subChanges, err := v.compare(node)
					if err != nil {
						return nil, err
					}
					changes = append(changes, subChanges...)
				}
			}
			for k, v := range other.Tree {
				if _, ok := m.Tree[k]; !ok {
					changes = append(changes, v.toChange(ActionCreate))
				}
			}
		} else {
			changes = append(changes, m.toChange(ActionUpdate))
		}
	}
	return changes, nil

}
