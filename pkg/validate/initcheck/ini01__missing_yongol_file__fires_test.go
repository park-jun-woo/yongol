//ff:func feature=validate type=test control=sequence topic=init-check
//ff:what TestINI01_MissingYongolFile_Fires

package initcheck

import (
	"strings"
	"testing"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestINI01_MissingYongolFile_Fires(t *testing.T) {
	tmp := t.TempDir()

	// No .yongol file in tmp
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := Run(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[INI-01]") {
		t.Errorf("want [INI-01] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "yongol init") {
		t.Errorf("want 'yongol init' in message, got %s", diags[0].Message)
	}
}
