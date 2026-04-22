//ff:func feature=ssac-parse type=parser control=sequence
//ff:what validates that input pairs without a colon are rejected — e.g. {query} form

package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseInputsNoColon(t *testing.T) {
	src := `package service

// @get []Gig gigs = Gig.List({query})
func ListGigs(c *gin.Context) {}
`
	dir := t.TempDir()
	path := filepath.Join(dir, "test.go")
	if err := os.WriteFile(path, []byte(src), 0644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParseFile(path)
	if len(diags) == 0 {
		t.Fatal("expected diagnostic for input without colon")
	}
	if !strings.Contains(diags[0].Message, "is not a valid input pair") {
		t.Errorf("unexpected diagnostic message: %s", diags[0].Message)
	}
}
