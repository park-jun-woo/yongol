//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what writeDiagErrors — 진단 에러를 문자열 빌더에 기록

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func writeDiagErrors(b *strings.Builder, diags []diagnostic.Diagnostic) {
	b.WriteString("Validate errors:\n")
	for _, d := range diags {
		fmt.Fprintf(b, "- %s\n", d.Message)
		if d.Advice != "" {
			fmt.Fprintf(b, "  Advice: %s\n", d.Advice)
		}
	}
}
