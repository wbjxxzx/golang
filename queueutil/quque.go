package queueutil

const MAXSIZE = 256

type Queue struct {
	data [MAXSIZE]interface{}
	head int
	rear int
}
type QueueError string

func (e QueueError) Error() string {
	return string(e)
}
func New() *Queue {
	q := &Queue{head: 0, rear: 0}
	return q
}
func (q *Queue) Len() int {
	len := (q.rear - q.head + MAXSIZE) % MAXSIZE
	return len

}
func (q *Queue) Put(i interface{}) error {
	if (q.rear+1)%MAXSIZE == q.head {
		return QueueError("Queue is full")
	}
	q.data[q.rear] = i
	q.rear = (q.rear + 1) % MAXSIZE
	return nil
}
func (q *Queue) Get() (interface{}, error) {
	if q.head == q.rear {
		return nil, QueueError("Queue is empty")
	}
	i := q.data[q.head]
	q.head = (q.head + 1) % MAXSIZE
	return i, nil
}
