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

type records []Record

// implement sort.Interface for records
func (r records) Len() int           { return len(r) }
func (r records) Less(i, j int) bool { return r[i].ID < r[j].ID }
func (r records) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Node represents a node in a tree
type Node struct {
	ID       int
	Children []*Node
}

func (n *Node) getNode(id int) *Node {
	if n.ID == id {
		return n
	}
	for _, c := range n.Children {
		node := c.getNode(id)
		if node != nil {
			return node
		}
	}
	return nil
}

func (n *Node) addChild(node Node) {
	n.Children = append(n.Children, &node)
}

// Build builds a tree of Nodes from a slice of records
func Build(recs records) (*Node, error) {
	// Sorting the input records ensures a parent record
	// is handled before any of its children
	sort.Sort(recs)
	if len(recs) == 0 {
		return nil, nil
	}
	root := recs[0]
	if root.ID != 0 || root.Parent != 0 {
		return nil, errors.New("invalid root node")
	}
	tree := &Node{ID: 0}
	for i, r := range recs[1:] {
		if i+1 != r.ID {
			return nil, errors.New("non continous")
		}
		if r.Parent >= r.ID {
			return nil, errors.New("invalid parent")
		}
		parent := tree.getNode(r.Parent)
		parent.addChild(Node{ID: r.ID, Children: nil})
	}
	return tree, nil
}
