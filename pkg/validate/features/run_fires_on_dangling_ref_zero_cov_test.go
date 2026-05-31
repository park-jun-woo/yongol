//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what zz_zerocov_test — validate/features.Run 0% 커버리지 단위 테스트
package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_FiresOnDanglingRef_ZeroCov(t *testing.T) {
	// has_many pointing at an undefined table fires FT-10 → exercises the
	// aggregation path with non-empty diags.
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {HasMany: []string{"ghost"}},
		},
	}
	diags := Run(fs)
	if len(diags) == 0 {
		t.Fatal("expected at least one diagnostic for dangling has_many ref")
	}
}
