
     ____ _____        / /_  ___  ____ _____
    / __ `/ __ \______/ __ \/ _ \/ __ `/ __ \
   / /_/ / /_/ /_____/ / / /  __/ /_/ / /_/ /
   \__, /\____/     /_/ /_/\___/\__,_/ .___/
  /____/                            /_/

A max heap (priority queue) implemented in Go with a ternary branching factor of 3, where each node can have up to three children rather than the traditional binary two.

// API
import "biddington/go-heap"

h := heap.New()                // empty heap with capacity 100
h := heap.NewFrom(items)       // build from []heap.Item
h.Insert(item)                 // insert and bubble up to maintain the heap invariant
n := h.Len()                   // number of elements currently stored


// TYPES
type Heap []Item          // slice-backed max heap
type Item struct { ... }  // priority + description (fields unexported)

heap.IndexOutOfBounds    // error returned by internal out-of-bounds checks


// BUILDING & TESTING
go test ./...
go test -v -run TestX   # single test
