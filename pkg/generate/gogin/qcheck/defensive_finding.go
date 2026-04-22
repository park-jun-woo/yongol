//ff:type feature=gen-gogin type=model
//ff:what DefensiveFinding — AST 스캐너가 보고하는 방어 패턴 위반 (파일:라인 + 카테고리)

package qcheck

import "go/token"

// DefensiveFinding is a single DF-* category violation produced by the
// defensive pattern scanners (scanUncheckedUnmarshal / scanUncheckedScan /
// scanMissingDeferClose). Category is the DF-NN identifier, Detail carries
// a human readable call expression like "json.Unmarshal" or "db.QueryContext"
// so fixture assertions can pinpoint which template emitted the regression.
type DefensiveFinding struct {
	Category string
	Detail   string
	Pos      token.Position
}
