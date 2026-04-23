//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitImports — blockAuthInit 가 필요한 Imports 슬라이스 구성

package boot

// authInitImports returns the imports blockAuthInit requires. Isolated so
// blockAuthInit stays a short orchestration that does no string literal
// composition of its own.
func authInitImports(modulePath string) []string {
	return []string{
		`"` + modulePath + `/internal/auth"`,
		`"net/http"`,
		`"time"`,
	}
}
