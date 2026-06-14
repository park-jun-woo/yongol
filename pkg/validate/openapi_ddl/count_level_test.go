//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=openapi-ddl
//ff:what countLevel — 진단 메시지에 주어진 prefix 가 포함된 건수 집계

package openapi_ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// countLevel counts diagnostics whose message contains the given prefix.
func countLevel(diags []diagnostic.Diagnostic, prefix string) int {
	n := 0
	for _, d := range diags {
		if strings.Contains(d.Message, prefix) {
			n++
		}
	}
	return n
}
