//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiSsacHelpers_ZeroCov — addExtPkgRef/appendSnakeRune 커버

package ssac

import (
	"strings"
	"testing"
	"unicode"
)

func TestAddExtPkgRef_ZeroCov(t *testing.T) {
	d := &importData{ExtPkgs: map[string]map[string]bool{}}
	addExtPkgRef(d, "mail", "Send")
	addExtPkgRef(d, "mail", "Queue") // existing pkg map reused
	if !d.ExtPkgs["mail"]["Send"] || !d.ExtPkgs["mail"]["Queue"] {
		t.Fatalf("addExtPkgRef = %v", d.ExtPkgs)
	}
}

// snakeViaRune mirrors how appendSnakeRune is driven, exercising all branches.
func snakeViaRune(s string) string {
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		prevUpper := i > 0 && unicode.IsUpper(runes[i-1])
		nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
		appendSnakeRune(&b, i, r, prevUpper, nextLower)
	}
	return b.String()
}

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
