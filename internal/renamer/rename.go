package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var separatorRe = regexp.MustCompile(`[ \-_]+`)

func Normalize(name string, style Style) string {
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	words := separatorRe.Split(strings.ToLower(base), -1)

	var stem string
	switch style {
	case Snake:
		stem = strings.Join(words, "_")
	case Camel:
		for i, w := range words {
			if len(w) > 0 {
				words[i] = strings.ToUpper(w[:1]) + w[1:]
			}
		}
		stem = strings.Join(words, "")
	default:
		stem = strings.Join(words, "-")
	}

	return stem + strings.ToLower(ext)
}

func Rename(path string, style Style) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newBase := Normalize(base, style)
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
