package greetings

import (
	"testing"
	"regexp"
)

// use the command go test or go test -v for verbose

// Testing the greetings.Hello function

// Test to Check for a valid return value
func TestHelloName(t *testing.T) {
	// the parameter is a pointer to the testing package's testing.T type

	name := "Parth"
	want := regexp.MustCompile(`\b`+name+`\b`)

	msg, err := Hello("Parth")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Parth") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

	// %#q → formats a string as a valid Go literal (safer/more precise than %q).
	// \b → regex word boundary; ensures you’re matching the whole word "Parth", not just a substring.
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want ""`, msg, err)
	}
}

// PS C:\VIT\Sem7\GoLang\creatingModules\greetings> go test
// PASS
// ok      example.com/greetings   0.217s

// PS C:\VIT\Sem7\GoLang\creatingModules\greetings> go test -v
// === RUN   TestHelloName
// --- PASS: TestHelloName (0.00s)
// === RUN   TestHelloEmpty
// --- PASS: TestHelloEmpty (0.00s)
// PASS
// ok      example.com/greetings   0.355s