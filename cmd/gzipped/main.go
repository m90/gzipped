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

func main() {
	file := flag.String("file", "", "file to read")
	flag.Parse()

	var reader *bufio.Reader
	var bytesRead int
	var buf bytes.Buffer

	gzipped := gzip.NewWriter(&buf)

	if location := *file; location != "" {
		b, err := ioutil.ReadFile(*file)
		if err != nil {
			fmt.Printf("Error reading file %v\n", err)
			os.Exit(1)
		}
		bytesRead = len(b)
		gzipped.Write(b)
	} else {
		reader = bufio.NewReader(os.Stdin)
		for {
			chunk, err := reader.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
			gzipped.Write([]byte{chunk})
			bytesRead++
		}
	}

	if err := gzipped.Flush(); err != nil {
		fmt.Printf("Error flushing reader %v\n", err)
		os.Exit(1)
	}
	if err := gzipped.Close(); err != nil {
		fmt.Printf("Error closing reader %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got %d bytes, compressed to %d bytes \n", bytesRead, len(buf.Bytes()))
}
