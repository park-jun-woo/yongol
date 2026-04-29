//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ssac-sqlc
//ff:what runXQS20 — XQS-20 디스패처 호출 후 메시지 슬라이스로 환원 (테스트 헬퍼)

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/yongol"

// runXQS20 runs the dispatcher and returns the diagnostic messages as a
// slice. Used by every XQS-20 test case to keep assertions terse.
func runXQS20(fs *yongol.Fullstack) []string {
	diags := xqs20ReturnTypeMatch(fs)
	out := make([]string, 0, len(diags))
	for _, d := range diags {
		out = append(out, d.Message)
	}
	return out
}
