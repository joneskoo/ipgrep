package main

import (
	"fmt"
	"os"

	"github.com/joneskoo/ipgrep/ipgrep"
)

const (
	name  = "ipgrep"
	usage = `bad usage
Usage: ipgrep IP[/mask] file.txt

    E.g. ipgrep 2001:db8::/64 log.txt
         ipgrep 127.0.0.1 log.txt`
)

func main() {
	if len(os.Args) != 2+1 {
		fatal(usage)
	}
	pattern := os.Args[1]
	fileName := os.Args[2]
	f, err := os.Open(fileName)
	if err != nil {
		fatal(err)
	}
	err = ipgrep.Grep(f, os.Stdout, pattern)
	if err != nil {
		fatal(err)
	}
}

func fatal(a ...interface{}) {
	fmt.Fprintf(os.Stderr, "%v: ", name)
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
