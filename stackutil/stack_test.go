package stackutil

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if len(s.arr) != MAXSIZE {
		t.Errorf("len(New()) == %q, want %q", len(s.arr), MAXSIZE)
	}
}
func TestPush(t *testing.T) {
	s := New()
	n := 1
	p := &n
	s.Push(p)
	s.Push(1)
	s.Push("abc")
	s.Push(struct{ int }{1})
	if s.top != 3 {
		t.Errorf("s.top = %q, want %q", s.top, 3)
	}
}
