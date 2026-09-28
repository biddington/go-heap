package priorityqueue

import (
	"slices"
	"testing"
)

func QueuesEqual(input PriorityQueue, expected PriorityQueue) bool {
	return slices.EqualFunc(input, expected, func(x, y Item) bool {
		return x.priority == y.priority && x.desc == y.desc
	})
}

func isMaxHeap(s []Item) bool {
	for i := range len(s) {
		firstChild := 3*i + 1
		if firstChild >= len(s) {
			break
		}
		for j := firstChild; j < len(s) && j < firstChild+3; j++ {
			if s[j].priority > s[i].priority {
				return false
			}
		}
	}
	return true
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
			name:     "bubbleUp a queue of 1",
			pq:       []Item{{priority: 1, desc: "lone task"}},
			index:    0,
			expected: []Item{{priority: 1, desc: "lone task"}},
		},
		{
			name:     "bubbleUp an invalid queue",
			pq:       []Item{{priority: 2, desc: "low"}, {priority: 7, desc: "high"}},
			index:    1,
			expected: []Item{{priority: 7, desc: "high"}, {priority: 2, desc: "low"}},
		},
		{
			name:     "bubbleUp a valid queue",
			pq:       []Item{{priority: 9, desc: "root"}, {priority: 4, desc: "child"}},
			index:    1,
			expected: []Item{{priority: 9, desc: "root"}, {priority: 4, desc: "child"}},
		},
		{
			name: "bubbleUp a valid queue",
			pq: []Item{
				{priority: 19, desc: "R"},
				{priority: 11, desc: "A"},
				{priority: 12, desc: "B"},
				{priority: 17, desc: "C"},
				{priority: 4, desc: "AA"},
				{priority: 8, desc: "AB"},
				{priority: 54, desc: "AC"},
			},
			index: 6,
			expected: []Item{
				{priority: 54, desc: "AC"},
				{priority: 19, desc: "R"},
				{priority: 12, desc: "B"},
				{priority: 17, desc: "C"},
				{priority: 4, desc: "AA"},
				{priority: 8, desc: "AB"},
				{priority: 11, desc: "A"},
			},
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
