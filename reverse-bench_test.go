package stringutil

import "testing"

func BenchmarkReverse(b *testing.B) {
	posCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for i := 0; i < b.N; i++ {
		for _, c := range posCases {
			got := Reverse(c.in)
			if got != c.want {
				b.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}
