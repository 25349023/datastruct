package priorityqueue

import (
	"testing"
)

func TestBinomialHeap_Empty(t *testing.T) {
	b := BinomialHeap{}
	if !b.Empty() {
		t.Fatal("b.Empty() should be true")
	}

	b.Insert(1)
	b.Insert(2)
	if b.Empty() {
		t.Fatal("b.Empty() should be false after insertion")
	}

	_, _ = b.DeleteMin()
	_, _ = b.DeleteMin()
	if !b.Empty() {
		t.Fatal("b.Empty() should be true after delete all elements")
	}
}

func TestBinomialHeap_Insert(t *testing.T) {
	b := BinomialHeap{}

	for _, v := range []int{5, 2, 4, 3, 1} {
		hd := b.Insert(v)
		if _, ok := hd.(*bHeapNode); !ok {
			t.Fatal("incorrect underlying type")
		}
		if hd.Data() != v {
			t.Fatalf("got: %d, expect: %d", hd.Data(), v)
		}
	}

	if b.min.data != 1 {
		t.Fatal("minimum of b should be 1, got", b.min.data)
	}
	if b.n != 5 {
		t.Fatal("b.n should be 5, got", b.n)
	}
}

func TestBinomialHeap_DeleteMin(t *testing.T) {
	b := BinomialHeap{}
	for _, v := range []int{5, 2, 4, 3, 1} {
		b.Insert(v)
	}
	for ans := 1; !b.Empty(); ans++ {
		v, err := b.DeleteMin()
		if err != nil {
			t.Fatal(err)
		}
		if v != ans {
			t.Fatalf("got: %d, expect: %d", v, ans)
		}
	}
	_, err := b.DeleteMin()
	if err == nil {
		t.Fatal("should report error when b is empty")
	}
	if b.n != 0 {
		t.Fatal("b.n should be 0")
	}
}

func TestBinomialHeap_Min(t *testing.T) {
	b := BinomialHeap{}

	_, err := b.Min()
	if err == nil {
		t.Fatal("should report error when b is empty")
	}

	for _, v := range [][2]int{{5, 5}, {2, 2}, {4, 2}, {3, 2}, {1, 1}, {6, 1}} {
		x, a := v[0], v[1]
		b.Insert(x)
		if y, err := b.Min(); y != a {
			if err != nil {
				t.Fatal(err)
			}
			t.Fatalf("got: %d, expect: %d", y, a)
		}
	}
}

func TestBinomialHeap_Meld(t *testing.T) {
	var b1, b2 BinomialHeap

	// b1 == b2 == empty
	err := b1.Meld(&b2)
	if err != nil {
		t.Fatal(err)
	}
	if !b1.Empty() || !b2.Empty() {
		t.Fatal("both b1 and b2 should be empty")
	}

	// b1 != empty, b2 == empty
	b1.Insert(1)
	b1.Insert(2)
	err = b1.Meld(&b2)
	if err != nil {
		t.Fatal(err)
	}
	if !b2.Empty() || b2.n != 0 {
		t.Fatal("b2 should be empty")
	}
	if b1.n != 2 {
		t.Fatal("b1.n should be 2")
	}
	if y, _ := b1.Min(); y != 1 {
		t.Fatal("b1.Min() should be 1")
	}

	// b1 == empty, b2 != empty
	b1, b2 = b2, b1
	err = b1.Meld(&b2)
	if err != nil {
		t.Fatal(err)
	}
	if !b2.Empty() || b2.n != 0 {
		t.Fatal("b2 should be empty")
	}
	if b1.n != 2 {
		t.Fatal("b1.n should be 2")
	}
	if y, _ := b1.Min(); y != 1 {
		t.Fatal("b1.Min() should be 1")
	}
}

func TestBinomialHeap_Meld2(t *testing.T) {
	var b1, b2 BinomialHeap
	for _, v := range []int{5, 2, 7, 6, 9} {
		b1.Insert(v)
	}
	for _, v := range []int{8, 3, 4, 1, 10} {
		b2.Insert(v)
	}

	err := b1.Meld(&b2)
	if err != nil {
		t.Fatal(err)
	}

	if !b2.Empty() || b2.n != 0 {
		t.Fatal("b2 should be empty")
	}

	if b1.n != 10 {
		t.Fatal("b1.n should be 10, got", b1.n)
	}

	if y, _ := b1.Min(); y != 1 {
		t.Fatalf("got: %d, expect: 1", y)
	}

	for ans := 1; !b1.Empty(); ans++ {
		v, _ := b1.DeleteMin()
		if v != ans {
			t.Fatalf("got %d, expect %d", v, ans)
		}
	}

	if !b1.Empty() {
		t.Fatal("b1 should be empty after delete elements")
	}
}

func TestBinomialHeap_Meld3(t *testing.T) {
	b1 := BinomialHeap{}
	var m MeldablePQ

	err := b1.Meld(m)
	if err == nil {
		t.Fatal("should report error when other is not Binomial Heap")
	}
}
