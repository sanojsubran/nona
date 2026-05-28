package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var separatorRe = regexp.MustCompile(`[ \-_]+`)

func Normalize(name string) string {
	name = strings.ToLower(name)
	name = separatorRe.ReplaceAllString(name, "-")
	return name
}

func Rename(path string) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newBase := Normalize(base)
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
