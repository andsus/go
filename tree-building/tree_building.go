package tree

import (
	"fmt"
	"sort"
)

//Record is a database item that specifies a parent record in a tree structure.
type Record struct {
	ID, Parent int
}

//String converts a record to a string.
func (r *Record) String() string {
	return fmt.Sprintf("(ID: %d, Parent: %d)", r.ID, r.Parent)
}

//Node is a node in a tree structure.
type Node struct {
	ID       int
	Children []*Node
}

// Build converts a set of records to a tree of Nodes.
func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	nodes := make([]*Node, len(records))
	for i, r := range records {
		nodes[i] = &Node{ID: r.ID}
		if i != r.ID || r.Parent > r.ID || r.ID > 0 && r.ID == r.Parent {
			return nil, fmt.Errorf("Not in sequence or has a bad parent: %v", r)
		}

		if r.ID > 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[i])
		}
	}
	return nodes[0], nil
}
