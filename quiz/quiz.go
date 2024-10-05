package quiz

import (
	"errors"
)

type Stack struct {
	size int
	st   []int
}

func (s *Stack) Push(val int) {
	s.st = append(s.st, val)
	s.size++
}

func (s *Stack) Peek() (int, error) {
	if s.Empty() {
		return -1, errors.New("stack is empty")
	}
	return s.st[len(s.st)-1], nil
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return -1, errors.New("stack is empty")
	}

	val, _ := s.Peek()
	s.st = s.st[:len(s.st)-1]
	s.size--
	return val, nil
}
func (s *Stack) Empty() bool {
	return s.size == 0
}
func (s *Stack) Size() int {
	return s.size
}

func NewStack() *Stack {
	return &Stack{size: 0, st: []int{}}
}

type Queue struct {
	last  int
	queue []int
	front int
}

// 3 ,4 ,4,5,6,7,8,9
/*
3 (f=0,l=0)
4  f=0,l=1
4  f=0 ,l=2

*/
func (q *Queue) Add(val int) {
	if q.front == -1 && q.last == -1 {
		q.front++
		q.last++
		q.queue = append(q.queue, val)
	} else {
		q.queue = append(q.queue, val)
		q.last++
	}

}
func (q *Queue) Poll() (int, error) {
	if q.Empty() {
		return -1, errors.New("Queue is Empry")
	}
	val, _ := q.Peek()
	q.front++
	if q.front > q.last {
		q.last = -1
		q.front = -1
	}
	return val, nil

}

func (q *Queue) Peek() (int, error) {
	if q.Empty() {
		return -1, errors.New("Queue is Empty")
	}
	return q.queue[q.front], nil
}

func (q *Queue) Empty() bool {
	if q.front == -1 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	if q.Empty() {
		return 0
	}
	return q.last - q.front + 1
}

func NewQueue() *Queue {
	return &Queue{last: -1, front: -1, queue: []int{}}
}
