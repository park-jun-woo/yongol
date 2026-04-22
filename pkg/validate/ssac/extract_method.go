//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what extractMethod — "Model.Method" 에서 Method 부분 추출

package ssac

import (
	"strings"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// extractMethod returns the Method portion of "Model.Method" (or "" if absent).
func extractMethod(seq parsessac.Sequence) string {
	if i := strings.IndexByte(seq.Model, '.'); i > 0 && i+1 < len(seq.Model) {
		return seq.Model[i+1:]
	}
	return ""
}
