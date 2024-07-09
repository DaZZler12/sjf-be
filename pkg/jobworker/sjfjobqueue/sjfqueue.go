package jobqueue

import (
	"container/heap"
	"sync"

	"github.com/DaZZler12/sjf-be/pkg/entities/sjf/model"
)

type JobQueueForSJF interface {
	Len() int
	Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
	Less(i, j int) bool
	Remove(i int)
	Peek() interface{}
	Get(i int) interface{}
}

type QueueForSJF struct {
	items []*model.SJF
	lock  sync.Mutex
}

// Len returns the length of the queue
func (q *QueueForSJF) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items)
}

// Less returns true if the duration of the job at index i is less than the duration of the job at index j
// Compare the duration of the jobs at index i and j
// to determine which job has a shorter duration (less time) making it the higher priority
// i.e. priority is based on the duration of the job
func (q *QueueForSJF) Less(i, j int) bool {
	return q.items[i].Duration < q.items[j].Duration
}

// Swap swaps the jobs at index i and j
func (q *QueueForSJF) Swap(i, j int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

// Push pushes a job into the queue
func (q *QueueForSJF) Push(x interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	item := x.(*model.SJF)
	q.items = append(q.items, item)
}

// Pop pops a job from the queue
func (q *QueueForSJF) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	n := len(q.items)
	if n == 0 {
		return nil // return nil if the queue is empty
	}
	item := q.items[n-1]
	q.items = q.items[0 : n-1]
	return item
}

// Remove removes a job from the queue
func (q *QueueForSJF) Remove(i int) {
	// give me the prodcuiton level code..
	q.lock.Lock()
	defer q.lock.Unlock()
	n := len(q.items)
	if i >= n {
		return // return if the index is out of bounds
	}
	q.items = append(q.items[:i], q.items[i+1:]...)
}

// Peek returns the job at the front of the queue
func (q *QueueForSJF) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

// Get returns the job at index i
func (q *QueueForSJF) Get(i int) interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if i >= len(q.items) {
		return nil
	}
	return q.items[i]
}

// Init initializes the queue
func (q *QueueForSJF) InitSJFQueue() {
	heap.Init(q)
}
