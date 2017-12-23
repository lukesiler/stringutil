package stringutil

import "testing"

func TestReversePos(t *testing.T) {
	posCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range posCases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestReverseNeg(t *testing.T) {
	negCases := []struct {
		in, noWant string
	}{
		{"Hello, world", "dlrow, olleH"},
		{"Hello, 世界", "界世, olleH"},
		{"", " "},
		{"something", "nothing"},
		{"blah", "blah"},
	}
	for _, c := range negCases {
		got := Reverse(c.in)
		if got == c.noWant {
			t.Errorf("Reverse(%q) == %q, don't want %q", c.in, got, c.noWant)
		}
	}
}
