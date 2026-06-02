// SPDX-License-Identifier: MIT
// Copyright (c) 2026 sanojsubran

package renamer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		input string
		style Style
		want  string
	}{
		// Kebab (default)
		{"hello", Kebab, "hello"},
		{"Hello World", Kebab, "hello-world"},
		{"foo_bar", Kebab, "foo-bar"},
		{"foo-bar", Kebab, "foo-bar"},
		{"foo--bar", Kebab, "foo-bar"},
		{"foo__bar", Kebab, "foo-bar"},
		{"foo - bar", Kebab, "foo-bar"},
		{"FOO_BAR_BAZ", Kebab, "foo-bar-baz"},
		{"already-normalized", Kebab, "already-normalized"},
		{"Screenshot 2024-01-15 at 10.30.45 AM.png", Kebab, "screenshot-2024-01-15-at-10.30.45-am.png"},
		{"Screenshot 2026-05-13 at 9.49.37 PM.png", Kebab, "screenshot-2026-05-13-at-9.49.37-pm.png"},

		// Snake
		{"Hello World", Snake, "hello_world"},
		{"foo-bar", Snake, "foo_bar"},
		{"FOO_BAR_BAZ", Snake, "foo_bar_baz"},
		{"foo - bar", Snake, "foo_bar"},
		{"Screenshot 2024-01-15 at 10.30.45 AM.png", Snake, "screenshot_2024_01_15_at_10.30.45_am.png"},
		{"Screenshot 2026-05-13 at 9.49.37 PM.png", Snake, "screenshot_2026_05_13_at_9.49.37_pm.png"},

		// Camel
		{"Hello World", Camel, "HelloWorld"},
		{"foo-bar", Camel, "FooBar"},
		{"foo_bar_baz", Camel, "FooBarBaz"},
		{"foo - bar", Camel, "FooBar"},
		{"Screenshot 2024-01-15 at 10.30.45 AM.png", Camel, "Screenshot20240115At10.30.45Am.png"},
		{"Screenshot 2026-05-13 at 9.49.37 PM.png", Camel, "Screenshot20260513At9.49.37Pm.png"},
	}

	for _, c := range cases {
		got := Normalize(c.input, c.style)
		if got != c.want {
			t.Errorf("Normalize(%q, %v) = %q, want %q", c.input, c.style, got, c.want)
		}
	}
}

func TestRename(t *testing.T) {
	t.Run("kebab: renames file when name changes", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "Hello World.txt")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src, Kebab); err != nil {
			t.Fatalf("Rename returned error: %v", err)
		}

		dst := filepath.Join(dir, "hello-world.txt")
		if _, err := os.Stat(dst); err != nil {
			t.Errorf("expected %q to exist: %v", dst, err)
		}
		if _, err := os.Stat(src); !os.IsNotExist(err) {
			t.Errorf("expected %q to be gone", src)
		}
	})

	t.Run("snake: renames file when name changes", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "Hello World.txt")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src, Snake); err != nil {
			t.Fatalf("Rename returned error: %v", err)
		}

		dst := filepath.Join(dir, "hello_world.txt")
		if _, err := os.Stat(dst); err != nil {
			t.Errorf("expected %q to exist: %v", dst, err)
		}
		if _, err := os.Stat(src); !os.IsNotExist(err) {
			t.Errorf("expected %q to be gone", src)
		}
	})

	t.Run("camel: renames file when name changes", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "hello world.txt")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src, Camel); err != nil {
			t.Fatalf("Rename returned error: %v", err)
		}

		dst := filepath.Join(dir, "HelloWorld.txt")
		if _, err := os.Stat(dst); err != nil {
			t.Errorf("expected %q to exist: %v", dst, err)
		}
		if _, err := os.Stat(src); !os.IsNotExist(err) {
			t.Errorf("expected %q to be gone", src)
		}
	})

	t.Run("no-op when name is already normalized", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "already-normalized.txt")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src, Kebab); err != nil {
			t.Fatalf("Rename returned error: %v", err)
		}

		if _, err := os.Stat(src); err != nil {
			t.Errorf("expected %q to still exist: %v", src, err)
		}
	})

	t.Run("macOS screenshot filename", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "Screenshot 2024-01-15 at 10.30.45 AM.png")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src, Kebab); err != nil {
			t.Fatalf("Rename returned error: %v", err)
		}

		dst := filepath.Join(dir, "screenshot-2024-01-15-at-10.30.45-am.png")
		if _, err := os.Stat(dst); err != nil {
			t.Errorf("expected %q to exist: %v", dst, err)
		}
		if _, err := os.Stat(src); !os.IsNotExist(err) {
			t.Errorf("expected %q to be gone", src)
		}
	})

	t.Run("returns error for nonexistent file", func(t *testing.T) {
		dir := t.TempDir()
		if err := Rename(filepath.Join(dir, "Nonexistent File.txt"), Kebab); err == nil {
			t.Error("expected error for nonexistent file, got nil")
		}
	})
}
