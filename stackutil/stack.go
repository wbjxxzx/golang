package stackutil

const MAXSIZE = 256

type Stack struct {
	data [MAXSIZE]interface{}
	top  int
}
type StackError string

func (e StackError) Error() string {
	return string(e)
}
func New() *Stack {
	t := &Stack{top: -1}
	return t
}
func (s *Stack) isFull() bool {
	if s.top == MAXSIZE-1 {
		return true
	}
	return false
}
func (s *Stack) isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}
func (s *Stack) Push(i interface{}) error {
	if s.isFull() {
		return StackError("Stack is full")
	}
	s.top++
	s.data[s.top] = i
	return nil
}
func (s *Stack) Pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, StackError("Stack is empty")
	}
	i := s.data[s.top]
	s.data[s.top] = nil
	s.top--
	return i, nil
}
