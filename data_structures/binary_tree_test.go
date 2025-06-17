package gogenerics

import "testing"

func TestBinaryHasAnyChildren(t *testing.T) {
	tree := &BinaryTree[int]{
		LeftNode:  nil,
		RightNode: nil,
		Value:     1,
	}

	if tree.HasAnyChildren() {
		t.Errorf("tree.HasAnyChildren()")
		t.FailNow()
	}

	tree = &BinaryTree[int]{
		Value: 1,
		LeftNode: &BinaryTree[int]{
			Value: 0,
		},
		RightNode: &BinaryTree[int]{
			Value: 2,
		},
	}

	if !tree.HasAnyChildren() {
		t.Errorf("!tree.HasAnyChildren()")
		t.FailNow()
	}
}

func TestBinaryTreeInsertLeftNode(t *testing.T) {
	tree := &BinaryTree[int]{
		LeftNode:  nil,
		RightNode: nil,
		Value:     1,
	}

	// tree.InsertInNode(&tree, 1)
	tree.InsertLeftNode(tree, 0)
	tree.InsertLeftNode(tree.LeftNode, -1)

	values := []int{}

	tree.PreorderTraversalHandle(tree, func(tr *BinaryTree[int]) {
		values = append(values, tr.Value)
	})

	expectedValues := []int{1, 0, -1}

	if len(values) != len(expectedValues) {
		t.Errorf("len(values) != len(expectedValues)")
		t.FailNow()
	}

	for i, v := range values {
		if v != expectedValues[i] {
			t.Errorf("v != expectedValues[i]")
			t.FailNow()
		}
	}
}

func TestBinaryTreeInsertRightNode(t *testing.T) {
	tree := &BinaryTree[int]{
		LeftNode:  nil,
		RightNode: nil,
		Value:     1,
	}

	// tree.InsertInNode(&tree, 1)
	tree.InsertRightNode(tree, 2)
	tree.InsertRightNode(tree.RightNode, 3)

	values := []int{}

	tree.InorderTraversalHandle(tree, func(tr *BinaryTree[int]) {
		values = append(values, tr.Value)
	})

	expectedValues := []int{1, 2, 3}

	if len(values) != len(expectedValues) {
		t.Errorf("len(values) != len(expectedValues)")
		t.FailNow()
	}

	for i, v := range values {
		if v != expectedValues[i] {
			t.Errorf("v != expectedValues[i]")
			t.FailNow()
		}
	}
}
