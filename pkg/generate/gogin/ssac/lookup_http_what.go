//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what lookupHTTPWhat — OpenAPI summary/description 에서 핸들러 한 줄 설명 추출

package ssac

import (
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// lookupHTTPWhat returns a short one-line description for a service handler.
// Preference order: OpenAPI operation.Summary → OpenAPI operation.Description →
// "HTTP handler" fallback. Used as the //ff:what body.
func lookupHTTPWhat(fs *yongol.Fullstack, sf ssacparser.ServiceFunc) string {
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return "HTTP handler"
	}
	for _, item := range fs.OpenAPIDoc.Paths.Map() {
		for _, op := range item.Operations() {
			if op == nil || op.OperationID != sf.Name {
				continue
			}
			if s := strings.TrimSpace(op.Summary); s != "" {
				return firstLine(s)
			}
			if s := strings.TrimSpace(op.Description); s != "" {
				return firstLine(s)
			}
		}
	}
	return "HTTP handler"
}
