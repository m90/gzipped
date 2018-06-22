package gzipped

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"strings"
)

// A Comparison contains info about a buffer before
// and after its content have been gzipped
type Comparison struct {
	In       string
	InBytes  uint64
	Out      string
	OutBytes uint64
	Ratio    float32
}

const (
	sizeByte = 1.0 << (10 * iota)
	sizeKiloByte
	sizeMegaByte
	sizeGigaByte
	sizeTeraByte
)

func humanize(size uint64) string {
	unit := ""
	value := float32(size)

	switch {
	case size >= sizeTeraByte:
		unit = "T"
		value = value / sizeTeraByte
	case size >= sizeGigaByte:
		unit = "G"
		value = value / sizeGigaByte
	case size >= sizeMegaByte:
		unit = "M"
		value = value / sizeMegaByte
	case size >= sizeKiloByte:
		unit = "K"
		value = value / sizeKiloByte
	case size >= sizeByte:
		unit = "B"
	case size == 0:
		return "0"
	}

	stringified := strings.TrimSuffix(fmt.Sprintf("%.1f", value), ".0")
	return fmt.Sprintf("%s%s", stringified, unit)
}

// Compare gzips the contents of the passed buffer and returns information
// about the sizes of the original and the compressed version
func Compare(inBuf *bytes.Buffer) (*Comparison, error) {
	if inBuf == nil || inBuf.Len() == 0 {
		return nil, errors.New("unable to use empty or nil buffer")
	}
	var outBuf bytes.Buffer
	gzipped := gzip.NewWriter(&outBuf)

	if _, err := gzipped.Write(inBuf.Bytes()); err != nil {
		return nil, err
	}
	if err := gzipped.Flush(); err != nil {
		return nil, err
	}
	if err := gzipped.Close(); err != nil {
		return nil, err
	}

	inBytes, outBytes := uint64(len(inBuf.Bytes())), uint64(len(outBuf.Bytes()))
	return &Comparison{
		In:       humanize(inBytes),
		InBytes:  inBytes,
		Out:      humanize(outBytes),
		OutBytes: outBytes,
		Ratio:    float32(outBytes) / float32(inBytes) * 100,
	}, nil
}
