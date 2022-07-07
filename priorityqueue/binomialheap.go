package priorityqueue

import (
	"fmt"
	"math"
)

type bHeapNode struct {
	data              int
	degree            int
	child, prev, next *bHeapNode
}

func findMinNode(list *bHeapNode) *bHeapNode {
	if list == nil {
		return nil
	}

	min := list
	for curr := list.next; curr != list; curr = curr.next {
		if curr.data < min.data {
			min = curr
		}
	}
	return min
}

func mergeLists(x, y *bHeapNode) {
	x.next.prev = y.prev
	y.prev.next = x.next
	x.next = y
	y.prev = x
}

func joinMinTrees(x, y *bHeapNode) *bHeapNode {
	if y == nil {
		return x
	}
	if x == nil {
		return y
	}

	if x.data < y.data {
		x.addChild(y)
		return x
	} else {
		y.addChild(x)
		return y
	}
}

func (head *bHeapNode) addSibling(s *bHeapNode) {
	s.next = head
	s.prev = head.prev
	head.prev.next = s
	head.prev = s
}

func (head *bHeapNode) addChild(ch *bHeapNode) {
	if ch == nil {
		return
	}
	if head.child == nil {
		head.child, ch.next, ch.prev = ch, ch, ch
		head.degree = 1
		return
	}

	head.child.addSibling(ch)
	head.degree++
}

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

// Insert x into the BinomialHeap.
func (b *BinomialHeap) Insert(x int) {
	node := &bHeapNode{data: x}

	defer func() { b.n++ }()

	if b.min == nil {
		b.min, node.prev, node.next = node, node, node
		return
	}

	b.min.addSibling(node)

	if x < b.min.data {
		b.min = node
	}
}

// DeleteMin pops the minimum from the BinomialHeap then returns it,
// error if the heap is empty.
func (b *BinomialHeap) DeleteMin() (int, error) {
	if b.Empty() {
		return 0, fmt.Errorf("cannot delete-min from empty binomial heap")
	}

	defer func() { b.n-- }()

	minValue := b.min.data

	if b.min.next == b.min {
		b.min = findMinNode(b.min.child)
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
			p = joinMinTrees(p, trees[d])
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
			b.min.addSibling(node)
			if node.data < b.min.data {
				b.min = node
			}
		}
	}
}

// Meld two BinomialHeap and leave other empty,
// error if the underlying type of other is not BinomialHeap
func (b *BinomialHeap) Meld(other MeldablePQ) error {
	if other, ok := other.(*BinomialHeap); ok {
		if other.Empty() {
			return nil
		}
		if b.Empty() {
			b.min, other.min = other.min, nil
			b.n = other.n
			return nil
		}

		mergeLists(b.min, other.min)

		if other.min.data < b.min.data {
			b.min = other.min
		}
		b.n += other.n

		other.min = nil
		return nil
	} else {
		return fmt.Errorf("cannot meld with non binomial heap")
	}
}
