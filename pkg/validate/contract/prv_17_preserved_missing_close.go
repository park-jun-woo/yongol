//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv17PreservedMissingClose — preserved 파일에서 Close 필요한 리소스 획득 후 defer 누락 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv17PreservedMissingClose runs PRV-17 against every preserved
// path. Per-file scanning is delegated to scanFileForMissingClose
// which looks for assignments whose RHS is a resource-returning call
// (`os.Open`, `db.Query`, `http.Get`, ...) and verifies the same
// function body contains a matching `defer <var>.Close()`.
func prv17PreservedMissingClose(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForMissingClose(path)...)
	}
	return diags
}
