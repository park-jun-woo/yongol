//ff:type feature=gen-gogin type=model
//ff:what PureLinesReport — 한 루프의 body 순수 라인 수 측정 결과 (filefunc Q4 근거)

package qcheck

// PureLinesReport carries the count of non-blank, non-comment lines inside a
// loop body. filefunc Q4 requires range-body ≤ 10 pure lines; generators
// consult this before emitting to stay within budget.
type PureLinesReport struct {
	Func      string
	LoopKind  string // "for" or "range"
	Line      int    // source line of the loop header
	PureLines int
}
