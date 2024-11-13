package main

import "fmt"

type Tree struct {
	X      int
	father *Tree
	L      *Tree
	R      *Tree
}

func (t Tree) printHeap() {
	fmt.Print(t.X, ":{")
	if t.L != nil {
		t.L.printHeap()
		fmt.Print(",")
		if t.R != nil {
			t.R.printHeap()
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("}")
}

func (t *Tree) heapify(A []int) (Anew []int) {
	N := len(A) - 1
	Anew = A
	if N >= 0 {
		t.L = &Tree{X: Anew[N], father: t, L: nil, R: nil}
		t.L.siftUp()
		Anew = Anew[:N]
		N--
		if N >= 1 {
			t.R = &Tree{X: Anew[N], father: t, L: nil, R: nil}
			t.R.siftUp()
			Anew = Anew[:N]
			N--
		}
		if N >= 1 {
			Anew = t.L.heapify(Anew)
		}
		if N >= 1 {
			Anew = t.R.heapify(Anew)
		}
	}
	return Anew
}

func (t *Tree) siftUp() {
	for t.father != nil {
		if t.X > t.father.X {
			app := t.X
			t.X = t.father.X
			t.father.X = app
		}
		t = t.father
	}
}

//func (t *Heap) siftDown(root int) {
//	for t.heap[root].L != nil {
//		child := t.heap[root].L
//		if t.heap[root].R != nil {
//			child_right := t.heap[root].R
//			if child.X > child_right.X {
//				index := child.index
//				child.index = child_right.index
//				child_right.index = index
//				t.heap[root].R = child
//				t.heap[root].L = child_right
//			}
//		}
//		child = t.heap[root].L
//		father := t.heap[root]
//		if father.X < child.X {
//			index := child.index
//			child.index = father.index
//			father.index = index
//			t.heap[root] = *child
//			t.heap[root].L = &father
//		} else {
//			goto end
//		}
//
//	}
//end:
//}

func heapSort(A []int) {

	// 1. build the heap, get it's size
	buildMaxHeap(A)
	size := len(A)

	for size > 1 {
		// 2. repeatedly swap first and last element, decrement size of considered list
		A[0], A[size-1] = A[size-1], A[0]
		size -= 1

		// 3. siftDown on first element to put it in place
		siftDown(A[:size], 0)
	}
}

// build a Max heap from a given array.
// a max heap adheres to the properties of:
//	1. being filled left to right
//	2. no child node is greater than its parents
func buildMaxHeap(A []int) {

	// since all leaf nodes are considered valid heaps, we start from the last
	// non-leaf node, and work right to left in the array
	for i := (len(A) / 2) - 1; i >= 0; i -= 1 {
		siftDown(A, i)
	}
}

// given an index i of A, sift it downwards, should it be smaller than its children
func siftDown(A []int, i int) {
	largest := i
	l, r := getChildIndices(i)
	if l < len(A) && A[l] > A[i] {
		largest = l
	}
	if r < len(A) && A[r] > A[largest] {
		largest = r
	}
	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		siftDown(A, largest)
	}
}

func getParentIndex(childIndex int) int {
	if childIndex%2 == 0 {
		return (childIndex - 2) / 2
	} else {
		return (childIndex - 1) / 2
	}
}

// return child indices as (left-index, right-index)
func getChildIndices(parentIndex int) (int, int) {
	return (2*parentIndex + 1), (2*parentIndex + 2)
}
