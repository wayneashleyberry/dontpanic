[_Don't panic™_](https://github.com/golang/go/wiki/CodeReviewComments#dont-panic)

---

[![Build Status](https://travis-ci.org/wayneashleyberry/dontpanic.svg?branch=master)](https://travis-ci.org/wayneashleyberry/dontpanic)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/dontpanic)](https://goreportcard.com/report/github.com/wayneashleyberry/dontpanic)
[![GolangCI](https://golangci.com/badges/github.com/wayneashleyberry/dontpanic.svg)](https://golangci.com/r/github.com/wayneashleyberry/dontpanic)

```
go get github.com/wayneashleyberry/dontpanic
```

```
dontpanic: don't allow panic() in your go code

Usage: dontpanic [-flag] [package]


Flags:  -V	print version and exit
  -all
    	no effect (deprecated)
  -c int
    	display offending line with this many lines of context (default -1)
  -cpuprofile string
    	write CPU profile to this file
  -debug string
    	debug flags, any subset of "fpstv"
  -flags
    	print analyzer flags in JSON
  -json
    	emit JSON output
  -memprofile string
    	write memory profile to this file
  -source
    	no effect (deprecated)
  -tags string
    	no effect (deprecated)
  -trace string
    	write trace log to this file
  -v	no effect (deprecated)
```
