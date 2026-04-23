//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitImports — blockAuthInit 가 필요한 Imports 슬라이스 구성

package boot

// authInitImports returns the imports blockAuthInit requires. Isolated so
// blockAuthInit stays a short orchestration that does no string literal
// composition of its own.
//
// Phase001 UserClaimUnification — auth import now targets ssac/pkg/auth
// directly; the project-local internal/auth package is no longer generated.
// The modulePath argument is retained so the signature is stable for the
// boot orchestration layer but is intentionally unused here.
func authInitImports(modulePath string) []string {
	_ = modulePath
	return []string{
		`"github.com/park-jun-woo/ssac/pkg/auth"`,
		`"net/http"`,
		`"time"`,
	}
}
