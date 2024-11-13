package randq

import (
	"github.com/pkg/errors"
	"math/rand"
	"sync"
)

var ErrEmptyQueue = errors.New("queue is empty")

// RandomizedQueue represents a queue that supports random element removal
type RandomizedQueue[T any] struct {
	elements []T
	mutex    sync.Mutex
}

// New creates a new RandomizedQueue
func New[T any]() *RandomizedQueue[T] {
	return &RandomizedQueue[T]{elements: []T{}}
}

// Enqueue adds an element to the queue
func (q *RandomizedQueue[T]) Enqueue(value T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.elements = append(q.elements, value)
}

// BatchEnqueue adds multiple elements to the queue
func (q *RandomizedQueue[T]) BatchEnqueue(values []T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.elements = append(q.elements, values...)
}

// Dequeue removes a random element from the queue and returns it
func (q *RandomizedQueue[T]) Dequeue() (T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	var zeroValue T // default value in case of empty queue
	if len(q.elements) == 0 {
		return zeroValue, ErrEmptyQueue
	}

	// Pick a random index
	randIndex := rand.Intn(len(q.elements))
	// Swap with the last element for easy removal
	element := q.elements[randIndex]
	q.elements[randIndex] = q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]

	return element, nil
}

// BatchDequeue removes all elements from the queue and returns them
func (q *RandomizedQueue[T]) BatchDequeue() ([]T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.elements) == 0 {
		return nil, ErrEmptyQueue
	}

	elements := q.elements
	q.elements = []T{}

	return elements, nil
}

// Size returns the number of elements in the queue
func (q *RandomizedQueue[T]) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.elements)
}

// IsEmpty checks if the queue is empty
func (q *RandomizedQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}
