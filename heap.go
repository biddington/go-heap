package heap

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

// Returns the index of the highest priority child or -1 if index has no children
//
// Pre: index is within the bounds of the heap
func (heap Heap) highestPriorityChildIndex(index int) int {

	// -1 is chosen as the sentinal value because it's outside the range
	// of values we consider legitimate indexes
	const NO_CHILDREN = -1

	// Although this captures the condition that an index must be within the bounds
	// of the heap, we collapse the meaning of the return value to that of the
	// NO_CHILDREN return value becuase to callers of this function it's functionally
	// equivalent to saying "no useful value"
	if index < 0 || index >= heap.Len() {
		return NO_CHILDREN
	}

	// Calculated first and last child of index
	firstChild := (index * BRANCHING_FACTOR) + 1
	lastChild := (index * BRANCHING_FACTOR) + BRANCHING_FACTOR

	// Functions as a signal in the loop below to store the index to child
	// in preparation for subsequent compares
	//
	// Is also the sentinal value for indicating no children
	child := NO_CHILDREN

	// firstChild and lastChild are _calculated_ indexes. As such the loop bounds guard against invalid
	// access where a calculated index doesn't actually exist within the heap
	for iter := firstChild; iter <= lastChild && iter < heap.Len(); iter++ {
		// In order to keep the loop mechanics straight forward the initial condition
		// is detected based on the NO_CHILDREN value. This signals to store iter which
		// will be the first child. Subsequent iterations will use the second branch because
		// we now have another index to compare against.
		//
		// This means we don't need to have the first child stored outside the loop
		// and the loop go through awkward BRANCHING_FACTOR - 1 iterations
		if child == NO_CHILDREN {
			child = iter
		} else if heap[iter].priority > heap[child].priority {
			child = iter
		}
	}

	return child
}

// Reinstates heap invariants with respect to an element at index and its children
//
// Assumes index is the only node that is out of place i.e. that is, violates the heaps
// parent > child invariant
//
// E.g. we've popped the min/max element at the root and replaced it with the element
// in tail position which will be of lower priority with respect to its (new) children
func (heap Heap) pushDown(index int) {

	parentIndex := index

	// An empty heap or heap of a single element is by definition, valid
	if heap.Len() < 2 {
		return
	}

	// We want to keep pushing down element at index as long as it has a lower priority
	// than any of its children. If we're at last position, i.e. there's no children so
	// we stop pushing down
	for {
		childIndex := heap.highestPriorityChildIndex(parentIndex)

		// childIndex refers to a leaf node. The heaps parent-child invariant has been
		// reinstated so we return
		if childIndex == -1 {
			return
		}

		// When the parent has a lower priority than it's child we swap them in order to
		// reinstate the heaps parent-child invariant. Our focus is now on childIndex and
		// we copy it to parentIndex in preparation for the next iteration
		if heap[parentIndex].priority < heap[childIndex].priority {
			heap.swap(parentIndex, childIndex)
			parentIndex = childIndex
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
