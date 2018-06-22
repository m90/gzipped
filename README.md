# gzipped
[![Build Status](https://travis-ci.org/m90/gzipped.svg?branch=master)](https://travis-ci.org/m90/gzipped)
[![godoc](https://godoc.org/github.com/m90/gzipped?status.svg)](http://godoc.org/github.com/m90/gzipped)

> Print the gzipped size of any file

A simple command to answer the age old question "But how big will it be when gzipped?".

## Installation:

Install the command:
```sh
go get github.com/m90/gzipped/cmd/gzipped
```

Install the library:
```sh
go get github.com/m90/gzipped
```

## Command usage

Use `-file` to pass the location of a file:
```sh
gzipped -file ./styles.css
```

or use pipes:
```sh
cat ./bundle.js | uglifyjs -mangle | gzipped
```

The following options are available:

```
Usage of gzipped:
  -bytes
    	display sizes in bytes
  -file string
    	file to be gzipped
```

Refer to the [godoc](http://godoc.org/github.com/m90/gzipped) for information about how to use it as a library.

### License
MIT © [Frederik Ring](http://www.frederikring.com)
