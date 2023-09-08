# rand

A cli utility that generates random strings based on given paramenters.

## Install

### Go install

Run `go install` to download the binary to the go's binary folder:

```
go install github.com/nu12/rand
```

Note: go's binary folder (tipically `~/go/bin`) should be added to your PATH.

### From release (x86_64 only)

Download a tagged release binary for your OS (ubuntu, macos, windows) placing it in a folder in your PATH and make it executable (may require elevated permissions):

```
wget -O /usr/local/bin/rand https://github.com/nu12/rand/releases/download/vX.Y/rand-ubuntu
chmod +x /usr/local/bin/rand
```

Note: replace `X.Y` with a valid version from the repository's releases and `ubuntu` with the appropriate OS.

### From source

Clone this repo and compile the source code:

```
git clone github.com/nu12/rand
cd rand
go build -o rand main.go
```

Move binary to a bin folder in your PATH (may require elevated permissions):
```
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
  -t, --toggle   Help message for toggle

Use "rand [command] --help" for more information about a command.
```

For further information about a specific command, use the `-h` flag (i.e. `rand char -h`).
