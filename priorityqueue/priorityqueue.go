package priorityqueue

type PriorityQueue interface {
	Empty() bool
	Insert(x int)
	DeleteMin() (int, error)
	Min() (int, error)
}

type MeldablePQ interface {
	PriorityQueue
	Meld(other MeldablePQ) error
}
