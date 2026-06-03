// SPDX-License-Identifier: MIT
// Copyright (c) 2026 sanojsubran

package main

import (
	"flag"
	"fmt"
	"os"

	"nona/internal/renamer"
)

var version = "dev"

func main() {
	styleFlag := flag.String("style", "kebab", "naming style: kebab, snake, or camel")
	versionFlag := flag.Bool("version", false, "print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: nona [options] <file> [file ...]\n\n")
		fmt.Fprintf(os.Stderr, "Rename files to a consistent naming style.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *versionFlag {
		fmt.Println("nona", version)
		return
	}

	if flag.NArg() == 0 {
		flag.Usage()
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
