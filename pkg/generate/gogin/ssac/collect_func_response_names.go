//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFuncResponseNames — ServiceFuncs 에서 @call Func Response 타입 이름 + full import path 수집

package ssac

import (
	"path"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// funcRespInfo holds the package alias and full import path for a Func
// Response type collected from @call sequences.
type funcRespInfo struct {
	PkgAlias   string // "dashboard"
	ImportPath string // "github.com/park-jun-woo/zenflow/internal/dashboard"
}

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
			// For @call, the package alias is the first element of Model
			// (e.g. "dashboard.Summarize" → "dashboard").
			pkgAlias := ""
			if parts := strings.SplitN(seq.Model, ".", 2); len(parts) == 2 {
				pkgAlias = parts[0]
			}
			importPath := ""
			for _, imp := range sf.Imports {
				if path.Base(imp) == pkgAlias {
					importPath = imp
					break
				}
			}
			result[typeName] = funcRespInfo{
				PkgAlias:   pkgAlias,
				ImportPath: importPath,
			}
		}
	}
	return result
}
