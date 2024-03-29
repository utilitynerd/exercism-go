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
	// Sorting the input records ensures a parent record
	// is handled before any of its children
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make(map[int]*Node)
	for i, r := range records {
		if (r.ID == 0 && r.Parent != 0) || i != r.ID || (r.ID > 0 && r.Parent >= r.ID) {
			return nil, errors.New("non sequential or invalid parent")
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID > 0 {
			p := nodes[r.Parent]
			p.Children = append(p.Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
