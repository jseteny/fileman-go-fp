# fileman-go-fp

You can find the following functional programming concepts in the code.

## Partial Application 

The following function implements the Partial Application functional programming
concept. In this case it makes it possible to optimize out the
unnecessary strings.ReplaceAll() calls.

See the benchmark results here: ![alt text](https://github.com/jseteny/fileman-go-fp/raw/master/benchmark_results/Benchmark.png?raw=true)(https://github.com/jseteny/fileman-fp-go/blob/master/image.jpg?raw=true)

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

You can use it in an optimised way doing 
the Partial Application outside the loop 
and calling the Partially Applied function
inside the loop.

```go
func BenchmarkFindPathOptimized(t *testing.B) {
	calcList := makeCalcList("a/o")
	for n := 0; n < t.N; n++ {
		calcList(dirs)
	}
}
```

or using the following short hand form

```
func BenchmarkFindPath(t *testing.B) {
	for n := 0; n < t.N; n++ {
		makeCalcList("a/o")(dirs)
	}
}
```
