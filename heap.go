package heap

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

func highestPriorityChild(index int) int {
	return 0
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

// Returns the highest priority child of element at index
//
// Pre: index is within the bounds of the heap
func (heap Heap) highestPriorityChild(index int) int {

	if index > heap.Len() || index < 0 {

	}

	return 0
}

// Reinstates heap invariants with respect to
// an element at index and its children
//
// E.g. we've popped the min/max element at
// the root and replaced it with the element
// in tail position which will be of lower
// priority with respect to its (new) children
func (heap Heap) pushDown(index int) {
	// We want to keep pushing element at index
	// down as long as it has a lower priority
	// than _any_ of its children
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
