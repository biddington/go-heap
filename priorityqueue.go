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

func parent(index int) int {
	return (index - 1) / 3
}

func (queue *PriorityQueue) swap(parent, child int) {
	(*queue)[parent], (*queue)[child] = (*queue)[child], (*queue)[parent]
}

// progressively swaps element at index with its parent
// until either element is in root position or it encounters
// a parent with a higher priority
func (pq PriorityQueue) bubbleUp(index int) {

	child := index

	for {
		parent := parent(child)

		for {
			if pq[parent].priority < pq[child].priority {
				pq.swap(parent, child)
			}
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
