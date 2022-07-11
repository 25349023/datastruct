package priorityqueue

import (
	"fmt"
	"math"
)

var logPhi = math.Log(math.Phi)

// FibonacciHeap implementation that is introduced in
// 'Fundamentals of Data Structures in C'
type FibonacciHeap struct {
	min *fHeapNode
	n   int
}

// Min peeks and returns the minimum of the heap.
func (f *FibonacciHeap) Min() (int, error) {
	if f.Empty() {
		return 0, fmt.Errorf("heap is empty")
	}
	return f.min.data, nil
}

// Empty returns whether the heap is empty or not.
func (f *FibonacciHeap) Empty() bool {
	return f.min == nil
}

// Insert x into the FibonacciHeap and return the inserted node.
//
// Amortized cost is O(1).
func (f *FibonacciHeap) Insert(x int) DataNode {
	node := &fHeapNode{data: x}

	defer func() { f.n++ }()

	if f.min == nil {
		f.min, node.prev, node.next = node, node, node
		return node
	}

	f.min.AddSibling(node)

	if x < f.min.data {
		f.min = node
	}
	return node
}

// DeleteMin pops the minimum from the FibonacciHeap then returns it,
// error if the heap is empty.
//
// Amortized cost is O(lg n).
func (f *FibonacciHeap) DeleteMin() (int, error) {
	if f.Empty() {
		return 0, fmt.Errorf("cannot delete-min from empty binomial heap")
	}

	defer func() { f.n-- }()

	minValue := f.min.data
	f.min.pruneParentFromChildren()

	// Step 1: delete min node
	if !isOnly(f.min) {
		f.min.prev.next = f.min.next
		f.min.next.prev = f.min.prev

		subtree := f.min.child
		f.min = f.min.next
		if subtree != nil {
			mergeLists(f.min, subtree)
		}
	} else if f.min.child != nil {
		// Note that in the case of FibonacciHeap,
		// f.min.child may have trees of the same degree.
		// We should merge them, too.
		f.min = f.min.child
	} else {
		f.min = nil
		return minValue, nil
	}

	// Step 2: merge min trees with same degree
	trees := f.mergeSameDegreeTrees()

	// Step 3 & 4: relink the min trees and find min node
	f.relink(trees)

	return minValue, nil
}

func (f *FibonacciHeap) mergeSameDegreeTrees() []*fHeapNode {
	maxDegree := int(math.Log(float64(f.n))/logPhi) + 1
	trees := make([]*fHeapNode, maxDegree)

	for p, next := f.min, f.min.next; ; p, next = next, next.next {
		d := p.degree
		for ; trees[d] != nil; d++ {
			p = joinMinTrees(p, trees[d]).(*fHeapNode)
			trees[d] = nil
		}
		trees[d] = p

		if next == f.min {
			break
		}
	}
	return trees
}

func (f *FibonacciHeap) relink(trees []*fHeapNode) {
	f.min = nil
	for _, node := range trees {
		if node == nil {
			continue
		}
		if f.min == nil {
			f.min, node.next, node.prev = node, node, node
		} else {
			f.min.AddSibling(node)
			if node.data < f.min.data {
				f.min = node
			}
		}
	}
}

// Meld two FibonacciHeap and leave other empty,
// error if the underlying type of other is not FibonacciHeap.
//
// Amortized cost is O(1).
func (f *FibonacciHeap) Meld(other MeldablePQ) error {
	if other, ok := other.(*FibonacciHeap); ok {
		if other.Empty() {
			return nil
		}
		if f.Empty() {
			f.min, f.n = other.min, other.n
			other.min, other.n = nil, 0
			return nil
		}

		mergeLists(f.min, other.min)

		if other.min.data < f.min.data {
			f.min = other.min
		}
		f.n += other.n

		other.min, other.n = nil, 0
		return nil
	} else {
		return fmt.Errorf("cannot meld with non binomial heap")
	}
}

// Delete the specified arbitrary node in the FibonacciHeap f,
// error if f is empty or the target's type is incorrect.
//
// Amortized cost is O(lg n).
func (f *FibonacciHeap) Delete(target DataNode) (int, error) {
	if f.Empty() {
		return 0, fmt.Errorf("cannot delete-min from empty binomial heap")
	}

	if target, ok := target.(*fHeapNode); ok {
		if target == f.min {
			return f.DeleteMin()
		}

		popValue := target.data

		f.cutChild(target, true)
		f.n--

		if target.child != nil {
			mergeLists(f.min, target.child)
		}

		if target.parent != nil {
			f.cascadingCut(target.parent)
		}

		return popValue, nil
	} else {
		return 0, fmt.Errorf("incorrect type of target")
	}
}

// DecreaseKey decrease the key of the specified node in f.
// error if key is greater than original key or the target's type is incorrect.
//
// Amortized cost is O(1).
func (f *FibonacciHeap) DecreaseKey(target DataNode, key int) error {
	if target, ok := target.(*fHeapNode); ok {
		if target.data < key {
			return fmt.Errorf("new key is greater than original key")
		}

		target.data = key

		if p := target.parent; p != nil && target.data < p.data {
			f.cutChild(target, false)
			mergeLists(f.min, target)
			f.cascadingCut(p)
		}

		if target.data < f.min.data {
			f.min = target
		}

		return nil
	} else {
		return fmt.Errorf("incorrect type of target")
	}
}

func (f *FibonacciHeap) cutChild(target *fHeapNode, delete bool) {
	f.removeFromParent(target)

	if delete {
		// set children's parents to nil
		target.pruneParentFromChildren()
	} else {
		target.parent = nil
	}
}

func (f *FibonacciHeap) removeFromParent(target *fHeapNode) {
	if p := target.parent; p != nil {
		if p.child == target {
			if isOnly(target) {
				p.child = nil
			} else {
				p.child = target.next
			}
		}
		p.degree--
	}

	target.prev.next = target.next
	target.next.prev = target.prev
	target.prev, target.next = target, target
}

func (f *FibonacciHeap) cascadingCut(target *fHeapNode) {
	parent := target.parent
	if parent == nil {
		return
	}

	if target.lostChild {
		f.cutChild(target, false)
		mergeLists(f.min, target)
		f.cascadingCut(parent)
	} else {
		target.lostChild = true
	}
}
