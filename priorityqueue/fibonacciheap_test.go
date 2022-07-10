package priorityqueue

import (
	"math/rand"
	"testing"
	"time"
)

func TestFibonacciHeap_Empty(t *testing.T) {
	f := FibonacciHeap{}
	if !f.Empty() {
		t.Fatal("f.Empty() should be true")
	}

	f.Insert(1)
	f.Insert(2)
	if f.Empty() {
		t.Fatal("f.Empty() should be false after insertion")
	}

	_, _ = f.DeleteMin()
	_, _ = f.DeleteMin()
	if !f.Empty() {
		t.Fatal("f.Empty() should be true after delete all elements")
	}
}

func TestFibonacciHeap_Insert(t *testing.T) {
	f := FibonacciHeap{}
	for _, v := range []int{5, 2, 4, 3, 1} {
		hd := f.Insert(v)
		if _, ok := hd.(*fHeapNode); !ok {
			t.Fatal("incorrect underlying type")
		}
		if hd.Data() != v {
			t.Fatalf("got: %d, expect: %d", hd.Data(), v)
		}
	}

	if f.min.data != 1 {
		t.Fatal("minimum of f should be 1, got", f.min.data)
	}
	if f.n != 5 {
		t.Fatal("f.n should be 5, got", f.n)
	}
}

func TestFibonacciHeap_DeleteMin(t *testing.T) {
	f := FibonacciHeap{}
	for _, v := range []int{5, 2, 4, 3, 1} {
		f.Insert(v)
	}
	for ans := 1; !f.Empty(); ans++ {
		v, err := f.DeleteMin()
		if err != nil {
			t.Fatal(err)
		}
		if v != ans {
			t.Fatalf("got: %d, expect: %d", v, ans)
		}
	}
	_, err := f.DeleteMin()
	if err == nil {
		t.Fatal("should report error when f is empty")
	}
	if f.n != 0 {
		t.Fatal("f.n should be 0")
	}
}

func TestFibonacciHeap_Min(t *testing.T) {
	f := FibonacciHeap{}

	_, err := f.Min()
	if err == nil {
		t.Fatal("should report error when f is empty")
	}

	for _, v := range [][2]int{{5, 5}, {2, 2}, {4, 2}, {3, 2}, {1, 1}, {6, 1}} {
		x, a := v[0], v[1]
		f.Insert(x)
		if y, err := f.Min(); y != a {
			if err != nil {
				t.Fatal(err)
			}
			t.Fatalf("get: %d, expect: %d", y, a)
		}
	}
}

func TestFibonacciHeap_Meld(t *testing.T) {
	var f1, f2 FibonacciHeap

	// f1 == f2 == empty
	err := f1.Meld(&f2)
	if err != nil {
		t.Fatal(err)
	}
	if !f1.Empty() || !f2.Empty() {
		t.Fatal("both f1 and f2 should be empty")
	}

	// f1 != empty, f2 == empty
	f1.Insert(1)
	f1.Insert(2)
	err = f1.Meld(&f2)
	if err != nil {
		t.Fatal(err)
	}
	if !f2.Empty() || f2.n != 0 {
		t.Fatal("f2 should be empty")
	}
	if f1.n != 2 {
		t.Fatal("f1.n should be 2")
	}
	if y, _ := f1.Min(); y != 1 {
		t.Fatal("f1.Min() should be 1")
	}

	// f1 == empty, f2 != empty
	f1, f2 = f2, f1
	err = f1.Meld(&f2)
	if err != nil {
		t.Fatal(err)
	}
	if !f2.Empty() || f2.n != 0 {
		t.Fatal("f2 should be empty")
	}
	if f1.n != 2 {
		t.Fatal("f1.n should be 2")
	}
	if y, _ := f1.Min(); y != 1 {
		t.Fatal("f1.Min() should be 1")
	}
}

func TestFibonacciHeap_Meld2(t *testing.T) {
	var f1, f2 FibonacciHeap
	for _, v := range []int{5, 2, 7, 6, 9} {
		f1.Insert(v)
	}
	for _, v := range []int{8, 3, 4, 1, 10} {
		f2.Insert(v)
	}

	err := f1.Meld(&f2)
	if err != nil {
		t.Fatal(err)
	}

	if !f2.Empty() || f2.n != 0 {
		t.Fatal("f2 should be empty")
	}

	if f1.n != 10 {
		t.Fatal("f1.n should be 10, got", f1.n)
	}

	if y, _ := f1.Min(); y != 1 {
		t.Fatalf("got: %d, expect: 1", y)
	}

	for ans := 1; !f1.Empty(); ans++ {
		v, _ := f1.DeleteMin()
		if v != ans {
			t.Fatalf("got %d, expect %d", v, ans)
		}
	}

	if !f1.Empty() || f1.n != 0 {
		t.Fatal("f1 should be empty after delete elements")
	}
}

func TestFibonacciHeap_Meld3(t *testing.T) {
	f1 := FibonacciHeap{}
	var m MeldablePQ

	err := f1.Meld(m)
	if err == nil {
		t.Fatal("should report error when other is not Fibonacci Heap")
	}
}

