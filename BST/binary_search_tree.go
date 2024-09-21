package main

import (
	"fmt"
)

// Node represents a node in the binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new node with the given value to the BST
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrder prints the in-order traversal of the BST (Left, Root, Right)
func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrder()
}

// PreOrder prints the pre-order traversal of the BST (Root, Left, Right)
func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

// PostOrder prints the post-order traversal of the BST (Left, Right, Root)
func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Printf("%d ", n.Value)
}

// Search returns true if the value exists in the BST
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if n.Value == value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// FindMin finds the node with the minimum value in the BST
func (n *Node) FindMin() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.FindMin()
}

// FindMax finds the node with the maximum value in the BST
func (n *Node) FindMax() *Node {
	if n.Right == nil {
		return n
	}
	return n.Right.FindMax()
}

// Delete removes a node with the specified value from the BST
func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		// Node to delete found
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// Node with two children: Get the inorder successor (smallest in the right subtree)
		minRight := n.Right.FindMin()
		n.Value = minRight.Value
		n.Right = n.Right.Delete(minRight.Value)
	}
	return n
}

// Main function demonstrating the usage of the BST
func main() {
	root := &Node{Value: 50}
	root.Insert(30)
	root.Insert(70)
	root.Insert(20)
	root.Insert(40)
	root.Insert(60)
	root.Insert(80)

	fmt.Println("In-order traversal (Left, Root, Right):")
	root.InOrder() // Should print values in ascending order
	fmt.Println()

	fmt.Println("Pre-order traversal (Root, Left, Right):")
	root.PreOrder() // Should print values in pre-order
	fmt.Println()

	fmt.Println("Post-order traversal (Left, Right, Root):")
	root.PostOrder() // Should print values in post-order
	fmt.Println()

	fmt.Println("Search for 40:", root.Search(40))   // Should return true
	fmt.Println("Search for 100:", root.Search(100)) // Should return false

	fmt.Println("Minimum value:", root.FindMin().Value) // Should return 20
	fmt.Println("Maximum value:", root.FindMax().Value) // Should return 80

	root.Delete(20)
	fmt.Println("In-order after deleting 20:")
	root.InOrder()

	root.Delete(30)
	fmt.Println("\nIn-order after deleting 30:")
	root.InOrder()

	root.Delete(50)
	fmt.Println("\nIn-order after deleting 50:")
	root.InOrder()
}
