//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT03_EmptySpecsDir_NoFire

package features

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT03_EmptySpecsDir_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: ""}
	diags := ft03HashMismatch(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
