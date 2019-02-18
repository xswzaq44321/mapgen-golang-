package BST

import (
	"errors"
)

type Comparable interface {
	LessThan(Comparable) bool
	EqualTo(Comparable) bool
	GreaterThan(Comparable) bool
}

type Node struct {
	Data  Comparable
	Left  *Node
	Right *Node
}

func (n *Node) Insert(data Comparable) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}
	switch {
	case data.EqualTo(n.Data):
		return nil
	case data.LessThan(n.Data):
		if n.Left == nil {
			n.Left = &Node{Data: data}
			return nil
		}
		return n.Left.Insert(data)
	case data.GreaterThan(n.Data):
		if n.Right == nil {
			n.Right = &Node{Data: data}
			return nil
		}
		return n.Right.Insert(data)
	}
	return nil
}

func (n *Node) Find(data Comparable) (Comparable, bool) {
	if n == nil {
		return nil, false
	}

	switch {
	case data.EqualTo(n.Data):
		return n.Data, true
	case data.LessThan(n.Data):
		return n.Left.Find(data)
	default:
		return n.Right.Find(data)
	}
}

//findMax finds the maximum element in a (sub-)tree.
//Return values: the node itself and its parent node.
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

//replaceNode replaces the parent’s child pointer to n with a pointer to the replacement node. parent must not be nil.
func (n *Node) replaceNode(parent, replacement *Node) error {
	if parent == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}
	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) Delete(s Comparable, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	// search the node to be deleted
	switch {
	case s.LessThan(n.Data):
		return n.Left.Delete(s, n)
	case s.GreaterThan(n.Data):
		return n.Right.Delete(s, n)
	// found the node to be deleted
	default:
		// no child, just remove it
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		// one child, replace it with it's child
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		// two child, find the maximum element in the left sub-tree
		replacement, replParent := n.Left.findMax(n)

		// replace the node's data with replacement's data
		n.Data = replacement.Data

		// remove the replacement
		return replacement.Delete(replacement.Data, replParent)
	}
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(data Comparable) error {
	if t.Root == nil {
		t.Root = &Node{Data: data}
		return nil
	}

	return t.Root.Insert(data)
}

func (t *Tree) Find(data Comparable) (Comparable, bool) {
	if t.Root == nil {
		return nil, false
	}
	return t.Root.Find(data)
}

func (t *Tree) Delete(data Comparable) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	//Passing a “fake” parent node here almost avoids having to treat the root node as a special case.
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(data, fakeParent)
	if err != nil {
		return err
	}

	// whether or not root got deleted, set t.Root to new root
	t.Root = fakeParent.Right
	return nil
}

// travel the tree in LVR order, and on each node, invoke f function
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}
