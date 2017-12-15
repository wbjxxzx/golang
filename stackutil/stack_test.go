package stackutil

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if len(s.data) != MAXSIZE {
		t.Errorf("len(New()) == %q, want %q", len(s.data), MAXSIZE)
	}
}

var s *Stack = &Stack{top: -1}
var n int = 1
var cases = []struct {
	in, want interface{}
}{
	{1, 1},
	{&n, &n},
	{"abc", "abc"},
	{struct{ int }{1}, struct{ int }{1}},
}

func TestPush(t *testing.T) {
	for _, v := range cases {
		err := s.Push(v.in)
		if err != nil {
			t.Error(err)
		}
	}
	if s.top != len(cases)-1 {
		t.Errorf("s.top = %q, want %q", s.top, len(cases)-1)
	}
}
func TestPop(t *testing.T) {
	data := make([]interface{}, 0)
	for _, v := range cases {
		data = append(data[:0], append([]interface{}{v.want}, data[0:]...)...)
		fmt.Printf("%#v\n", data)
	}
	for _, v := range data {
		i, err := s.Pop()
		if i != v {
			t.Errorf("s.Pop() = %q, want %q", i, v)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
