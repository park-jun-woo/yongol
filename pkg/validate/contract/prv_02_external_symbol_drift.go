//ff:func feature=validate-contract type=rule control=iteration dimension=1
//ff:what prv02ExternalSymbolDrift — preserved 파일의 sqlc / @call / DDL 참조가 SSOT 에 없을 때 ERROR

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// prv02ExternalSymbolDrift checks every external symbol a preserved
// function still references against the SSOT-derived expected sets.
// Missing entries turn into one Diagnostic per category per file so
// the user sees the full drift surface in a single pass.
//
// Per-file logic lives in checkOnePreservedSymbols so this
// orchestrator stays flat and within Q4 body-length limits.
func prv02ExternalSymbolDrift(fs *yongol.Fullstack, preservedPaths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	expQueries, expCalls, expFields := expectedExternalSymbols(fs, fs.Ground())
	for _, path := range preservedPaths {
		diags = append(diags, checkOnePreservedSymbols(path, expQueries, expCalls, expFields)...)
	}
	return diags
}
