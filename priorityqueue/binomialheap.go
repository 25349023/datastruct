package priorityqueue

import (
	"fmt"
	"math"
)

// BinomialHeap implementation that is introduced in
// 'Fundamentals of Data Structures in C'
type BinomialHeap struct {
	min *bHeapNode
	n   int
}

// Min peeks and returns the minimum of the heap.
func (b *BinomialHeap) Min() (int, error) {
	if b.Empty() {
		return 0, fmt.Errorf("heap is empty")
	}
	return b.min.data, nil
}

// Empty returns whether the heap is empty or not.
func (b *BinomialHeap) Empty() bool {
	return b.min == nil
}

// Insert x into the BinomialHeap and return the inserted node.
//
// Amortized cost is O(1).
func (b *BinomialHeap) Insert(x int) DataNode {
	node := &bHeapNode{data: x}

	defer func() { b.n++ }()

	if b.min == nil {
		b.min, node.prev, node.next = node, node, node
		return node
	}

	b.min.AddSibling(node)

	if x < b.min.data {
		b.min = node
	}

	return node
}

// DeleteMin pops the minimum from the BinomialHeap then returns it,
// error if the heap is empty.
//
// Amortized cost is O(lg n).
func (b *BinomialHeap) DeleteMin() (int, error) {
	if b.Empty() {
		return 0, fmt.Errorf("cannot delete-min from empty binomial heap")
	}

	defer func() { b.n-- }()

	minValue := b.min.data

	if isOnly(b.min) {
		b.min = findMinNode(b.min.child).(*bHeapNode)
		return minValue, nil
	}

	// Step 1: delete min node
	b.min.prev.next = b.min.next
	b.min.next.prev = b.min.prev

	subtree := b.min.child
	b.min = b.min.next
	if subtree != nil {
		mergeLists(b.min, subtree)
	}

	// Step 2: merge min trees with same degree
	trees := b.mergeSameDegreeTrees()

	// Step 3 & 4: relink the min trees and find min node
	b.relink(trees)

	return minValue, nil
}

func (b *BinomialHeap) mergeSameDegreeTrees() []*bHeapNode {
	maxDegree := int(math.Log2(float64(b.n))) + 1
	trees := make([]*bHeapNode, maxDegree)

	for p, next := b.min, b.min.next; ; p, next = next, next.next {
		d := p.degree
		for ; trees[d] != nil; d++ {
			p = joinMinTrees(p, trees[d]).(*bHeapNode)
			trees[d] = nil
		}
		trees[d] = p

		if next == b.min {
			break
		}
	}
	return trees
}

func (b *BinomialHeap) relink(trees []*bHeapNode) {
	b.min = nil
	for _, node := range trees {
		if node == nil {
			continue
		}
		if b.min == nil {
			b.min, node.next, node.prev = node, node, node
		} else {
			b.min.AddSibling(node)
			if node.data < b.min.data {
				b.min = node
			}
		}
	}
}

// Meld two BinomialHeap and leave other empty,
// error if the underlying type of other is not BinomialHeap.
//
// Amortized cost is O(1).
func (b *BinomialHeap) Meld(other MeldablePQ) error {
	if other, ok := other.(*BinomialHeap); ok {
		if other.Empty() {
			return nil
		}
		if b.Empty() {
			b.min, b.n = other.min, other.n
			other.min, other.n = nil, 0
			return nil
		}

		mergeLists(b.min, other.min)

		if other.min.data < b.min.data {
			b.min = other.min
		}
		b.n += other.n

		other.min, other.n = nil, 0
		return nil
	} else {
		return fmt.Errorf("cannot meld with non binomial heap")
	}
}
