//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what extractModel — "Model.Method" 에서 Model 부분 추출

package ssac

import (
	"strings"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// extractModel returns the Model portion of "Model.Method" (or "" if absent).
func extractModel(seq parsessac.Sequence) string {
	if i := strings.IndexByte(seq.Model, '.'); i > 0 {
		return seq.Model[:i]
	}
	return ""
}
