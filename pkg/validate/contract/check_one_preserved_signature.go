//ff:func feature=validate-contract type=util control=sequence
//ff:what checkOnePreservedSignature — preserved 파일 1건의 signature drift 단일 Diagnostic(또는 없음) 반환

package contract

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/contract"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// checkOnePreservedSignature inspects a single preserved file and
// returns the PRV-01 Diagnostic it warrants, or an empty slice when it
// is still in contract with the SSOT. The caller iterates over all
// preserved files and flattens the results.
func checkOnePreservedSignature(g *rule.Ground, path string) []diagnostic.Diagnostic {
	opID := findOperationIDForFile(path)
	if opID == "" {
		return nil
	}
	actual, err := contract.ExtractSignature(path)
	if err != nil || actual.Name == "" {
		return nil
	}
	expected, ok := expectedSignature(g, opID)
	if !ok {
		return []diagnostic.Diagnostic{prv01MissingOpIDDiag(path, actual.Name, opID)}
	}
	diffs := compareSignature(expected, actual)
	if len(diffs) == 0 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    path,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-01] %s — preserved, but signature drifted: %s", actual.Name, strings.Join(diffs, "; ")),
		Advice: strings.Join([]string{
			"(a) revert the SSOT change that reshaped this handler",
			"(b) align the preserved body with the new signature",
			"(c) release preserve by deleting the file",
		}, "\n"),
	}}
}
