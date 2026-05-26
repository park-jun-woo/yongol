//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what hasActiveBlock — BootPlan 에서 지정된 이름의 활성 블록 존재 여부 확인

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// hasActiveBlock checks if the plan has an active block with the given name.
func hasActiveBlock(plan *ir.BootPlan, name string) bool {
	for _, block := range plan.ActiveBlocks {
		if block.Name == name && block.Active {
			return true
		}
	}
	return false
}
