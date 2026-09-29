package priorityqueue

type Item struct {
	priority int
	desc     string
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func NewFrom(items []Item) *PriorityQueue {
	heapified := heapify(items)

	pq := PriorityQueue(heapified)

	return &pq
}

func New() *PriorityQueue {
	pq := PriorityQueue(make([]Item, 0, 100))

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
func (queue PriorityQueue) swap(parent, child int) {
	queue[parent], queue[child] = queue[child], queue[parent]
}

// Progressively swaps element at index with its parent
// until heap invariants are satisified or element
// reaches root position.
//
// Pre: pq must be a valid heap
func (pq PriorityQueue) bubbleUp(index int) {

	const ROOT = 0
	child := index

	for {
		// Element is now at the top of the heap
		if child == ROOT {
			break
		}

		parent := parent(child)

		if pq[parent].priority < pq[child].priority {
			pq.swap(parent, child)
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

// Reinstates heap invariants with respect to
// an element at index and its children
//
// E.g. we've popped the min/max element at
// the root and replaced it with the element
// in tail position which will be of lower
// priority with respect to its (new) children
func pushDown(index int) {
	// We want to keep pushing element at index
	// down as long as it has a lower priority
	// than _any_ of its children
}

func heapify(xs []Item) PriorityQueue {
	return PriorityQueue(xs)
}

func (pq PriorityQueue) Enqueue(item Item) {
	pq = append(pq, item)
	pq.bubbleUp(pq.Len() - 1)
}

func (pq PriorityQueue) Dequeue() {

}
