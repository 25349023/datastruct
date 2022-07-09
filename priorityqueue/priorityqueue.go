package priorityqueue

type PriorityQueue interface {
	Empty() bool
	Insert(x int) BiDirTreeNode
	DeleteMin() (int, error)
	Min() (int, error)
}

type MeldablePQ interface {
	PriorityQueue
	Meld(other MeldablePQ) error
}

type CompletePQ interface {
	MeldablePQ
	Delete(target BiDirTreeNode) (int, error)
	DecreaseKey(target BiDirTreeNode, key int) error
}
