package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/m90/gzipped"
)

func main() {
	var (
		file      = flag.String("file", "", "file to be gzipped")
		showBytes = flag.Bool("bytes", false, "display sizes in bytes")
	)
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

	result, err := gzipped.Compare(bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("Error compressing data: %v\n", err)
		os.Exit(1)
	}

	if *showBytes {
		fmt.Printf("Original is %d Bytes, compressed is %d Bytes, ratio is %.1f%%\n", result.InBytes, result.OutBytes, result.Ratio)
	} else {
		fmt.Printf("Original is %s, compressed is %s, ratio is %.1f%%\n", result.In, result.Out, result.Ratio)
	}
}
