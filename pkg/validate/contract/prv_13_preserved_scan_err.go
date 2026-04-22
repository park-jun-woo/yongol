//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv13PreservedScanErr — preserved 파일에서 sql.Scan 에러 무시 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv13PreservedScanErr runs PRV-13 against every preserved path.
// Each file is parsed once by scanFileForScanErr which walks blocks
// looking for `.Scan(...)` calls whose returned error is discarded.
func prv13PreservedScanErr(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForScanErr(path)...)
	}
	return diags
}
