// SPDX-License-Identifier: MIT
// Copyright (c) 2026 sanojsubran

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"nona/internal/renamer"
)

var version = "dev"

func main() {
	styleFlag := flag.String("style", "kebab", "naming style: kebab, snake, or camel")
	replaceFlag := flag.String("replace", "", "replace a substring in the filename: `old=new`")
	versionFlag := flag.Bool("version", false, "print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  nona [--style kebab|snake|camel] <file> [file ...]\n")
		fmt.Fprintf(os.Stderr, "  nona --replace old=new <file> [file ...]\n\n")
		fmt.Fprintf(os.Stderr, "Normalize filenames to a consistent style, or replace a substring across files.\n\n")
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

	if *replaceFlag != "" {
		var styleExplicit bool
		flag.Visit(func(f *flag.Flag) {
			if f.Name == "style" {
				styleExplicit = true
			}
		})
		if styleExplicit {
			fmt.Fprintln(os.Stderr, "error: --replace and --style cannot be used together")
			os.Exit(1)
		}

		parts := strings.SplitN(*replaceFlag, "=", 2)
		if len(parts) != 2 || parts[0] == "" {
			fmt.Fprintln(os.Stderr, "error: --replace must be in old=new format")
			os.Exit(1)
		}
		var failed bool
		for _, arg := range flag.Args() {
			if err := renamer.ReplaceInName(arg, parts[0], parts[1]); err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				failed = true
			}
		}
		if failed {
			os.Exit(1)
		}
		return
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
