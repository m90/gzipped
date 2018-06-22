package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/m90/gzipped"
)

func main() {
	var (
		file        = flag.String("file", "", "location of file to be gzipped")
		showBytes   = flag.Bool("bytes", false, "display sizes in raw bytes instead of humanized formats")
		readTimeout = flag.Duration("timeout", time.Second*2, "deadline for stdin to supply data")
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
		deadline := time.NewTimer(*readTimeout).C
		cancelDeadline := &sync.Once{}
		startedReading := make(chan bool)

		go func() {
			select {
			case <-deadline:
				fmt.Printf("Received no input on stdin for %v, did you mean to pass -file instead?\n", *readTimeout)
				os.Exit(1)
			case <-startedReading:
				return
			}
		}()

		reader := bufio.NewReader(os.Stdin)
		for {
			chunk, err := reader.ReadByte()
			cancelDeadline.Do(func() { startedReading <- true })
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
