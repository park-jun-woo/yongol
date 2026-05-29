//ff:func feature=validate type=test control=sequence topic=init-check
//ff:what TestINI01_EmptySpecsDir_NoFire

package initcheck

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestINI01_EmptySpecsDir_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: ""}
	diags := Run(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
