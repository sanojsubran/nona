# nona

`nona` renames files to lowercase kebab-case — spaces, hyphens, and underscores are collapsed into a single hyphen, and the name is lowercased.

## Usage

```
nona <file> [file ...]
```

### Examples

```
nona "My File Name.txt"        # -> my-file-name.txt
nona "some_snake_case.go"      # -> some-snake-case.go
nona "Mixed--Separators.md"    # -> mixed-separators.md
```

Files whose names are already normalized are left untouched. Each rename is printed to stdout as `old -> new`.

## Install

```
go install nona@latest
```

Or build from source:

```
git clone https://github.com/sanojsubran/nona
cd nona
go build -o nona .
```

## License

MIT
