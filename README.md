# doktor

take two and call me in the morning (⌐■_■)

## Installation 
```
> brew tap adamenger/doktor
> brew install doktor
```

## Usage
```
NAME:
   doktor - take two and call me in the morning (⌐■_■)

USAGE:
   doktor [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   malware scanning and removal tool

COMMANDS:
     clean, c  clean all infections doktor finds
     list, l   list all the infections doktor about
     scan, s   doktor will scan the host for known malware signatures
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Development

You will need a golang dev environment to contribute to doktor.

You'll need to run the build script and redirect the output to a `*.go` file. Once you do this, you should be ready to compile doktor.

```
$> ./build.sh > signatureData.go
$> go build
```


## Why?

doktor is meant to be a very simple malware scanner for OSX written in Go. Currently the scanner is very simple, it will read the paths from the signature file and check if they exist.

If the path exists, this is labeled as a malware hit. This is not meant to be sophisticated tool, I just wanted something I could run from the command line and install using brew.
