//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitImports — blockAuthInit 가 필요한 Imports 슬라이스 구성

package boot

import "fmt"

// authInitImports returns the imports blockAuthInit requires. Isolated so
// blockAuthInit stays a short orchestration that does no string literal
// composition of its own.
//
// Phase001 UserClaimUnification — auth import now targets ssac/pkg/auth
// directly; the project-local internal/auth package is no longer generated.
//
// Phase002 (ssac/purify) — the yongol-generated postgres RefreshStore lives
// at `<module>/internal/infra/auth` and is imported as `infraauth`. The
// block installs it at boot via auth.Init(infraauth.NewPostgres(queries)).
func authInitImports(modulePath string) []string {
	return []string{
		`"github.com/park-jun-woo/ssac/pkg/auth"`,
		`"net/http"`,
		`"time"`,
		fmt.Sprintf(`infraauth "%s/internal/infra/auth"`, modulePath),
	}
}
