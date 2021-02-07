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

package ipgrep

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"
	"strconv"
	"strings"

	"inet.af/netaddr"
)

// Grep finds IP addresses matching pattern
// in Reader and writes matching to Writer
func Grep(r io.Reader, w io.Writer, search string) error {
	searchPrefix, err := parseIPPrefix(search)
	if err != nil {
		return fmt.Errorf("bad pattern: %w", err)
	}
	lineScanner := bufio.NewScanner(r)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if err := lineScanner.Err(); err != nil {
			return fmt.Errorf("reading standard input: %v", err)
		}
		for _, word := range strings.Fields(line) {
			ipp, err := parseIPPrefix(word)
			if err != nil {
				continue
			}
			if searchPrefix.Contains(ipp.IP) || ipp.Contains(searchPrefix.IP) {
				fmt.Fprintln(w, line)
				break // next line
			}
		}
	}
	return nil
}

const (
	acceptMForMask      = true
	acceptLegacyNetmask = true
)

func parseIPPrefix(s string) (prefix netaddr.IPPrefix, err error) {
	is := strings.IndexByte(s, '/')
	im := -1
	if acceptMForMask {
		im = strings.IndexByte(s, 'm')
	}
	// i is position of the separator or end of string if no separator
	var i int
	switch {
	case is >= 0:
		i = is
	case im >= 0:
		i = im
	default:
		i = len(s)
	}

	// Parse IP part of the prefix
	ip, err := netaddr.ParseIP(s[:i])
	if err != nil {
		return netaddr.IPPrefix{}, err
	}

	// s is a single IP (convert to single IP CIDR)
	if i == len(s) {
		return netaddr.IPPrefix{IP: ip, Bits: ip.BitLen()}, nil
	}

	s = s[i+1:]
	prefixLen, err := strconv.Atoi(s)
	if prefixLen < 0 || prefixLen > 128 {
		return netaddr.IPPrefix{}, fmt.Errorf("bad prefix length %q: %v", s, err)
	}
	if err != nil {
		if !acceptLegacyNetmask {
			return netaddr.IPPrefix{}, fmt.Errorf("bad prefix %q: %v", s, err)
		}
		mask, err := netaddr.ParseIP(s)
		if err != nil {
			return netaddr.IPPrefix{}, err
		}
		maskBytes := mask.As4()
		prefixLen = bits.OnesCount32(binary.BigEndian.Uint32(maskBytes[:]))

	}
	return ip.Prefix(uint8(prefixLen))
}
