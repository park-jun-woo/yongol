//ff:func feature=gen-splitter type=util control=selection
//ff:what primaryTypeName — 단일 spec GenDecl 의 대표 이름(type/const/var) 추출
package splitter

import "go/ast"

// primaryTypeName returns the sole declared name of a GenDecl when it
// contains exactly one spec — otherwise "". Multi-spec blocks are grouped
// into consts.* or vars.* files by the caller.
func primaryTypeName(d *ast.GenDecl) string {
	if len(d.Specs) != 1 {
		return ""
	}
	switch s := d.Specs[0].(type) {
	case *ast.TypeSpec:
		return s.Name.Name
	case *ast.ValueSpec:
		if len(s.Names) == 1 {
			return s.Names[0].Name
		}
	}
	return ""
}
