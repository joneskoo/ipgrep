// Copyright (c) 2016 ipgrep contributors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/joneskoo/ipgrep/internal/ipgrep"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", name, err)
		os.Exit(1)
	}
}

func run() error {
	output := os.Stdout
	input, pattern, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	return ipgrep.Grep(input, output, pattern)
}

type inputFile struct {
	name string
	io.Reader
}

func parseArgs(args []string) (input inputFile, pattern string, err error) {
	switch {
	// ipgrep CIDR
	// ipgrep CIDR -
	case len(args) == 1+1:
		fallthrough
	case len(args) == 1+2 && args[2] == "-":
		input = inputFile{"-", os.Stdin}
	// ipgrep CIDR FILE
	case len(args) == 1+2:
		fileName := args[2]
		var f *os.File
		f, err = os.Open(fileName)
		input = inputFile{fileName, f}
	default:
		err = fmt.Errorf("%s", usage)
	}
	if err != nil {
		return inputFile{}, "", err
	}

	pattern = args[1]

	return input, pattern, nil
}

const (
	name  = "ipgrep"
	usage = `bad usage
Usage: ipgrep PATTERN [FILE...]

    E.g. ipgrep 2001:db8::/64 log.txt
	 ipgrep 127.0.0.1 log.txt
	 cat log.txt | ipgrep 127.0.0.1`
)
