// SPDX-License-Identifier: MIT
// Copyright (c) 2026 sanojsubran

package main

import (
	"flag"
	"fmt"
	"os"

	"nona/internal/renamer"
)

func main() {
	styleFlag := flag.String("style", "kebab", "naming style: kebab, snake, camel")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: nona [--style kebab|snake|camel] <file> [file ...]")
		os.Exit(1)
	}

	var style renamer.Style
	switch *styleFlag {
	case "snake":
		style = renamer.Snake
	case "camel":
		style = renamer.Camel
	case "kebab":
		style = renamer.Kebab
	default:
		fmt.Fprintf(os.Stderr, "unknown style %q: must be kebab, snake, or camel\n", *styleFlag)
		os.Exit(1)
	}

	var failed bool
	for _, arg := range flag.Args() {
		if err := renamer.Rename(arg, style); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
}
