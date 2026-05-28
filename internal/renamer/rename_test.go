package renamer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"hello", "hello"},
		{"Hello World", "hello-world"},
		{"foo_bar", "foo-bar"},
		{"foo-bar", "foo-bar"},
		{"foo--bar", "foo-bar"},
		{"foo__bar", "foo-bar"},
		{"foo - bar", "foo-bar"},
		{"FOO_BAR_BAZ", "foo-bar-baz"},
		{"already-normalized", "already-normalized"},
		{"Screenshot 2024-01-15 at 10.30.45 AM.png", "screenshot-2024-01-15-at-10.30.45-am.png"},
	}

	for _, c := range cases {
		got := Normalize(c.input)
		if got != c.want {
			t.Errorf("Normalize(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}

func TestRename(t *testing.T) {
	t.Run("renames file when name changes", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "Hello World.txt")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src); err != nil {
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

	t.Run("no-op when name is already normalized", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "already-normalized.txt")
		if err := os.WriteFile(src, nil, 0644); err != nil {
			t.Fatal(err)
		}

		if err := Rename(src); err != nil {
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

		if err := Rename(src); err != nil {
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
		if err := Rename(filepath.Join(dir, "Nonexistent File.txt")); err == nil {
			t.Error("expected error for nonexistent file, got nil")
		}
	})
}
