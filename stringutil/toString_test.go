package stringutil

import "testing"

func TestStr(t *testing.T) {
	cases := []struct {
		i    interface{}
		want string
	}{
		{1, "1"},
		{uint(1), "1"},
		{1.0, "1.0000"},
		{float32(1.0), "1.0000"},
		{"", ""},
		{"abc", "abc"},
		{'c', "c"},
		{true, "true"},
		{struct{ int }{1}, "unknown type"},
	}
	for _, c := range cases {
		got := Str(c.i)
		if got != c.want {
			t.Errorf("Str(%q) = %q, want %q", c.i, got, c.want)
		}
	}
}
