package merkle_dir

import "fmt"

type Action uint8

const (
	ActionCreate Action = iota
	ActionUpdate
	ActionDelete
)

//go:generate stringer -type=Action

type Changes []Change

type Change struct {
	FullPath string
	Action   Action
}

func (c Changes) Apply(mapFunc map[Action]func(*Change) (error)) (error) {
	for _, change := range c {
		if _, ok := mapFunc[change.Action]; !ok {
			return fmt.Errorf("Action %s is not registered for change on %s", change.Action, change.FullPath)
		}
	}
	return nil
}

func NewChange(fullpath string, action Action) Change {
	return Change{FullPath: fullpath, Action: action}
}

func (m *MerklePath) toChange(action Action) Change {
	return Change{FullPath: m.Name(), Action: action}
}
