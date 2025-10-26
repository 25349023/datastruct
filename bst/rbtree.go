package bst

import (
	"fmt"
	"math"
	"strings"

	"github.com/gookit/color"
)

type RBTree[T any] struct {
	Nil  *RBNode[T]
	Root *RBNode[T]
	Less func(a, b T) bool
}

func (rbt *RBTree[T]) Init(less func(a, b T) bool) {
	rbt.Nil = &RBNode[T]{
		color: BLACK,
	}
	rbt.Nil.left, rbt.Nil.right, rbt.Nil.parent = rbt.Nil, rbt.Nil, rbt.Nil
	rbt.Root = rbt.Nil
	rbt.Less = less
}

func (rbt *RBTree[T]) Equal(a, b RBNode[T]) bool {
	return !rbt.Less(a.Data, b.Data) && !rbt.Less(b.Data, b.Data)
}

func (rbt *RBTree[T]) NewRBNode(data T) *RBNode[T] {
	return &RBNode[T]{
		Data:   data,
		left:   rbt.Nil,
		right:  rbt.Nil,
		parent: rbt.Nil,
		color:  RED,
	}
}

func (rbt *RBTree[T]) Min() *RBNode[T] {
	if rbt.Root == rbt.Nil {
		return rbt.Nil
	}

	x := rbt.Root
	for ; x.left != rbt.Nil; x = x.left {
		fmt.Printf("%v ", x.Data)
	}
	return x
}

func (rbt *RBTree[T]) minUnder(node *RBNode[T]) *RBNode[T] {
	if node == rbt.Nil {
		return rbt.Nil
	}

	x := node
	for ; x.left != rbt.Nil; x = x.left {
	}
	return x
}

func (rbt *RBTree[T]) Max() *RBNode[T] {
	if rbt.Root == rbt.Nil {
		return rbt.Nil
	}

	x := rbt.Root
	for ; x.right != rbt.Nil; x = x.right {
		fmt.Printf("%v ", x.Data)
	}
	return x
}

func (rbt *RBTree[T]) Inorder(node *RBNode[T]) {
	if node == rbt.Nil {
		return
	}
	rbt.Inorder(node.left)
	fmt.Printf("%v ", node.Data)
	rbt.Inorder(node.right)
}

func (rbt *RBTree[T]) rotateLeft(x *RBNode[T]) {
	y := x.right
	x.right = y.left
	if x.right != rbt.Nil {
		x.right.parent = x
	}
	y.left = x

	if x == rbt.Root {
		rbt.Root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.parent = x.parent
	x.parent = y
}

func (rbt *RBTree[T]) rotateRight(x *RBNode[T]) {
	y := x.left
	x.left = y.right
	if x.left != rbt.Nil {
		x.left.parent = x
	}
	y.right = x

	if x == rbt.Root {
		rbt.Root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.parent = x.parent
	x.parent = y
}

func (rbt *RBTree[T]) Insert(data T) *RBNode[T] {
	z := rbt.NewRBNode(data)

	rbt.Root = rbt.insertTo(rbt.Nil, rbt.Root, z)
	rbt.insertFixup(z)
	rbt.Root.color = BLACK

	return z
}

func (rbt *RBTree[T]) insertTo(p, x, z *RBNode[T]) *RBNode[T] {
	if x == rbt.Nil {
		z.parent = p
		return z
	}

	if rbt.Less(z.Data, x.Data) {
		x.left = rbt.insertTo(x, x.left, z)
	} else {
		x.right = rbt.insertTo(x, x.right, z)
	}
	return x
}

func (rbt *RBTree[T]) insertFixup(z *RBNode[T]) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			if u := z.parent.parent.right; u.color == RED { // Case 1: z's uncle is red
				z.parent.parent.color = RED
				z.parent.color = BLACK
				u.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.right { // Case 2: z's uncle is black, z is right child
					rbt.rotateLeft(z.parent)
					z = z.left
				}
				z.parent.parent.color = RED // Case 3: z's uncle is black, z is left child
				z.parent.color = BLACK
				rbt.rotateRight(z.parent.parent)
			}
		} else {
			if u := z.parent.parent.left; u.color == RED { // Case 1: z's uncle is red
				z.parent.parent.color = RED
				z.parent.color = BLACK
				u.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.left { // Case 2: z's uncle is black, z is left child
					rbt.rotateRight(z.parent)
					z = z.right
				}
				z.parent.parent.color = RED // Case 3: z's uncle is black, z is right child
				z.parent.color = BLACK
				rbt.rotateLeft(z.parent.parent)
			}
		}
	}
}

