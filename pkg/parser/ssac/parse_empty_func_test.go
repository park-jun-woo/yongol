//ff:func feature=ssac-parse type=parser control=sequence
//ff:what S-74 test — verifies that an SSaC function with no annotations produces an error diagnostic

package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseEmptyFunc(t *testing.T) {
	dir := t.TempDir()
	subdir := filepath.Join(dir, "workflow")
	os.MkdirAll(subdir, 0755)

	src := `package service

func ActivateWorkflow() {}
`
	path := filepath.Join(subdir, "activate.ssac")
	os.WriteFile(path, []byte(src), 0644)

	_, diags := ParseFile(path)
	if len(diags) == 0 {
		t.Fatal("expected S-74 diagnostic for empty function, got none")
	}
	if !strings.Contains(diags[0].Message, "[S-74]") {
		t.Errorf("expected [S-74] in message, got: %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "ActivateWorkflow") {
		t.Errorf("expected function name in message, got: %s", diags[0].Message)
	}
}
