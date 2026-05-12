//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=response
//ff:what extractFuncRespInfo — @call 시퀀스에서 패키지 별칭 + import 경로를 추출하여 funcRespInfo 반환

package ssac

import (
	"path"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// extractFuncRespInfo extracts the package alias and import path from a
// @call sequence. The package alias is derived from Model (e.g.
// "dashboard.Summarize" → "dashboard") and the import path is found in
// the ServiceFunc.Imports whose path.Base matches the alias.
func extractFuncRespInfo(seq ssacparser.Sequence, imports []string) funcRespInfo {
	pkgAlias := ""
	if parts := strings.SplitN(seq.Model, ".", 2); len(parts) == 2 {
		pkgAlias = parts[0]
	}
	importPath := ""
	for _, imp := range imports {
		if path.Base(imp) == pkgAlias {
			importPath = imp
			break
		}
	}
	return funcRespInfo{
		PkgAlias:   pkgAlias,
		ImportPath: importPath,
	}
}