func (rbt *RBTree[T]) Delete(z *RBNode[T]) {
	y := z // y is the node to be deleted actually
	yOriginalColor := y.color
	x := rbt.Nil // x is the node that will replace y's original position

	if z.left == rbt.Nil {
		x = z.right
		rbt.transplant(z, z.right)
	} else if z.right == rbt.Nil {
		x = z.left
		rbt.transplant(z, z.left)
	} else {
		y = rbt.minUnder(z.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent != z {
			rbt.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		} else {
			x.parent = y // for the case that x is Nil
		}
		rbt.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		rbt.deleteFixup(x)
	}

	// Correcting drawing
	rbt.Nil.left, rbt.Nil.right, rbt.Nil.parent = rbt.Nil, rbt.Nil, rbt.Nil
}

func (rbt *RBTree[T]) deleteFixup(x *RBNode[T]) {
	for x != rbt.Root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED { // Case 1: x's sibling is red
				x.parent.color, w.color = RED, BLACK
				rbt.rotateLeft(x.parent)
				w = x.parent.right
			}
			// Case 2: w's both children are black
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				// Case 3: w.left is red, w.right is black
				if w.right.color == BLACK {
					w.color, w.left.color = RED, BLACK
					rbt.rotateRight(w)
					w = x.parent.right
				}
				// Case 4: w.right is red
				x.parent.color, w.color = w.color, x.parent.color
				rbt.rotateLeft(x.parent)
				w.right.color = BLACK
				x = rbt.Root
			}
		} else {
			w := x.parent.left
			if w.color == RED {
				x.parent.color, w.color = RED, BLACK
				rbt.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.color, w.right.color = RED, BLACK
					rbt.rotateLeft(w)
					w = x.parent.left
				}
				x.parent.color, w.color = w.color, x.parent.color
				rbt.rotateRight(x.parent)
				w.left.color = BLACK
				x = rbt.Root
			}
		}
	}
	x.color = BLACK
}

func (rbt *RBTree[T]) transplant(u, v *RBNode[T]) {
	if u.parent == rbt.Nil {
		rbt.Root = v
	}
	if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// DrawTree should restrict key data to a maximum of 2 digits to avoid layout distortion
func (rbt *RBTree[T]) DrawTree() {
	h := rbt.height(rbt.Root)
	btmWidth := int(math.Pow(2, float64(h))) * 2
	pos := btmWidth/2 - 1
	spaces := strings.Repeat("  ", pos)
	red := color.FgRed.Render

	type nodeInfo struct {
		level int
		dir   int
		*RBNode[T]
	}

	queue := make([]nodeInfo, 0, btmWidth)
	queue = append(queue, nodeInfo{0, 0, rbt.Root})
	lastLvl := 0
	for len(queue) > 0 {
		x := queue[0].RBNode
		lvl := queue[0].level
		if lvl > lastLvl {
			btmWidth /= 2
			fmt.Println()
			if btmWidth == 1 {
				break
			}

			pos = btmWidth/2 - 1
			spaces = strings.Repeat("  ", pos)

			// draw edges
			for i := 0; i < len(queue); i++ {
				fmt.Print(spaces)
				if queue[i].RBNode == rbt.Nil {
					fmt.Print("    ")
				} else if queue[i].dir == 0 {
					fmt.Printf("  ／")
				} else {
					fmt.Printf("＼  ")
				}
				fmt.Print(spaces)
			}
			fmt.Println()
		}
		lastLvl = lvl

		queue = queue[1:]
		queue = append(queue, nodeInfo{lvl + 1, 0, x.left})
		queue = append(queue, nodeInfo{lvl + 1, 1, x.right})

		fmt.Print(spaces)
		if x == rbt.Nil {
			fmt.Print("    ")
		} else if x.color == RED {
			fmt.Printf("%s%2v", red("██"), x.Data)
		} else {
			fmt.Printf("██%2v", x.Data)
		}
		fmt.Print(spaces)
	}
}

func (rbt *RBTree[T]) height(node *RBNode[T]) int {
	if node == rbt.Nil {
		return 0
	}
	return max(rbt.height(node.left), rbt.height(node.right)) + 1
}
