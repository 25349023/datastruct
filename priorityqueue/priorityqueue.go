package priorityqueue

type PriorityQueue interface {
	Empty() bool
	Insert(x int) DataNode
	DeleteMin() (int, error)
	Min() (int, error)
}

type MeldablePQ interface {
	PriorityQueue
	Meld(other MeldablePQ) error
}

type CompletePQ interface {
	MeldablePQ
	Delete(target DataNode) (int, error)
	DecreaseKey(target DataNode, key int) error
}
