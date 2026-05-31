//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestFastapiSsacHelpers_ZeroCov — addExtPkgRef/appendSnakeRune 커버
package ssac

import (
	"testing"
)

func TestAppendSnakeRune_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"OrgID":         "org_id",
		"ResolveRootID": "resolve_root_id",
		"user":          "user",
		"ID":            "id",
	}
	for in, want := range cases {
		if got := snakeViaRune(in); got != want {
			t.Errorf("snake(%q) = %q, want %q", in, got, want)
		}
	}
}
