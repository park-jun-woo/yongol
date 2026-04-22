//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what AST TypeSpec에서 StructInfo를 추출
package ssac

import "go/ast"

// extractStructInfo는 AST TypeSpec에서 StructInfo를 추출한다.
func extractStructInfo(spec ast.Spec) *StructInfo {
	ts, ok := spec.(*ast.TypeSpec)
	if !ok {
		return nil
	}
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return nil
	}
	si := &StructInfo{Name: ts.Name.Name}
	for _, field := range st.Fields.List {
		if len(field.Names) > 0 {
			si.Fields = append(si.Fields, StructField{
				Name: field.Names[0].Name,
				Type: exprToString(field.Type),
			})
		}
	}
	return si
}
