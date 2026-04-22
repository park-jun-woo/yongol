//ff:func feature=validate-contract type=util control=sequence
//ff:what resolveImportName — import spec의 alias 또는 경로 basename으로 패키지 식별자 결정

package contract

import (
	"go/ast"
	"path"
	"strings"
)

// resolveImportName picks the effective package identifier for an
// import: the explicit alias when set (and not `_` / `.`), otherwise
// the last segment of the quoted import path. The quoted form — as
// stored in go/ast — is stripped of its surrounding double quotes
// before basename extraction.
func resolveImportName(aliasIdent *ast.Ident, quotedPath string) string {
	if aliasIdent != nil && aliasIdent.Name != "" && aliasIdent.Name != "_" && aliasIdent.Name != "." {
		return aliasIdent.Name
	}
	unquoted := strings.Trim(quotedPath, "\"")
	return path.Base(unquoted)
}
