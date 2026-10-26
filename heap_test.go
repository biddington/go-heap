package heap

import (
	"slices"
	"testing"
)

func AreHeapsEqual(input Heap, expected Heap) bool {
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
		input    Heap
		index    int
		expected Heap
	}{
		{
			name:     "bubbleUp on an empty heap returns an empty heap",
			input:    []Item{},
			index:    0,
			expected: []Item{},
		},
		{
			name:     "bubbleUp on a heap of length 1 returns the heap",
			input:    []Item{{priority: 1, desc: "lone task"}},
			index:    0,
			expected: []Item{{priority: 1, desc: "lone task"}},
		},
		{
			name:     "bubbleUp on valid heap returns the heap",
			input:    []Item{{priority: 9, desc: "root"}, {priority: 4, desc: "child"}},
			index:    1,
			expected: []Item{{priority: 9, desc: "root"}, {priority: 4, desc: "child"}},
		},
		{
			name: "bubbleUp on a heap with a single element out of place returns a valid heap",
			input: []Item{
				{priority: 19, desc: "R"},
				{priority: 11, desc: "A"},
				{priority: 12, desc: "B"},
				{priority: 17, desc: "C"},
				{priority: 4, desc: "AA"},
				{priority: 8, desc: "AB"},
				{priority: 54, desc: "AC"}, // Element AC has the max priority and should be at the root of the heap
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
			tt.input.bubbleUp(tt.index)

			areEqual := AreHeapsEqual(tt.input, tt.expected)

			if !areEqual {
				t.Errorf("bubbleUp(%d): got %+v; want %+v", tt.index, tt.input, tt.expected)
			}
		})
	}

}

func TestPushDown(t *testing.T) {
	tests := []struct {
		desc     string
		input    Heap
		index    int
		expected Heap
	}{
		{
			desc:     "pushDown on an empty heap is a no-op",
			input:    []Item{},
			expected: []Item{},
			index:    0,
		},
		{
			desc:     "pushDown on a heap of a single value is a no-op",
			input:    []Item{{priority: 1, desc: "Half-life 3"}},
			expected: []Item{{priority: 1, desc: "Half-life 3"}},
			index:    0,
		},
		// {
		// 	desc:     "pushDown should restore heap invariants when a single node is in violation",
		// 	input:    []Item{{priority: 16, desc: "16"}, {priority: 23, desc: "23"}, {priority: 27, desc: "27"}, {priority: 45, desc: "45"}},
		// 	expected: []Item{{priority: 45, desc: "45"}, {priority: 23, desc: "23"}, {priority: 27, desc: "27"}, {priority: 16, desc: "16"}},
		// 	index:    0,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.input.pushDown(tt.index)

			areEqual := AreHeapsEqual(tt.input, tt.expected)

			// if !isMaxHeap(tt.input) {
			// 	t.Errorf("heap invariant violation: got %+v; want %+v", tt.index, tt.expected)
			// }

			if !areEqual {
				t.Errorf("pushDown(%d): got %+v; want %+v", tt.index, tt.index, tt.expected)
			}
		})
	}
}

func TestHeapify(t *testing.T) {
	tests := []struct {
		name   string
		input  []Item
		output Heap
	}{
		{
			name:   "Heapify - Single element",
			input:  []Item{{priority: 9, desc: "Low"}},
			output: Heap{{priority: 9, desc: "Low"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := heapify(tt.input)

			if !AreHeapsEqual(got, tt.output) {
				t.Errorf("Got %v, want %v", got, tt.output)
			}
		})
	}
}
