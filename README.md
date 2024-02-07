# rand

A Go package and CLI utility that generates random strings based on given paramenters.

# Package

Get the package with:

```bash
go get github.com/nu12/rand@v1.1.0
```

Use any of the available functions below:

```go
package main

import (
	rand "github.com/nu12/rand/pkg"
)

func main() {
	length := 12
	lowerOnly := false
	upperOnly := false
	rand.Alpha(length, lowerOnly, upperOnly)
	rand.Char(length, lowerOnly, upperOnly)
	rand.Num(length)
	rand.Password(length, lowerOnly, upperOnly)
	rand.Special(length)
	rand.UUID()
}
```

# CLI

## Install

### Go install

Run `go install` to download the binary to the go's binary folder:

```bash
go install github.com/nu12/rand@latest
```

Note: go's binary folder (tipically `~/go/bin`) should be added to your PATH.

### From release (x86_64 only)

Download a tagged release binary for your OS (ubuntu, macos, windows) placing it in a folder in your PATH and make it executable (may require elevated permissions):

```bash
wget -O /usr/local/bin/rand https://github.com/nu12/rand/releases/download/vX.Y.Z/rand-ubuntu
chmod +x /usr/local/bin/rand
```

Note: replace `X.Y.Z` with a valid version from the repository's releases and `ubuntu` with the appropriate OS.

### From source

Clone this repo and compile the source code:

```bash
git clone github.com/nu12/rand
cd rand
go build -o rand main.go
```

Move binary to a bin folder in your PATH (may require elevated permissions):
```bash
mv rand /usr/local/bin/
```

## Usage

General usage for all commands is `rand [command] length [flags]`. Find out all available commands with `rand`:

```
A cli utility that generates random strings based on given paramenters

Usage:
  rand [command]

Available Commands:
  alpha       Generate a random string of characters and numbers of a given length
  char        Generate a random string of a given length
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  num         Generate a random number of a given length
  password    Generate a random string of characters, numbers and special signs of a given length
  special     Generate a random string of special characters of a given length
  uuid        Generate an UUID

Flags:
  -h, --help     help for rand

Use "rand [command] --help" for more information about a command.
```

For further information about a specific command, use the `-h` flag (i.e. `rand char -h`).

#### Docker

Run the following docker command to run `rand` without a local instalation:

```
docker run --rm ghcr.io/nu12/rand [command] length [flags]
```
