//ff:type feature=generate type=test-helper
//ff:what prepareSessionBackendCase — TestPrepareSessionBackend 의 단일 테이블 엔트리

package prepared

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// prepareSessionBackendCase captures one row of the 2x2 activation
// matrix (manifest declared × SSaC used).
type prepareSessionBackendCase struct {
	name     string
	manifest *pmanifest.ProjectConfig
	funcs    []ssac.ServiceFunc
	wantNil  bool
	wantBE   string
}
