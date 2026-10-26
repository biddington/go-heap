# AGENTS.md – go-heap

## Structure

- **Module:** `biddington/go-heap` | **Go:** 1.27
- Single package, all files in root: `heap.go`, `heap_test.go`
- Test data in `testdata/rapid/TestBubbleUpUp/*.fail`

## Commands

```
go test ./...          # run tests (no special flags)
go test -v -run TestX  # single test
```

## API surface

| Exported | Signature |
|---|---|
| `func New() *Heap` | empty heap with capacity 100 |
| `func NewFrom([]Item) *Heap` | heapify from slice |
| `func (Heap) Insert(Item)` | insert item, bubble up |
| `func (Heap) Len() int` | length |

- `Heap []Item`, `Item { priority int; desc string }` — both fields unexported.
