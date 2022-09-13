package main

import "fmt"


// Node
type Node struct {
	Key int
	Left *Node
	Right *Node
}

var count int


// Insert will add a note to the tree
func (n *Node) Insert (k int) {
	if n.Key < k {
		// move right 
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search
func (n *Node) Search (k int) bool{
    count++

    if n == nil {
        return false
    }

    if n.Key < k {
        // move right
        return n.Right.Search(k)
    } else if n.Key > k {
        // move left
        return n.Left.Search(k)
    }
    return true
}

func main() {
	tree := &Node{Key:100}
	tree.Insert(200)
	tree.Insert(300)
    tree.Insert(76)
	tree.Insert(19)
    tree.Insert(150)
	tree.Insert(310)
    tree.Insert(7)
	tree.Insert(24)
    tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(310))
    fmt.Println(count)

}