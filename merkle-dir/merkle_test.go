package merkle_dir

import "testing"
import (
	"github.com/uepoch/distributed-experiments/lib/utils"
	"crypto"
	"github.com/stretchr/testify/assert"
)

func benchmarkUpdate(hash crypto.Hash, b *testing.B) {
	m, _ := NewMerkleRoot(hash)
	strGen, errCh := utils.StringGenerator(5, 50)
	defer close(errCh)
	for n := 0; n < b.N; n++ {
		m.Update(<-strGen, HashableInt(n))
	}
}

func TestUpdate(t *testing.T) {
	node, err := Update("a/b/c/", HashableInt(0))

	assert.Error(t, err, "There should be an error")
	assert.Nil(t, node, "This should be nul")

	node, err = Update("a/b/c", HashableInt(0))

	assert.NoError(t, err, "This should not be an error")
	assert.NotNil(t, node, "This should not be nil")
	assert.Equal(t, node.Path, "c", "This should be 'c'")
	parent, err := Get("a/b/")
	assert.NoError(t, err, "This should return the parent node and no error")
	assert.Equal(t, parent, node.Parent,"This should return the parent node")
}

func BenchmarkUpdateSHA512(b *testing.B) {
	benchmarkUpdate(crypto.SHA512, b)
}
func BenchmarkUpdateMD5(b *testing.B) {
	benchmarkUpdate(crypto.MD5, b)
}
func BenchmarkUpdateSHA256(b *testing.B) {
	benchmarkUpdate(crypto.SHA256, b)
}
