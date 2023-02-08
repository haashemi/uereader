package uereader

import "fmt"

type Error struct {
	err     error
	Name    string
	IndexID int32
}

func (e Error) Error() string {
	return fmt.Sprintf("category %s index %d: %v", e.Name, e.IndexID, e.err)
}
