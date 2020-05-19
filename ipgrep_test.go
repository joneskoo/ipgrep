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

package ipgrep_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/joneskoo/ipgrep/internal/ipgrep"
)

func TestIpgrep(t *testing.T) {
	cases := []struct {
		input   string
		pattern string
		want    string
		err     bool
	}{
		{pattern: "127.0.0.1", input: "127.0.0.1\n", want: "127.0.0.1\n"},
		{pattern: "127.0.0.2", input: "127.0.0.1\n", want: ""},
		{pattern: "127.0.0.2/24", input: "127.0.0.1\n", want: "127.0.0.1\n"},
		{pattern: "0.0.0.0/0", input: "127.0.0.1\n", want: "127.0.0.1\n"},
		{input: "127.0.0.1\n", pattern: "127.0.0.1", want: "127.0.0.1\n"},
		{input: "127.0.0.2\n", pattern: "127.0.0.1", want: ""},
		{input: "127.0.0.2/24\n", pattern: "127.0.0.1", want: "127.0.0.2/24\n"},
		{input: "0.0.0.0/0\n", pattern: "127.0.0.1", want: "0.0.0.0/0\n"},
		{pattern: "::/0", input: "127.0.0.1\n", want: ""},
	}
	for _, c := range cases {
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
	}
}
