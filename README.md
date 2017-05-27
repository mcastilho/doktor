# doktor

take two and call me in the morning (⌐■_■).

doktor is an extremely efficient and simple CLI malware scanner written in Golang. I have collected malware/adware/badware/shitware/crapware/tearwear/bearwear signatures from the internet for your convenience in `signatures.yaml`. 

If you have a correction or want to add a signature that I don't have, I would be very appreciative if you could do so through a PR.

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

## Credit
As with all things in technology, I owe credit to the efforts of the people who have come before me. I'd like to give credit to these fine folks for inspiring me to create this. I've also incorporated the pubic signatures you've collected into doktor, so thank you for taking the time curate and open source them.

* Whoever made https://etrecheck.com/ and
* Thomas Reed formerly of Adware Medic now MBAM : http://www.adwaremedic.com/signatures.xml
