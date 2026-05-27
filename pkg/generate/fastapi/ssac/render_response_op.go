//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderResponseOp — ResponseOp → Python return 문 렌더링 (source 필드 snake_case 변환)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderResponseOp writes a return statement for the response. Source
// expressions with dot-access (e.g. "token.AccessToken") are converted to
// dict key access with snake_case keys (e.g. "token[\"access_token\"]").
func renderResponseOp(b *strings.Builder, op *ir.ResponseOp, indent string) {
	if op == nil {
		return
	}
	if op.SingleVar != "" {
		b.WriteString(fmt.Sprintf("%sreturn %s\n", indent, op.SingleVar))
		return
	}
	b.WriteString(fmt.Sprintf("%sreturn {\n", indent))
	for _, f := range op.Fields {
		src := pySourceExpr(f.Source)
		b.WriteString(fmt.Sprintf("%s    \"%s\": %s,\n", indent, snakeCase(f.Name), src))
	}
	b.WriteString(fmt.Sprintf("%s}\n", indent))
}
