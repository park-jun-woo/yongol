//ff:type feature=gen-gogin type=model topic=response
//ff:what funcRespInfo — @call Func Response 패키지 별칭 + import 경로 보유 모델

package ssac

// funcRespInfo holds the package alias and full import path for a Func
// Response type collected from @call sequences.
type funcRespInfo struct {
	PkgAlias   string // "dashboard"
	ImportPath string // "github.com/park-jun-woo/zenflow/internal/dashboard"
}
