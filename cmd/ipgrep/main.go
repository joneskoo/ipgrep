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

	"github.com/joneskoo/ipgrep"
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

	args := os.Args[1:]
	if len(args) < 1 {
		return fmt.Errorf("%s", usage)
	}
	pattern := args[0]
	if len(args) == 1 {
		args = append(args, "-")
	}

	for _, fileName := range args[1:] {
		var f *os.File
		if fileName == "-" {
			f = os.Stdin
		} else {
			var err error
			f, err = os.Open(fileName)
			if err != nil {
				return err
			}
		}
		if err := ipgrep.Grep(f, output, pattern); err != nil {
			return err
		}
	}
	return nil
}

type inputFile struct {
	name string
	io.Reader
}

const (
	name  = "ipgrep"
	usage = `not enough arguments

Usage: ipgrep PATTERN [FILE...]

    E.g. ipgrep 2001:db8::/64 log.txt
	 ipgrep 127.0.0.1 log.txt
	 cat log.txt | ipgrep 127.0.0.1`
)
