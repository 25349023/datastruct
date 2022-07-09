package priorityqueue

import "reflect"

type BiDirTreeNode interface {
	Data() int
	Next() BiDirTreeNode
	Prev() BiDirTreeNode
	SetNext(n BiDirTreeNode)
	SetPrev(n BiDirTreeNode)
	AddChild(ch BiDirTreeNode)
}

func isNilPtr(x interface{}) bool {
	return x == nil || (reflect.ValueOf(x).Kind() == reflect.Ptr && reflect.ValueOf(x).IsNil())
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
	if y == nil {
		return x
	}
	if x == nil {
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
