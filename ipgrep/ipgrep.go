package ipgrep

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

// Grep finds IP addresses matching pattern
// in Reader and writes matching to Writer
func Grep(r io.Reader, w io.Writer, pattern string) (err error) {
	p := parseIPNet(pattern)
	if p == nil {
		return fmt.Errorf("bad pattern: %v", pattern)
	}

	lineScanner := bufio.NewScanner(r)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if err := lineScanner.Err(); err != nil {
			return fmt.Errorf("reading standard input: %v", err)
		}
		for _, word := range strings.Fields(line) {
			ipnet := parseIPNet(word)
			if ipnet != nil && (p.Contains(ipnet.IP) || ipnet.Contains(p.IP)) {
				fmt.Fprintln(w, line)
				break
			}
		}
	}
	return
}

func parseIPNet(word string) (ipnet *net.IPNet) {
	ipnet = &net.IPNet{}
	parts := strings.FieldsFunc(word, func(c rune) bool { return c == '/' || c == 'm' })
	if parts == nil || len(parts) == 0 || len(parts) > 2 {
		return nil
	}
	if ipnet.IP = net.ParseIP(parts[0]); ipnet.IP == nil {
		return nil
	}

	ipv4 := false
	if ipnet.IP.To4() != nil {
		ipv4 = true
		ipnet.Mask = net.CIDRMask(32, 32)
	} else if ipnet.IP.To16() != nil {
		ipnet.Mask = net.CIDRMask(128, 128)
	} else {
		return nil
	}

	if len(parts) == 2 {
		ipnet.Mask = parseMask(parts[1], ipv4)
	}
	return
}

func parseMask(mask string, ipv4 bool) net.IPMask {
	prefixLen, err := strconv.Atoi(mask)
	if err == nil {
		if ipv4 {
			return net.CIDRMask(prefixLen, 32)
		}
		return net.CIDRMask(prefixLen, 128)

	}
	return net.IPMask(net.ParseIP(mask))
}
