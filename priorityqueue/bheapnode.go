package priorityqueue

type bHeapNode struct {
	data              int
	degree            int
	child, prev, next *bHeapNode
}

func (n *bHeapNode) Data() int {
	return n.data
}

func (n *bHeapNode) Next() BiDirTreeNode {
	return n.next
}

func (n *bHeapNode) Prev() BiDirTreeNode {
	return n.prev
}

func (n *bHeapNode) SetNext(other BiDirTreeNode) {
	n.next = other.(*bHeapNode)
}

func (n *bHeapNode) SetPrev(other BiDirTreeNode) {
	n.prev = other.(*bHeapNode)
}

func (n *bHeapNode) AddSibling(s *bHeapNode) {
	s.next = n
	s.prev = n.prev
	n.prev.next = s
	n.prev = s
}

func (n *bHeapNode) AddChild(ch BiDirTreeNode) {
	if isNilPtr(ch) {
		return
	}

	if ch, ok := ch.(*bHeapNode); ok {
		if n.child == nil {
			n.child, ch.next, ch.prev = ch, ch, ch
			n.degree = 1
			return
		}

		n.child.AddSibling(ch)
		n.degree++
	} else {
		panic("the type of ch is incorrect")
	}
}
