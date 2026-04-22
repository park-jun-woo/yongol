//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv12PreservedUnmarshalErr — preserved 파일에서 json/yaml Unmarshal 에러 무시 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv12PreservedUnmarshalErr runs PRV-12 against every preserved path.
// Per-file scanning is delegated to scanFileForUnmarshalErr so each
// file is parsed once and walked statement-by-statement to catch the
// ignored-error pattern.
func prv12PreservedUnmarshalErr(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForUnmarshalErr(path)...)
	}
	return diags
}
