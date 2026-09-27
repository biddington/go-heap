package priorityqueue

import (
	"slices"
	"testing"

	"pgregory.net/rapid"
)

func QueuesEqual(input PriorityQueue, expected PriorityQueue) bool {
	return slices.EqualFunc(input, expected, func(x, y Item) bool {
		return x.priority == y.priority && x.desc == y.desc
	})
}

func isMaxHeap(s []Item) {

}

func TestNewIsEmpty(t *testing.T) {
	pq := New()
	len := pq.Len()

	if len != 0 {
		t.Errorf("Expected length 0; got %d", len)
	}
}

func TestParent(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 0},
		{input: 1, expected: 0},
		{input: 3, expected: 0},
		{input: 5, expected: 1},
		{input: 8, expected: 2},
		{input: 11, expected: 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := parent(tt.input)

			if got != tt.expected {
				t.Errorf("parent(%d) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestBubbleUp(t *testing.T) {
	tests := []struct {
		name     string
		pq       PriorityQueue
		index    int
		expected PriorityQueue
	}{
		{
			name:     "bubbleUp on an empty queue",
			pq:       []Item{},
			index:    0,
			expected: []Item{},
		},
		{
			name:     "bubbleUp on a queue of 1",
			pq:       []Item{{priority: 1, desc: "lone task"}},
			index:    0,
			expected: []Item{{priority: 1, desc: "lone task"}},
		},
		{
			name:     "bubbleUp on an invalid queue",
			pq:       []Item{{priority: 2, desc: "low"}, {priority: 7, desc: "high"}},
			index:    1,
			expected: []Item{{priority: 7, desc: "high"}, {priority: 2, desc: "low"}},
		},
		{
			name:     "bubbleUp on a valid queue",
			pq:       []Item{{priority: 9, desc: "root"}, {priority: 4, desc: "child"}},
			index:    1,
			expected: []Item{{priority: 9, desc: "root"}, {priority: 4, desc: "child"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.bubbleUp(tt.index)

			areEqual := QueuesEqual(tt.pq, tt.expected)

			if !areEqual {
				t.Errorf("bubbleUp(%d): got %+v; want %+v", tt.index, tt.pq, tt.expected)
			}
		})
	}

}

func TestHeapify(t *testing.T) {
	tests := []struct {
		name   string
		input  []Item
		output PriorityQueue
	}{
		{
			name:   "Heapify - Single element",
			input:  []Item{{priority: 9, desc: "Low"}},
			output: PriorityQueue{{priority: 9, desc: "Low"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := heapify(tt.input)

			if !QueuesEqual(got, tt.output) {
				t.Errorf("Got %v, want %v", got, tt.output)
			}
		})
	}
}

func BubbleUpUp(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		itemGen := rapid.Custom(func(t *rapid.T) Item {
			return Item{
				priority: rapid.IntRange(1, 50).Draw(t, "priority"),
				desc:     "",
			}
		})

		items := make([]Item, 5)
		for i := range items {
			items[i] = itemGen.Draw(t, "")
		}

		pq := PriorityQueue(items)
		parentIndex := parent(4)

		if pq[parentIndex].priority < pq[4].priority {
			pq.swap(parentIndex, 4)
		}

		pq.bubbleUp(4)

		n := pq.Len()
		for i := 0; i < n; i++ {
			for j := 3*i + 1; j <= 3*i+3 && j < n; j++ {
				if pq[j].priority > pq[i].priority {
					t.Fatalf("heap violated: node[%d](%d) < child[%d](%d)", i, pq[i].priority, j, pq[j].priority)
				}
			}
		}
	})
}
