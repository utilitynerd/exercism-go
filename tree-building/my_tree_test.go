package tree

import "testing"

var n = &Node{
	ID: 0,
	Children: []*Node{
		{
			ID: 1,
			Children: []*Node{
				{ID: 4},
				{ID: 5},
			},
		},
		{
			ID: 2,
			Children: []*Node{
				{ID: 3},
				{ID: 6},
			},
		},
	},
}

func TestGetNOde(t *testing.T) {
	for i := 0; i <= 6; i++ {
		node := n.getNode(i)
		if node.ID != i {
			t.Errorf("expected %v got %v", i, node.ID)
		}
	}

	node := n.getNode(2)
	if node.ID != 2 {
		t.Errorf("Expected %v got %v", 2, node.ID)
	}
}
