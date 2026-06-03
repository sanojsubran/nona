# nona

![nona logo](docs/nona-logo.svg)

`nona` renames files by normalizing their names into a consistent style. Spaces, hyphens, and underscores are treated as word separators and collapsed. The default output style is kebab-case.

## Usage

```
nona [--style kebab|snake|camel] <file> [file ...]
```

The `--style` flag is optional and defaults to `kebab`.

## Styles

| Flag value | Example output |
|------------|---------------|
| `kebab` (default) | `hello-world.txt` |
| `snake` | `hello_world.txt` |
| `camel` | `HelloWorld.txt` |

## Examples

```
nona "My File Name.txt"                    # -> my-file-name.txt
nona --style snake "My File Name.txt"      # -> my_file_name.txt
nona --style camel "My File Name.txt"      # -> MyFileName.txt

nona "Screenshot 2024-01-15 at 10.30.45 AM.png"
# -> screenshot-2024-01-15-at-10.30.45-am.png
```

Files whose names are already normalized are left untouched. Each rename is printed to stdout as `old -> new`.

## Install

```
brew tap sanojsubran/apps
brew install nona
```

Or with `go install`:

```
go install github.com/sanojsubran/nona/cmd/nona@latest
```

Or build from source:

```
git clone https://github.com/sanojsubran/nona
cd nona
make build
```

## License

MIT
