package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("should not be able to add a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < len(q.items); i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order: %v, want %v", item, i)
		}
	}
	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("should not be any more items in queue, got: %v", item)
	}
}
