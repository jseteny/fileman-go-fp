# fileman-go-fp

You can find the following functional programming concepts in the code.

## Partial Application 

The following function implements the Partial Application functional programming
concept. In this case it makes it possible to optimize out the
unnecessary strings.ReplaceAll() calls.

Please see https://en.wikipedia.org/wiki/Partial_application
```go
func makeCalcList(path string) func(dirs []string) []string {
	exp := strings.ReplaceAll(path, "/", ".*/.*")

	return func(dirs []string) []string {
		result := make([]string, 0)
		for _, dir := range dirs {
			ok, err := regexp.MatchString(exp, dir)
			if err != nil {
				panic(err)
			}
			if ok {
				result = append(result, dir)
			}
		}
		return result
	}
}
```

You can use it this way

```go
calcList := makeCalcList(pattern)
...
list := calcList(dirs)
```

or using the following short hand form

```
got := makeCalcList(tt.pattern)(dirs)

```
