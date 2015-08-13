package pathutil 

import (
	pathlib "path"
	"fmt"
	"os/user"
	"testing"
)

func TestExpand(t *testing.T) {
	usr, _ := user.Current()
	home := usr.HomeDir
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"~", "~"},
		{"~/", home},
		{"a/~/b", "a/~/b"},
		{"~/b", pathlib.Join(home, "b")},
	}
	for _, c := range cases {
		fmt.Printf("Testing Expand(%s)\n", c.in)
		got := Expand(c.in)
		if got != c.want {
			t.Errorf("Expand(%q) wanted %q, got %q", c.in, c.want, got)
		}
	}
}
