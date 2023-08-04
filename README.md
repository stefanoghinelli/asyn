# asyn

> Package to list synonyms for a word from a thesaurus using the Datamuse API 

## Installation

> Note: go >= 1.20 required

```
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
$ go install -v github.com/stefanoghinelli/asyn@latest
```

## Usage

```
$ asyn list -w bike -r 5
Synonyms for bike:
 - cycle
 - wheel
 - pedal
 - bicycle
 - motorcycle
```

## Options

```
  -h, --help          help for list
  -r, --results int   Number of results in output (default 10)
  -w, --word string   The word to find synonyms for
```

Run `asyn --help` for more information about it.
