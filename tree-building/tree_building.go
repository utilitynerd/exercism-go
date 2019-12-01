package tree

import (
	"errors"
	"sort"
)

// Record is one row from the forums data store
type Record struct {
	ID     int
	Parent int
}

// Node represents a node in a tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree of Nodes from a slice of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// Sorting the input records ensures a parent record
	// is handled before any of its children
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	root := records[0]
	if root.ID != 0 || root.Parent != 0 {
		return nil, errors.New("invalid root node")
	}
	nodes := make(map[int]*Node)
	tree := &Node{ID: 0}
	nodes[0] = tree
	for i, r := range records[1:] {
		if i+1 != r.ID {
			return nil, errors.New("non continuous")
		}
		if r.Parent >= r.ID {
			return nil, errors.New("invalid parent")
		}
		n := &Node{ID: r.ID, Children: nil}
		nodes[r.ID] = n

		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, n)
	}
	return tree, nil
}
