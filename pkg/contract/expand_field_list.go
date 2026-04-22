//ff:func feature=contract type=util control=iteration dimension=1
//ff:what expandFieldList — FieldList 를 FuncParam 슬라이스로 전개 (다중 이름 그룹 분해)

package contract

import (
	"go/ast"
	"go/token"
)

// expandFieldList converts an ast.FieldList (parameters or results)
// into a flat []FuncParam. A group like `a, b int` expands to two
// entries sharing the same type string; an anonymous field (results
// only in normal Go, but also valid as unnamed params) yields a
// single entry with Name == "".
//
// When requireNames is false and a field has no declared names, one
// entry is appended with the type only — this matches how results
// like `func Foo() int` report a single unnamed result.
func expandFieldList(fset *token.FileSet, list *ast.FieldList, requireNames bool) []FuncParam {
	if list == nil {
		return nil
	}
	var out []FuncParam
	for _, field := range list.List {
		typeStr := printType(fset, field.Type)
		out = appendFieldEntries(out, field, typeStr, requireNames)
	}
	return out
}
