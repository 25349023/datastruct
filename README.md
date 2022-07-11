# datastruct

Implement some useful data structures using Go.

[Module page on pkg.go.dev](https://pkg.go.dev/github.com/25349023/datastruct)

## Examples

### priorityqueue
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



