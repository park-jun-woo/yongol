//ff:func feature=validate type=util control=selection topic=ssac-sqlc
//ff:what formatReturningReason — XQS-20 진단 메시지의 reason 단편 포맷

package ssac_sqlc

import "fmt"

// formatReturningReason formats the human-readable reason fragment included
// in the XQS-20 diagnostic message for a given RETURNING shape and the raw
// captured clause string.
func formatReturningReason(shape ReturningShape, clause string) string {
	switch shape {
	case ShapeFull:
		if clause == "*" {
			return "RETURNING * → model"
		}
		return "RETURNING <full column list> → model"
	case ShapePartial:
		return fmt.Sprintf("RETURNING %s → partial Row", clause)
	default:
		return "no RETURNING → model"
	}
}
