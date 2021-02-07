// Copyright (c) 2016 - 2021 ipgrep contributors
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

package ipgrep_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/joneskoo/ipgrep"
)

func TestIpgrep(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		pattern string
		want    string
		err     bool
	}{
		{
			name:    "match specific IPv4 with CIDR",
			pattern: "127.0.0.2/24",
			input:   "127.0.0.1\n",
			want:    "127.0.0.1\n",
		},
		{
			name:    "match specific IPv4 with all IPv4 CIDR",
			pattern: "0.0.0.0/0",
			input:   "127.0.0.1\n",
			want:    "127.0.0.1\n",
		},
		{
			name:    "no match specific IPv4 with specific",
			pattern: "127.0.0.2",
			input:   "127.0.0.1\n",
			want:    "",
		},

		{
			name:    "no match IPv4 with IPv6 CIDR",
			pattern: "::/0",
			input:   "127.0.0.1\n",
			want:    "",
		},

		{
			name:    "match all IPv4 CIDR with specific IP",
			pattern: "127.0.0.1",
			input:   "0.0.0.0/0\n",
			want:    "0.0.0.0/0\n",
		},
		{
			name:    "match IPv4 CIDR with specific IP",
			pattern: "127.0.0.1",
			input:   "127.0.0.2/24\n",
			want:    "127.0.0.2/24\n",
		},

		{
			name:    "no match IPv4 with specific IPv4",
			pattern: "1.2.3.5",
			input:   "1.2.3.4\n",
			want:    "",
		},
		{
			name:    "no match IPv4 CIDR with specific IPv4",
			pattern: "1.2.2.4",
			input:   "1.2.3.4/24\n",
			want:    "",
		},
		{
			name:    "no match IPv4 m CIDR with specific IPv4",
			pattern: "1.2.2.4",
			input:   "1.2.3.4m24\n",
			want:    "",
		},
		{
			name:    "match IPv6 with IPv6",
			pattern: "2001:db8::1",
			input:   "2001:db8::1\n",
			want:    "2001:db8::1\n",
		},
		{
			name:    "match IPv6 CIDR with IPv6 CIDR",
			pattern: "2001:db8::2",
			input:   "2001:db8::1/64\n",
			want:    "2001:db8::1/64\n",
		},

		{
			name:    "no match IPv4 with specific IPv4",
			pattern: "1.2.3.4",
			input:   "1.2.3.5\n",
			want:    "",
		},
		{
			name:    "no match IPv4 CIDR with specific IPv4",
			pattern: "1.2.3.4",
			input:   "1.2.2.4/24\n",
			want:    "",
		},
		{
			name:    "no match IPv6 with IPv6",
			pattern: "2001:db8::2",
			input:   "2001:db8::1\n",
			want:    "",
		},
		{
			name:    "no match IPv6 CIDR with IPv6",
			pattern: "2001:db9::1",
			input:   "2001:db8::1/64\n",
			want:    "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := strings.NewReader(c.input)
			buf := &bytes.Buffer{}
			err := ipgrep.Grep(r, buf, c.pattern)
			if (err != nil) != c.err {
				t.Errorf("wanted error=%v, got %v", c.err, err)
			}
			got := buf.String()
			if got != c.want {
				t.Errorf("wanted Grep(%q) to write %q, got %q", c.input, c.want, got)
			}
		})
	}
}

func BenchmarkGrepCIDR(b *testing.B) {
	b.ReportAllocs()
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		b.Errorf("Failed to read README.md: %v", err)
	}
	for i := 0; i < b.N; i++ {
		br := bytes.NewReader(input)
		err := ipgrep.Grep(br, ioutil.Discard, "1.2.3.4/24")
		if err != nil {
			b.Errorf("Grep failed unexpectedly: %v", err)
		}
	}
}

func BenchmarkGrepIP(b *testing.B) {
	b.ReportAllocs()
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		b.Errorf("Failed to read README.md: %v", err)
	}
	for i := 0; i < b.N; i++ {
		br := bytes.NewReader(input)
		err := ipgrep.Grep(br, ioutil.Discard, "1.2.3.4")
		if err != nil {
			b.Errorf("Grep failed unexpectedly: %v", err)
		}
	}
}
