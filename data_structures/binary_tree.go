package gogenerics

import "fmt"

type BinaryTree[T any] struct {
	LeftNode  *BinaryTree[T]
	Value     T
	RightNode *BinaryTree[T]
}

func (t BinaryTree[T]) HasAnyChildren() bool {
	return t.LeftNode != nil && t.RightNode != nil
}

func (t BinaryTree[T]) PreorderTraversalHandle(root *BinaryTree[T], handler func(t *BinaryTree[T])) {
	if root == nil {
		return
	}

	handler(root)
	t.PreorderTraversalHandle(root.LeftNode, handler)
	t.PreorderTraversalHandle(root.RightNode, handler)
}

func (t BinaryTree[T]) InorderTraversalHandle(root *BinaryTree[T], handler func(t *BinaryTree[T])) {
	if root == nil {
		return
	}

	t.InorderTraversalHandle(root.LeftNode, handler)
	handler(root)
	t.InorderTraversalHandle(root.RightNode, handler)
}

func (t BinaryTree[T]) PostorderTraversalHandle(root *BinaryTree[T], handler func(t *BinaryTree[T])) {
	if root == nil {
		return
	}

	t.PostorderTraversalHandle(root.LeftNode, handler)
	t.PostorderTraversalHandle(root.RightNode, handler)
	handler(root)
}

func (t *BinaryTree[T]) Print() {
	t.PreorderTraversalHandle(t, func(t *BinaryTree[T]) {
		fmt.Println(t.Value)
	})
}

func (t *BinaryTree[T]) InsertLeftNode(treeNode *BinaryTree[T], value T) {
	if treeNode == nil {
		return
	}

	t.InorderTraversalHandle(t, func(node *BinaryTree[T]) {
		if node == treeNode {
			node.LeftNode = &BinaryTree[T]{
				Value:     value,
				LeftNode:  nil,
				RightNode: nil,
			}
		}
	})
}

func (t *BinaryTree[T]) InsertRightNode(treeNode *BinaryTree[T], value T) {
	if treeNode == nil {
		return
	}

	t.InorderTraversalHandle(t, func(node *BinaryTree[T]) {
		if node == treeNode {
			node.RightNode = &BinaryTree[T]{
				Value:     value,
				LeftNode:  nil,
				RightNode: nil,
			}
		}
	})
}
