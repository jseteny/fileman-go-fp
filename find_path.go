package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	FindPath("go")
}

// FindPath find files and directories matching the given pattern
func FindPath(pattern string) {
	dirs := make([]string, 0)
	for dir := range collectDirs(".") {
		dirs = append(dirs, dir)
		list := calcList(pattern, dirs)
		fmt.Println()
		fmt.Printf("%q\n", list)
	}
}

func collectDirs(root string) <-chan string {
	dirs := make(chan string)
	go func() {
		defer close(dirs)
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			dirs <- path
			return err
		})
		if err != nil {
			log.Fatalln(err)
		}
	}()
	return dirs
}

func calcList(path string, dirs []string) []string {
	exp := strings.ReplaceAll(path, "/", ".*/.*")
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
