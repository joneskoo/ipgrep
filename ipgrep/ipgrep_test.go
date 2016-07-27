package ipgrep_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/joneskoo/ipgrep/ipgrep"
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
