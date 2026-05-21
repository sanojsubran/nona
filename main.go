package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var separatorRe = regexp.MustCompile(`[ \-_]+`)

func normalize(name string) string {
	name = strings.ToLower(name)
	name = separatorRe.ReplaceAllString(name, "-")
	return name
}

func rename(path string) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newBase := normalize(base)
	if newBase == base {
		return nil
	}
	newPath := filepath.Join(dir, newBase)
	if err := os.Rename(path, newPath); err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", path, newPath)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: lpn <file> [file ...]")
		os.Exit(1)
	}
	var failed bool
	for _, arg := range os.Args[1:] {
		if err := rename(arg); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
}
