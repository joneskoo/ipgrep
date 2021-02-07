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
	if acceptMForMask {
		s = strings.Replace(s, "m", "/", 1)
	}
	if acceptLegacyNetmask {
		pos := strings.Index(s, "/")
		if pos > 0 {
			mask, err := netaddr.ParseIP(s[pos+1:])
			if err == nil {
				maskBytes := mask.As4()
				bitLen := bits.OnesCount32(binary.BigEndian.Uint32(maskBytes[:]))
				s = fmt.Sprintf("%s%d", s[:pos+1], bitLen)
			}
		}
	}

	prefix, err = netaddr.ParseIPPrefix(s)
	if err != nil {
		// Search pattern is a single IP (convert to single IP CIDR)
		var ip netaddr.IP
		ip, err = netaddr.ParseIP(s)
		if err != nil {
			return prefix, err
		}
		// Ignore error: Prefix of BitLen cannot be too large.
		prefix, _ = ip.Prefix(ip.BitLen())
	}
	return prefix, nil
}
