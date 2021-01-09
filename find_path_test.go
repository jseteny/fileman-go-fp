package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cd(t *testing.T) {
	dirs := []string{".", ".idea", ".idea/.gitignore", ".idea/fileman.iml", ".idea/modules.xml", ".idea/workspace.xml", "cd.go", "cd_test.go", "go.mod"}

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
