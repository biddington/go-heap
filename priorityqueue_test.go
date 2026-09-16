package priorityqueue

import (
	"testing"
)

func TestNewIsEmpty(t *testing.T) {
	pq := New()
	len := pq.Len()

	if len != 0 {
		t.Errorf("Expected length 0; got %d", len)
	}
}

func TestBubbleUp(t *testing.T) {

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
