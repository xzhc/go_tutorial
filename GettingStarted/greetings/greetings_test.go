package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "xzh"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(name) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TesingHelloEmpty calls greeting.Hello with an empty string, checking for an error
func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello(name) = %q, %v, want "", error`, msg, err)
	}

}
