//ff:func feature=validate-contract type=util control=sequence
//ff:what checkOnePreservedSymbols — preserved 파일 1건의 외부 심볼 drift Diagnostic 목록 반환

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/contract"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkOnePreservedSymbols extracts the external symbols from path and
// diffs each category against the SSOT-derived expected sets. The
// file's own imports are used to:
//
//   - suppress false positives where the DDLFields extractor captures
//     package-qualified selectors such as `sql.ErrNoRows`.
//   - reclassify `qtx.Method(...)` call targets (whose receiver is a
//     local transaction variable, not an imported package) as sqlc
//     query references.
func checkOnePreservedSymbols(path string, expQueries, expCalls, expFields map[string]bool) []diagnostic.Diagnostic {
	sym, err := contract.ExtractExternalSymbols(path)
	if err != nil {
		return nil
	}
	if len(sym.SqlcQueries)+len(sym.CallTargets)+len(sym.DDLFields) == 0 {
		return nil
	}
	pkgs := collectFileImports(path)
	sym.DDLFields = filterPackageSelectors(sym.DDLFields, pkgs)
	pkgCalls, localMethods := reclassifyCallTargets(sym.CallTargets, pkgs)
	sym.CallTargets = pkgCalls
	sym.SqlcQueries = append(sym.SqlcQueries, localMethods...)
	ms := compareExternalSymbols(sym, expQueries, expCalls, expFields)
	return flattenMissingSymbols(path, ms)
}
