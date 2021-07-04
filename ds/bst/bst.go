// binary search tree
package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move right
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)

	} else if n.Key > k {
		// move right
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 10}
	tree.Insert(3)
	tree.Insert(2)
	fmt.Println(tree, tree.Left)
	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(43)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(14)
	tree.Insert(19)
	fmt.Println(tree.Search(43))
	fmt.Println(tree.Search(90))

}
