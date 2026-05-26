//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what hasPublishOp — Op 배열에 Publish 연산 포함 여부 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// hasPublishOp returns true if any op is a publish.
func hasPublishOp(ops []ir.Op) bool {
	for _, op := range ops {
		if op.Kind == ir.OpPublish {
			return true
		}
	}
	return false
}
