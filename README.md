# Check

Super small and simple check functions for Go testing. Stolen from [Ben Johnson](https://github.com/benbjohnson/testing)

## Install

```
go get github.com/doctordesh/check
```

## Usage

```golang
package foo

import (
	"testing"

	"github.com/doctordesh/check"
)

func TestBar(t *testing.T) {
	expectedValue := 100
	actualValue, err := DoSomething()
	check.OK(t, err)
	check.Equals(t, expectedValue, actualValue)
	check.Assert(t, expectedValue == actualValue, "expected equal")
}

```

## Godoc

```golang
func Assert(tb TestingHarness, condition bool)
func AssertWithMessage(tb TestingHarness, condition bool, msg string, v ...interface{})
func Equals(tb TestingHarness, exp, act interface{})
func EqualsWithMessage(tb TestingHarness, exp, act interface{}, msg string, v ...interface{})
func NotOK(tb TestingHarness, err error)
func NotOKWithMessage(tb TestingHarness, err error, msg string, v ...interface{})
func OK(tb TestingHarness, err error)
func OKWithMessage(tb TestingHarness, err error, msg string, v ...interface{})
type TestingHarness interface{ ... }
```
