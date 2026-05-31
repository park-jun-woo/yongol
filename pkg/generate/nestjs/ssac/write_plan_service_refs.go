//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writePlanServiceRefs — providers/exports 배열에 plan 별 Service 참조 항목 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writePlanServiceRefs writes one service reference entry per plan, used by both
// the providers and exports arrays.
func writePlanServiceRefs(b *strings.Builder, plans []*ir.ServicePlan) {
	for _, p := range plans {
		b.WriteString(fmt.Sprintf("    %sService,\n", p.OperationID))
	}
}
