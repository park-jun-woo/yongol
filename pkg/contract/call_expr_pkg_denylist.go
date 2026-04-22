//ff:type feature=contract type=model
//ff:what callExprPkgDenylist — CallTargets 집계에서 제외할 식별자 목록 (지역 변수·수신자)

package contract

// callExprPkgDenylist names identifiers that frequently appear as
// selector bases without representing an imported package. They are
// excluded from CallTargets so context-local helpers do not look like
// external dependencies.
var callExprPkgDenylist = map[string]struct{}{
	"ctx":         {},
	"err":         {},
	"req":         {},
	"res":         {},
	"resp":        {},
	"w":           {},
	"r":           {},
	"t":           {},
	"s":           {},
	"c":           {},
	"currentUser": {},
	"server":      {},
	"tx":          {},
	"db":          {},
}
