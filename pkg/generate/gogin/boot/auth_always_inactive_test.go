//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what authAlwaysInactive — auth 비활성 placeholder MainBlock 용 predicate
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestAuthAlwaysInactive(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{
		nil,
		{},
		{Manifest: &pmanifest.ProjectConfig{}},
	} {
		if authAlwaysInactive(fs) {
			t.Errorf("authAlwaysInactive must always return false, got true for %+v", fs)
		}
	}
}
