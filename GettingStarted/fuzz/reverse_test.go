package main

import (
	"testing"
	"unicode/utf8"
)

// Unit test
func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testCases {
		rev, err := Reverse(tc.in)
		if err != nil {
			return
		}
		if tc.want != rev {
			t.Errorf("Reverse: %q, want: %q", rev, tc.want)
		}
	}
}

// Fuzz test
func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc) //Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		// t.Logf("Numbers of runes: orig=%d, rev=%d, doubleRev=%d",
		// 	utf8.RuneCountInString(orig),
		// 	utf8.RuneCountInString(rev),
		// 	utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
