package bst

type Color int

const (
	RED Color = iota
	BLACK
)

type RBNode[T any] struct {
	Data        T
	left, right *RBNode[T]
	parent      *RBNode[T]
	color       Color
}
