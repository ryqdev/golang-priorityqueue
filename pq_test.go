package golang_priorityqueue_test

import (
	"github.com/stretchr/testify/assert"
	"golang_priorityqueue"
	"testing"
)

// TestPriorityQueueInt is the unit test for a PriorityQueue with integers
func TestPriorityQueueInt(t *testing.T) {
	rule := func(a, b int) bool { return a < b }

	// Create a priority queue of integers
	pq := golang_priorityqueue.PriorityQueue[int]{
		Array: []int{},
		Rule:  rule,
	}

	// Test Empty
	assert.True(t, pq.Empty(), "PriorityQueue should be empty initially")

	// Test Push and Top
	pq.Push(10)
	assert.False(t, pq.Empty(), "PriorityQueue should not be empty after push")
	assert.Equal(t, 10, pq.Top(), "Top should return the only element")

	// Push more elements
	pq.Push(5)
	pq.Push(15)
	assert.Equal(t, 5, pq.Top(), "Top should return the smallest element after multiple pushes")

	// Test Pop and heap property
	pq.Pop()
	assert.Equal(t, 10, pq.Top(), "Top should return the next smallest element after pop")
	pq.Pop()
	assert.Equal(t, 15, pq.Top(), "Top should return the last element after pop")

	// Pop the last element
	pq.Pop()
	assert.True(t, pq.Empty(), "PriorityQueue should be empty after popping all elements")
}

// TestPriorityQueueString is the unit test for a PriorityQueue with strings
func TestPriorityQueueString(t *testing.T) {
	// Max-heap rule: lexicographically larger strings have higher priority
	rule := func(a, b string) bool { return a > b }

	// Create a priority queue of strings
	pq := golang_priorityqueue.PriorityQueue[string]{
		Array: []string{},
		Rule:  rule,
	}

	// Test Empty
	assert.True(t, pq.Empty(), "PriorityQueue should be empty initially")

	// Test Push and Top
	pq.Push("apple")
	assert.False(t, pq.Empty(), "PriorityQueue should not be empty after push")
	assert.Equal(t, "apple", pq.Top(), "Top should return the only element")

	// Push more elements
	pq.Push("banana")
	pq.Push("cherry")
	assert.Equal(t, "cherry", pq.Top(), "Top should return the lexicographically largest element after multiple pushes")

	// Test Pop and heap property
	pq.Pop()
	assert.Equal(t, "banana", pq.Top(), "Top should return the next largest element after pop")
	pq.Pop()
	assert.Equal(t, "apple", pq.Top(), "Top should return the last element after pop")

	// Pop the last element
	pq.Pop()
	assert.True(t, pq.Empty(), "PriorityQueue should be empty after popping all elements")
}

// TestPriorityQueueHeapify tests the heapify operation
func TestPriorityQueueHeapify(t *testing.T) {
	rule := func(a, b int) bool { return a < b }

	// Create a priority queue with unordered elements
	pq := golang_priorityqueue.PriorityQueue[int]{
		Array: []int{10, 20, 5, 7, 2},
		Rule:  rule,
	}

	// Call Heapify to reorder elements into a valid heap
	pq.Heapify()

	// Test if the top is the smallest element after heapify
	assert.Equal(t, 2, pq.Top(), "Top should return the smallest element after heapify")
}
