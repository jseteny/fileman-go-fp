package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cd(t *testing.T) {
	dirs := []string{".", ".idea", ".idea/.gitignore", ".idea/fileman.iml", ".idea/modules.xml", ".idea/workspace.xml", "cd.go", "cd_test.go", "go.mod"}

	tests := []struct {
		path string
		want []string
	}{
		{path: "a", want: []string{".idea", ".idea/.gitignore", ".idea/fileman.iml", ".idea/modules.xml", ".idea/workspace.xml"}},
		{path: "o", want: []string{".idea/.gitignore", ".idea/modules.xml", ".idea/workspace.xml", "cd.go", "cd_test.go", "go.mod"}},
		{path: "a/o", want: []string{".idea/.gitignore", ".idea/modules.xml", ".idea/workspace.xml"}},
		{path: "go", want: []string{"cd.go", "cd_test.go", "go.mod"}},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := calcList(tt.path, dirs)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
