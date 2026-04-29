//ff:func feature=validate type=test-helper control=sequence topic=ssac-sqlc
//ff:what makeSeq — XQS-20 테스트용 SSaC Sequence 생성 헬퍼 (Result 바인딩 포함)

package ssac_sqlc

import ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// makeSeq is a small constructor helper for SSaC sequences with a result
// binding, used by every XQS-20 test case.
func makeSeq(seqType, declaredType, model string) ssacparser.Sequence {
	return ssacparser.Sequence{
		Type:   seqType,
		Model:  model,
		Result: &ssacparser.Result{Type: declaredType, Var: "v"},
		Line:   42,
	}
}
