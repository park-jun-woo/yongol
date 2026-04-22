//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what 단일 AST 선언에서 struct 목록을 추출
package ssac

import "go/ast"

// collectStructsFromDecl은 단일 AST GenDecl에서 struct 목록을 추출한다.
func collectStructsFromDecl(decl ast.Decl) []StructInfo {
	gd, ok := decl.(*ast.GenDecl)
	if !ok {
		return nil
	}
	var structs []StructInfo
	for _, spec := range gd.Specs {
		si := extractStructInfo(spec)
		if si != nil {
			structs = append(structs, *si)
		}
	}
	return structs
}
