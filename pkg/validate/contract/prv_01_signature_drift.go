//ff:func feature=validate-contract type=rule control=iteration dimension=1
//ff:what prv01SignatureDrift — preserved 파일 FuncSignature 와 SSOT 기대값 비교, 차이 시 ERROR

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// prv01SignatureDrift compares the FuncSignature of each preserved file
// against the SSOT-derived expected signature (Ground-based). Per-file
// logic lives in checkOnePreservedSignature so this orchestrator stays
// flat and within Q4 body-length limits.
func prv01SignatureDrift(fs *yongol.Fullstack, preservedPaths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	g := fs.Ground()
	for _, path := range preservedPaths {
		diags = append(diags, checkOnePreservedSignature(g, path)...)
	}
	return diags
}
