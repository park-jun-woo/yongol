//ff:func feature=validate type=util control=selection topic=ssac-sqlc
//ff:what buildXqs20Advice — XQS-20 advice 문자열 (방향별 가이드)

package ssac_sqlc

import "fmt"

// buildXqs20Advice composes the actionable hint shown alongside an XQS-20
// diagnostic. Two directions are supported:
//
//   - Model declared but RETURNING is partial → use <QueryName>Row, or expand
//     the RETURNING clause to all columns.
//   - Row declared but RETURNING is full      → use <Model>, or restrict the
//     RETURNING clause to specific columns.
func buildXqs20Advice(seqType, declared, expected string, shape ReturningShape) string {
	switch shape {
	case ShapePartial:
		return fmt.Sprintf("Change SSaC declaration to '@%s %s <var> = ...' or expand RETURNING to all columns", seqType, expected)
	case ShapeFull:
		return fmt.Sprintf("Change SSaC declaration to '@%s %s <var> = ...' or restrict RETURNING to specific columns", seqType, expected)
	default:
		return fmt.Sprintf("Change SSaC declaration to '@%s %s <var> = ...'", seqType, expected)
	}
}
