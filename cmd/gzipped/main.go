package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Result struct {
	InBytes  uint64
	OutBytes uint64
}

func compare(inBuf *bytes.Buffer) (*Result, error) {
	var outBuf bytes.Buffer
	gzipped := gzip.NewWriter(&outBuf)
	gzipped.Write(inBuf.Bytes())

	if err := gzipped.Flush(); err != nil {
		return nil, err
	}
	if err := gzipped.Close(); err != nil {
		return nil, err
	}

	return &Result{
		InBytes:  uint64(len(inBuf.Bytes())),
		OutBytes: uint64(len(outBuf.Bytes())),
	}, nil
}

func main() {
	file := flag.String("file", "", "file to read")
	flag.Parse()

	var b []byte

	if location := *file; location != "" {
		var readErr error
		b, readErr = ioutil.ReadFile(*file)
		if readErr != nil {
			fmt.Printf("Error reading file: %v\n", readErr)
			os.Exit(1)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			chunk, err := reader.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
			b = append(b, chunk)
		}
	}

	result, err := compare(bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("Error compressing data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got %d bytes, compressed to %d bytes \n", result.InBytes, result.OutBytes)
}
