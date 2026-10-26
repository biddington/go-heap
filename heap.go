package heap

import (
	"errors"
)

// The max number of children each tree node can have
const BRANCHING_FACTOR = 3

type Item struct {
	priority int
	desc     string
}

type Heap []Item

func (heap Heap) Len() int {
	return len(heap)
}

func NewFrom(items []Item) *Heap {
	heapified := heapify(items)

	pq := Heap(heapified)

	return &pq
}

func New() *Heap {
	pq := Heap(make([]Item, 0, 100))

	return &pq
}

// Calculates the would-be parent of index as if
// the heap was a tree structure with branching
// factor 3
func parent(index int) int {
	return (index - 1) / 3
}

// Swap is essentially a thin wrapper around tuple
// re-assignment.
//
// Take the values at child and parent. Put the
// child into the parent position and parent into
// the child position, all without a temp variable
func (heap Heap) swap(parent, child int) {
	heap[parent], heap[child] = heap[child], heap[parent]
}

// Progressively swaps element at index with its parent
// until heap invariants are satisified or element
// reaches root position.
//
// Pre: pq must be a valid heap
func (heap Heap) bubbleUp(index int) {

	const ROOT = 0
	child := index

	for {
		// Element is now at the top of the heap
		if child == ROOT {
			break
		}

		parent := parent(child)

		if heap[parent].priority < heap[child].priority {
			heap.swap(parent, child)
			// After swapping, our bubbling value now lives at parent (index)
			// which we copy to child for the next loop
			child = parent
		} else {
			// Heap properties have settled so we can exit
			// the loop and function
			break
		}

	}
}

var IndexOutOfBounds = errors.New("index out of bounds")

// Returns the index of the highest priority child or -1 if index has no children
//
// Pre: index is within the bounds of the heap
func (heap Heap) highestPriorityChild(index int) (int, error) {

	// -1 is chosen as the sentinal value because it's outside the range
	// of values we consider legitimate indexes
	const NO_CHILDREN = -1

	if index < 0 || index >= heap.Len() {
		return -1, IndexOutOfBounds
	}

	// Calculated first and last child of index
	firstChild := (index * BRANCHING_FACTOR) + 1
	lastChild := (index * BRANCHING_FACTOR) + BRANCHING_FACTOR

	child := NO_CHILDREN

	// firstChild and lastChild are calculated and are not gauranteed to be
	// available on the heap. We need to guard against accessing them if
	// they're not actually children of index, by testing against heap.Len()
	for iter := firstChild; iter <= lastChild && iter < heap.Len(); iter++ {
		if child == NO_CHILDREN {
			child = iter
		} else if heap[iter].priority > heap[child].priority {
			child = iter
		}
	}

	return child, nil
}

// Reinstates heap invariants with respect to
// an element at index and its children
//
// E.g. we've popped the min/max element at
// the root and replaced it with the element
// in tail position which will be of lower
// priority with respect to its (new) children
func (heap Heap) pushDown(index int) {

	// An empty heap or heap of a single element is by definition, valid
	if heap.Len() < 2 {
		return
	}

	// We want to keep pushing element at index down as long as it has a lower priority
	// than _any_ of its children
	// If we're at last position, i.e. there's no children, we stop pushing down
	parent := heap[index]

	for {
		child, err := heap.highestPriorityChild(index)

		if err != nil {
			panic(err.Error())
		}

		if heap[child].priority > parent.priority {
			heap.swap(index, child)
		}
	}
}

func heapify(xs []Item) Heap {
	return Heap(xs)
}

func (heap Heap) Insert(item Item) {
	heap = append(heap, item)
	heap.bubbleUp(heap.Len() - 1)
}

func (heap Heap) Take() {

}
