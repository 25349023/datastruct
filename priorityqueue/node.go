package priorityqueue

import "reflect"

// DataNode exposes only one API, which returns the contained data
type DataNode interface {
	Data() int
}

// BiDirTreeNode defines operations of bidirectional linked-list node,
// which supports getting/setting next/prev node,
// and add new node to the children list
type BiDirTreeNode interface {
	DataNode
	Next() BiDirTreeNode
	Prev() BiDirTreeNode
	SetNext(n BiDirTreeNode)
	SetPrev(n BiDirTreeNode)
	AddChild(ch BiDirTreeNode)
}

func isNilPtr(x interface{}) bool {
	return x == nil || (reflect.ValueOf(x).Kind() == reflect.Ptr && reflect.ValueOf(x).IsNil())
}

func isOnly(x BiDirTreeNode) bool {
	return x == x.Next()
}

func findMinNode(list BiDirTreeNode) BiDirTreeNode {
	if isNilPtr(list) {
		return list
	}

	min := list
	for curr := list.Next(); curr != list; curr = curr.Next() {
		if curr.Data() < min.Data() {
			min = curr
		}
	}
	return min
}

func mergeLists(x, y BiDirTreeNode) {
	x.Next().SetPrev(y.Prev())
	y.Prev().SetNext(x.Next())
	x.SetNext(y)
	y.SetPrev(x)
}

func joinMinTrees(x, y BiDirTreeNode) BiDirTreeNode {
	if isNilPtr(y) {
		return x
	}
	if isNilPtr(x) {
		return y
	}

	if x.Data() < y.Data() {
		x.AddChild(y)
		return x
	} else {
		y.AddChild(x)
		return y
	}
}
