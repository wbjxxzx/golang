package stackutil

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if len(s.data) != MAXSIZE {
		t.Errorf("len(New()) == %q, want %q", len(s.data), MAXSIZE)
	}
}
func TestPushPop(t *testing.T) {
	s := New()
	var i interface{}
	n := 1
	p := &n
	s.Push(p)
	s.Push(1)
	s.Push("abc")
	s.Push(struct{ int }{1})
	if s.top != 3 {
		t.Errorf("s.top = %q, want %q", s.top, 3)
	}
	i, _ = s.Pop()
	if i != struct{ int }{1} {
		t.Errorf("s.Pop() = %q, want %q", i, struct{ int }{1})
	}
	i, _ = s.Pop()
	if i != "abc" {
		t.Errorf("s.Pop() = %q, want %q", i, "abc")
	}
	i, _ = s.Pop()
	if i != 1 {
		t.Errorf("s.Pop() = %q, want %q", i, 1)
	}
	i, _ = s.Pop()
	if i != p {
		t.Errorf("s.Pop() = %q, want %q", i, p)
	}
}
