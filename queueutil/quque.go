package queueutil

const MAXSIZE = 256

type Queue struct {
	data [MAXSIZE]interface{}
	font int
	rear int
}
type QueueError string

func (e QueueError) Error() string {
	return string(e)
}
func New() *Queue {
	q := &Queue{font: -1, rear: -1}
	return q
}
