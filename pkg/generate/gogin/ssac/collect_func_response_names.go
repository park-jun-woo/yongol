//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what collectFuncResponseNames — ServiceFuncs 에서 @call Func Response 타입 이름 + full import path 수집

package ssac

import (
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectFuncResponseNames scans all ServiceFuncs for @call sequences that
// produce a typed Result. For each such sequence it extracts:
//   - Result.Type → the Func Response type name (e.g. "SummarizeResponse")
//   - Sequence.Model (split on ".") → the package alias (e.g. "dashboard")
//   - ServiceFunc.Imports → the full import path whose path.Base matches the alias
//
// For @call sequences, Package is not populated by the parser — the package
// is encoded in Model as "pkg.Func". We split on "." to extract the alias.
//
// The returned map is keyed by type name → funcRespInfo.
func collectFuncResponseNames(serviceFuncs []ssacparser.ServiceFunc) map[string]funcRespInfo {
	result := make(map[string]funcRespInfo)
	for _, sf := range serviceFuncs {
		for _, seq := range sf.Sequences {
			if seq.Type != "call" || seq.Result == nil || seq.Result.Type == "" {
				continue
			}
			typeName := seq.Result.Type
			if _, exists := result[typeName]; exists {
				continue
			}
			result[typeName] = extractFuncRespInfo(seq, sf.Imports)
		}
	}
	return result
}
