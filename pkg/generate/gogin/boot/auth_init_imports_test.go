//ff:func feature=gen-gogin type=test control=sequence
//ff:what authInitImports — blockAuthInit 가 필요한 Imports 슬라이스 구성

package boot

import (
	"strings"
	"testing"
)

func TestAuthInitImports(t *testing.T) {
	imports := authInitImports("example.com/zenflow")
	body := strings.Join(imports, "\n")
	for _, must := range []string{
		`"github.com/park-jun-woo/ssac/pkg/auth"`,
		`"net/http"`,
		`"time"`,
		`infraauth "example.com/zenflow/internal/infra/auth"`,
	} {
		if !strings.Contains(body, must) {
			t.Errorf("authInitImports missing %q, got:\n%s", must, body)
		}
	}
}
