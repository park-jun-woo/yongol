//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what compareExternalSymbols — preserved 외부 심볼 중 SSOT 에 없는 항목만 카테고리별로 분리 반환

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/contract"
)

// compareExternalSymbols returns the preserved-but-unresolved references
// in each category. Expected sets come from expectedExternalSymbols. An
// empty result means the preserved body is internally consistent with
// the current SSOT surface.
//
// CallTargets are gated by the SSOT's known-package set so std-library
// and unrelated-package calls do not pollute the drift output — only
// calls whose package appears in expectedCalls are subject to the
// drift rule.
func compareExternalSymbols(actual contract.ExternalSymbols, expectedQueries, expectedCalls, expectedFields map[string]bool) missingSymbols {
	var ms missingSymbols
	knownPkgs := knownCallPackages(expectedCalls)
	for _, q := range actual.SqlcQueries {
		if queryMethodDenylist[q] {
			continue
		}
		if !expectedQueries[q] {
			ms.Queries = append(ms.Queries, q)
		}
	}
	for _, c := range actual.CallTargets {
		if !callPkgIsKnown(c, knownPkgs) {
			continue
		}
		if !expectedCalls[c] {
			ms.Calls = append(ms.Calls, c)
		}
	}
	for _, f := range actual.DDLFields {
		if fieldIsDDLDrift(f, expectedFields) {
			ms.Fields = append(ms.Fields, f)
		}
	}
	return ms
}
