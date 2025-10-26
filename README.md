# datastruct

Implement some useful data structures using Go.

[Module page on pkg.go.dev](https://pkg.go.dev/github.com/25349023/datastruct)

## Examples

### Priority Queue (Fibonacci Heap)
```go
package main

import (
	"fmt"
	"github.com/25349023/datastruct/priorityqueue"
	"log"
)

func main() {
	var heap1, heap2 priorityqueue.FibonacciHeap

	for _, v := range []int{7, 3, 4, 10, 5} {
		heap1.Insert(v)
	}

	for _, v := range []int{2, 6, 8, 17, 1} {
		heap2.Insert(v)
	}

	err := heap1.Meld(&heap2)
	if err != nil {
		log.Fatal(err)
	}

	for !heap1.Empty() {
		m, err := heap1.DeleteMin()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}
}
```

### Red-Black Tree
```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/25349023/datastruct/bst"
)

func main() {
	var tree bst.RBTree[int]
	tree.Init(func(a, b int) bool {
		return a < b
	})

	var nodes []*bst.RBNode[int]
	fmt.Print("Insert: ")
	for i := 0; i < 15; i++ {
		v := rand.IntN(98) + 1
		n := tree.Insert(v)
		nodes = append(nodes, n)
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	tree.DrawTree()

	fmt.Print("Deleting Node: ")
	for i := 7; i >= 0; i-- {
		fmt.Printf("%v ", nodes[i].Data)
		tree.Delete(nodes[i])
	}
	fmt.Println()
	tree.DrawTree()
}

```



