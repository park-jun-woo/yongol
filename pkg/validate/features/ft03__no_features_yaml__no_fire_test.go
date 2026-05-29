//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT03_NoFeaturesYAML_NoFire

package features

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT03_NoFeaturesYAML_NoFire(t *testing.T) {
	tmp := t.TempDir()

	// No features.yaml at all
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := ft03HashMismatch(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
