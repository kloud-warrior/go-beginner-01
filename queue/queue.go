package types

import (
	"fmt"
)

var (
	ErrQueueIsEmpty = fmt.Errorf("queue is empty")
)

// New ...
func New() *Queue {
	return &Queue{}
}

func NewPriorityQueue() *Queue {
	return &Queue{}
}

// Queue ...
type Queue struct {
}

func (q Queue) IsEmpty() bool {
	return false
}

func (q Queue) Size() int {
	return -1
}

func (q *Queue) Push(num int) {
}

func (q Queue) Peek() (int, error) {
	return -1, nil
}

func (q *Queue) Pop() (int, error) {
	return -1, nil
}
