package queueutil

import "testing"

func TestNew(t *testing.T) {
	q := New()
	if q == nil {
		t.Errorf("queueutil.New is not success")
	}
}

var q *Queue = &Queue{head: 0, rear: 0}

var cases = []struct {
	in, want interface{}
}{
	{1, 1},
	{"abc", "abc"},
	{[3]int{1, 2, 3}, [3]int{1, 2, 3}},
	{struct{ int }{1}, struct{ int }{1}},
}

func TestLen(t *testing.T) {
	n := 0
	if q.Len() != 0 {
		t.Errorf("Len(q) should by 0, but get %q", n)
	}
}
func TestPut(t *testing.T) {
	for _, v := range cases {
		err := q.Put(v.in)
		if err != nil {
			t.Errorf("Put(i) return error %q", err)
		}
	}
}
func TestGet(t *testing.T) {
	for _, v := range cases {
		i, err := q.Get()
		if i != v.want {
			t.Errorf("Get() = %q, want %q, ", i, v.want)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
