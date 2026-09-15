package priorityqueue

import (
	"testing"
)

func TestNew(t *testing.T) {
	pq := New()
	len := pq.Len()

	if len != 0 {
		t.Errorf("Expected length 0; got %d", len)
	}
}
