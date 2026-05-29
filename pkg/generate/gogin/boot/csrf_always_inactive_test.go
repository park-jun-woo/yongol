//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what csrfAlwaysInactive — csrf 비활성 MainBlock 용 고정 predicate

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCsrfAlwaysInactive(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{
		nil,
		{},
		{Manifest: &pmanifest.ProjectConfig{}},
	} {
		if csrfAlwaysInactive(fs) {
			t.Errorf("csrfAlwaysInactive must always return false, got true for %+v", fs)
		}
	}
}
