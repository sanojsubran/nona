// SPDX-License-Identifier: MIT
// Copyright (c) 2026 sanojsubran

package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var separatorRe = regexp.MustCompile(`[ \-_\x{202F}]+`)

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

func ReplaceInName(path, from, to string) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	stem := strings.TrimSuffix(base, ext)
	newBase := strings.ReplaceAll(stem, from, to) + ext
	if newBase == base {
		_, err := os.Stat(path)
		return err
	}
	newPath := filepath.Join(dir, newBase)
	if err := os.Rename(path, newPath); err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", path, newPath)
	return nil
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
