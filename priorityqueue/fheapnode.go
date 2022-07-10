package priorityqueue

type fHeapNode struct {
	data          int
	degree        int
	lostChild     bool
	prev, next    *fHeapNode
	child, parent *fHeapNode
}

func (n *fHeapNode) Data() int {
	return n.data
}

func (n *fHeapNode) Next() BiDirTreeNode {
	return n.next
}

func (n *fHeapNode) Prev() BiDirTreeNode {
	return n.prev
}

func (n *fHeapNode) SetNext(other BiDirTreeNode) {
	n.next = other.(*fHeapNode)
}

func (n *fHeapNode) SetPrev(other BiDirTreeNode) {
	n.prev = other.(*fHeapNode)
}

func (n *fHeapNode) AddSibling(s *fHeapNode) {
	s.next = n
	s.prev = n.prev
	n.prev.next = s
	n.prev = s
}

func (n *fHeapNode) AddChild(ch BiDirTreeNode) {
	if isNilPtr(ch) {
		return
	}

	if ch, ok := ch.(*fHeapNode); ok {
		ch.parent = n
		ch.lostChild = false

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

func (n *fHeapNode) pruneParentFromChildren() {
	for c := n.child; c != nil && c.parent != nil; c = c.next {
		c.parent = nil
	}
}
