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

// Calculates the would-be parent of index
//
// Guards negative index values, or where
// index would result in a 0 numerator, and
// returns 0
func parent(index int) int {
	if index < 2 {
		return 0
	}

	return (index - 1) / 3
}

func (queue *PriorityQueue) swap(parent, child int) {
	(*queue)[parent], (*queue)[child] = (*queue)[child], (*queue)[parent]
}

// progressively swaps element at index with its parent
// until either element is in root position or it encounters
// a parent with a higher priority
func (pq PriorityQueue) bubbleUp(index int) {

	const ROOT = 0
	child := index

	for {
		// We've reached the top of the heap
		if child == ROOT {
			break
		}

		parent := parent(child)

		if pq[parent].priority < pq[child].priority {
			pq.swap(parent, child)
			// Our bubbling value now lives at parent which we
			// copy to child for the next loop
			child = parent
		} else {
			// Heap properties have settled so we can exit
			// the loop and function
			break
		}

	}
}

func pushDown() {}

func heapify(xs []Item) []Item {
	return xs
}

/*
 * + enqueue
 * + dequeue
 * - heapify
 * - pushDown
 * - bubbleUp
 */
func (PriorityQueue) Enqueue() {

}

func Dequeue() {

}
