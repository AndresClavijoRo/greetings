package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Alice"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Alice") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}
	messages, err := Hellos(names)

	if err != nil {
		t.Fatalf(`Hellos(%q) returned an error: %v`, names, err)
	}

	if len(messages) != len(names) {
		t.Fatalf(`Hellos(%q) = %d messages, want %d`, names, len(messages), len(names))
	}

	for _, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg := messages[name]
		if !want.MatchString(msg) {
			t.Errorf(`Hellos(%q) = %q, want match for %#q`, names, msg, want)
		}
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{""}

	messages, err := Hellos(names)
	if messages != nil || err == nil {
		t.Fatalf(`Hellos(%q) = %v, %v, want nil, error`, names, messages, err)
	}
}
