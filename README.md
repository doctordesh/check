# Check

Super small and simple check functions for Go testing. Stolen from [Ben Johnson](https://github.com/benbjohnson/testing)

```golang
func TestThing(t *testing.T) {
	value, err := DoSomething()
	check.OK(t, err)
	check.Equals(t, value, 100)
	check.Assert(t, value == 100)
}
```