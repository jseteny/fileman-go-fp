package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dirs = []string{".", ".idea", ".idea/.gitignore", ".idea/fileman.iml", ".idea/modules.xml", ".idea/workspace.xml", "cd.go", "cd_test.go", "go.mod"}

func Test_FindPath(t *testing.T) {

	tests := []struct {
		pattern string
		want    []string
	}{
		{pattern: "a", want: []string{".idea", ".idea/.gitignore", ".idea/fileman.iml", ".idea/modules.xml", ".idea/workspace.xml"}},
		{pattern: "o", want: []string{".idea/.gitignore", ".idea/modules.xml", ".idea/workspace.xml", "cd.go", "cd_test.go", "go.mod"}},
		{pattern: "a/o", want: []string{".idea/.gitignore", ".idea/modules.xml", ".idea/workspace.xml"}},
		{pattern: "go", want: []string{"cd.go", "cd_test.go", "go.mod"}},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			got := makeCalcList(tt.pattern)(dirs)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func BenchmarkFindPath(t *testing.B) {
	for n := 0; n < t.N; n++ {
		makeCalcList("a/o")(dirs)
	}
}

func BenchmarkFindPathOptimized(t *testing.B) {
	calcList := makeCalcList("a/o")
	for n := 0; n < t.N; n++ {
		calcList(dirs)
	}
}
