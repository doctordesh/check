package check

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"testing"
)

func init() {
	os.Setenv("TERM", "dumb")
}

func TestColor(t *testing.T) {
	os.Setenv("TERM", "")
	Equals(t, true, showColors())

	os.Setenv("TERM", "dumb")
	Equals(t, false, showColors())

	os.Setenv("TERM", "")
	os.Setenv("NO_COLOR", "lkjsdf")

	Equals(t, false, showColors())
}

type mock struct{}

func (self mock) FailNow() {}

// TestMain
func TestMain(m *testing.M) {
	noColors() // turn of color codes for tests
	os.Exit(m.Run())
}

func TestOK(t *testing.T) {
	bak := out
	out = &bytes.Buffer{}
	defer func() { out = bak }()

	_, _, line, _ := runtime.Caller(0)
	OK(mock{}, fmt.Errorf("my error"))
	Equals(t, fmt.Sprintf("check_test.go:%d: unexpected error: my error\n\n", line+1), out.(*bytes.Buffer).String())

	out = &bytes.Buffer{}
	_, _, line, _ = runtime.Caller(0)
	OKWithMessage(mock{}, fmt.Errorf("my error"), "iteration %d", 5)
	Equals(t, fmt.Sprintf("check_test.go:%d: unexpected error: my error - iteration 5\n\n", line+1), out.(*bytes.Buffer).String())
}

func TestNotOK(t *testing.T) {
	bak := out
	out = &bytes.Buffer{}
	defer func() { out = bak }()

	_, _, line, _ := runtime.Caller(0)
	NotOK(mock{}, nil)
	Equals(t, fmt.Sprintf("check_test.go:%d: expected error but got nil\n\n", line+1), out.(*bytes.Buffer).String())
	out = &bytes.Buffer{}
	_, _, line, _ = runtime.Caller(0)
	NotOKWithMessage(mock{}, nil, "iteration %d", 5)
	Equals(t, fmt.Sprintf("check_test.go:%d: expected error but got nil - iteration 5\n\n", line+1), out.(*bytes.Buffer).String())
}

func TestAssert(t *testing.T) {
	bak := out
	out = &bytes.Buffer{}
	defer func() { out = bak }()

	_, _, line, _ := runtime.Caller(0)
	Assert(mock{}, false)
	Equals(t, fmt.Sprintf("check_test.go:%d: condition not met\n\n", line+1), out.(*bytes.Buffer).String())

	out = &bytes.Buffer{}
	_, _, line, _ = runtime.Caller(0)
	AssertWithMessage(mock{}, false, "iteration %d", 5)
	Equals(t, fmt.Sprintf("check_test.go:%d: condition not met - iteration 5\n\n", line+1), out.(*bytes.Buffer).String())
}

func TestEquals(t *testing.T) {
	bak := out
	out = &bytes.Buffer{}
	defer func() { out = bak }()

	_, _, line, _ := runtime.Caller(0)
	Equals(mock{}, 5, 20)

	act := out.(*bytes.Buffer).String()
	exp := fmt.Sprintf("check_test.go:%d: \n\texp: 5\n\tgot: 20\n\n", line+1)
	if exp != act {
		fmt.Printf("\texp: %#v\ngot: %#v\n\n", exp, act)
		t.FailNow()
	}

	out = &bytes.Buffer{}
	_, _, line, _ = runtime.Caller(0)
	EqualsWithMessage(mock{}, 5, 20, "iteration %d", 5)

	act = out.(*bytes.Buffer).String()
	exp = fmt.Sprintf("check_test.go:%d: \n\texp: 5\n\tgot: 20\n\titeration 5\n\n", line+1)
	if exp != act {
		fmt.Printf("\texp: %#v\ngot: %#v\n\n", exp, act)
		t.FailNow()
	}
}
