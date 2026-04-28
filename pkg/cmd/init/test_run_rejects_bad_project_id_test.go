//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunRejectsBadProjectID — Run rejects ProjectID containing hyphens

package cliinit

import (
	"bytes"
	"testing"
)

func TestRunRejectsBadProjectID(t *testing.T) {
	tmp := t.TempDir()
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:   "my-app", // hyphen disallowed
		Description: "Test",
		Dir:         tmp,
		Module:      "github.com/test/myapp",
	})
	if err == nil {
		t.Fatalf("Run should reject ProjectID with hyphen")
	}
}
