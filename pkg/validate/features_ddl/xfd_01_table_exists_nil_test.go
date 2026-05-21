//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what XFD-01 — FeatureTables nil 시 단락 테스트
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXFD01_TableExists_NilFeatureTables(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := xfd01TableExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil FeatureTables, got %d", len(diags))
	}
}