func TestFibonacciHeap_Delete(t *testing.T) {
	f := FibonacciHeap{}
	hd := make([]DataNode, 10)
	for _, v := range []int{5, 2, 7, 6, 9, 1, 8, 4, 3} {
		hd[v] = f.Insert(v)
	}

	n := 8
	for _, v := range []int{1, 8, 4, 2, 3} {
		p, err := f.Delete(hd[v])
		if err != nil {
			t.Fatal(err)
		}
		if p != v {
			t.Fatalf("got: %d, expect: %d", p, v)
		}
		if f.n != n {
			t.Fatal("f.n should be", n)
		}
		n--
	}

	if p, _ := f.Min(); p != 5 {
		t.Fatalf("got: %d, expect: 5", p)
	}

	for _, v := range []int{6, 5, 9, 7} {
		p, err := f.Delete(hd[v])
		if err != nil {
			t.Fatal(err)
		}
		if p != v {
			t.Fatalf("got: %d, expect: %d", p, v)
		}
		if f.n != n {
			t.Fatal("f.n should be", n)
		}
		n--
	}

	if !f.Empty() {
		t.Fatal("f should be empty: ", f.min.data)
	}
}

func TestFibonacciHeap_Delete2(t *testing.T) {
	f := FibonacciHeap{}
	fn := &fHeapNode{data: 0}
	_, err := f.Delete(fn)
	if err == nil {
		t.Fatal("should report empty deletion")
	}

	f.Insert(5)

	bn := &bHeapNode{data: 5}
	_, err = f.Delete(bn)
	if err == nil {
		t.Fatal("should report incorrect type")
	}
}

func TestFibonacciHeap_RandomDelete(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	f := FibonacciHeap{}
	hd := make([]DataNode, 130)

	insert := rng.Perm(128)
	t.Logf("insert values: %v", insert)
	for _, v := range insert {
		hd[v] = f.Insert(v)
	}
	min, err := f.DeleteMin()
	if err != nil {
		t.Fatal(f)
	}
	if min != 0 {
		t.Fatalf("got: %d, expect: %d", min, 0)
	}

	n := 127
	deletion := rng.Perm(128)
	t.Logf("delete values: %v", deletion)
	for _, v := range deletion {
		if v == 0 || hd[v] == nil {
			continue
		}

		p, err := f.Delete(hd[v])
		n--
		if err != nil {
			t.Fatal(err)
		}
		if p != v {
			t.Fatalf("got: %d, expect: %d", p, v)
		}
		if f.n != n {
			t.Fatal("f.n should be", n)
		}
	}

	if !f.Empty() {
		t.Fatal("f should be empty: ", f.min.data)
	}
}

func TestFibonacciHeap_DecreaseKey(t *testing.T) {
	f := FibonacciHeap{}
	hd := make([]DataNode, 10)
	for _, v := range []int{5, 2, 7, 6, 9, 1, 8, 4, 3} {
		hd[v] = f.Insert(v)
	}

	for _, v := range []int{1, 8, 4, 2, 3} {
		err := f.DecreaseKey(hd[v], -v)
		if err != nil {
			t.Fatal(err)
		}
	}

	if p, _ := f.DeleteMin(); p != -8 {
		t.Fatalf("got: %d, expect: -8", p)
	}

	for _, ans := range []int{-4, -3, -2, -1, 5, 6, 7, 9} {
		p, err := f.DeleteMin()
		if err != nil {
			t.Fatal(err)
		}
		if p != ans {
			t.Errorf("got: %d, expect: %d", p, ans)
		}
	}

	if !f.Empty() {
		t.Fatal("f should be empty: ", f.min.data)
	}
}

func TestFibonacciHeap_DecreaseKey2(t *testing.T) {
	f := FibonacciHeap{}
	hd := make([]DataNode, 10)
	for _, v := range []int{5, 2, 7, 6} {
		hd[v] = f.Insert(v)
	}

	err := f.DecreaseKey(hd[5], 10)
	if err == nil {
		t.Fatal("increase a key is not valid")
	}

	err = f.DecreaseKey(&bHeapNode{data: 5}, 1)
	if err == nil {
		t.Fatal("should report incorrect type")
	}
}

func TestFibonacciHeap_RandomDecreaseKey(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	f := FibonacciHeap{}
	hd := make([]DataNode, 130)

	insert := rng.Perm(128)
	t.Logf("insert values: %v", insert)
	for _, v := range insert {
		hd[v] = f.Insert(v)
	}

	min, err := f.DeleteMin()
	if err != nil {
		t.Fatal(f)
	}
	if min != 0 {
		t.Fatalf("got: %d, expect: %d", min, 0)
	}

	decrement := rng.Perm(128)
	t.Logf("decrease values: %v", decrement)
	for _, v := range decrement {
		if v == 0 || hd[v] == nil {
			continue
		}

		err := f.DecreaseKey(hd[v], -v)
		if err != nil {
			t.Fatal(err)
		}
	}

	for ans := -127; ans < 0; ans++ {
		p, err := f.DeleteMin()
		if err != nil {
			t.Fatal(err)
		}
		if p != ans {
			t.Errorf("got: %d, expect: %d", p, ans)
		}
	}

	if !f.Empty() {
		t.Fatal("f should be empty: ", f.min.data)
	}
}
