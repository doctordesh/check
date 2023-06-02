package check

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

var out io.Writer = os.Stdout

type TestingHarness interface {
	FailNow()
}

//
// OK
//

func OK(tb TestingHarness, err error) {
	if err == nil {
		return
	}

	p(tb, "unexpected error: %s", err.Error())
}

func OKWithMessage(tb TestingHarness, err error, msg string, v ...interface{}) {
	if err == nil {
		return
	}

	p(tb, "unexpected error: %s - %s", err.Error(), fmt.Sprintf(msg, v...))
}

//
// NotOK
//
func NotOK(tb TestingHarness, err error) {
	if err != nil {
		return
	}

	p(tb, "expected error but got nil")
}

func NotOKWithMessage(tb TestingHarness, err error, msg string, v ...interface{}) {
	if err != nil {
		return
	}

	p(tb, "expected error but got nil - %s", fmt.Sprintf(msg, v...))
}

//
// Assert
//

func Assert(tb TestingHarness, condition bool) {
	if condition {
		return
	}

	p(tb, "condition not met")
}

func AssertWithMessage(tb TestingHarness, condition bool, msg string, v ...interface{}) {
	if condition {
		return
	}

	p(tb, "condition not met - %s", fmt.Sprintf(msg, v...))
}

//
// Equals
//

func Equals(tb TestingHarness, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		return
	}

	p(tb, "\n\texp: %#v\n\tgot: %#v", exp, act)
}

func EqualsWithMessage(tb TestingHarness, exp, act interface{}, msg string, v ...interface{}) {
	if reflect.DeepEqual(exp, act) {
		return
	}

	p(tb, "\n\texp: %#v\n\tgot: %#v\n\t%s", exp, act, fmt.Sprintf(msg, v...))
}

//
// Helpers
//

// p is the function that actually outputs the text
func p(tb TestingHarness, s string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	msg := fmt.Sprintf(s, v...)
	msg = "%s:%d: " + msg
	fmt.Fprintf(
		out,
		colorsStart+msg+colorsEnd,
		filepath.Base(file),
		line,
	)
	fmt.Fprintf(out, "\n\n")
	tb.FailNow()
}
