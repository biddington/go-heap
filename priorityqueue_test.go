package priorityqueue

import (
	"slices"
	"testing"

	"pgregory.net/rapid"
)

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
		input    PriorityQueue
		expected PriorityQueue
	}{
		{
			input:    []Item{{priority: 6, desc: "A"}, {priority: 4, desc: "B"}},
			expected: []Item{{priority: 4, desc: "B"}, {priority: 6, desc: "A"}},
		},
	}

	for _, tt := range tests {
		t.Run("bubble-up", func(t *testing.T) {
			tt.input.swap(0, 1)

			isEqual := slices.EqualFunc(tt.input, tt.expected, func(x, y Item) bool {
				return x.priority == y.priority
			})

			if !isEqual {
				t.Errorf("swap(input) = expected")
			}
		})
	}

}

func TestBubbleUpUp(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// your test logic
	})
}

func TestBubbleUpBubbleUp(t *testing.T) {}
