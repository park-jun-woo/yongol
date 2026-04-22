//ff:func feature=contract type=util control=iteration dimension=1
//ff:what appendFieldEntries — 단일 FieldGroup 의 이름들을 FuncParam 슬라이스에 추가

package contract

import "go/ast"

// appendFieldEntries appends one FuncParam per declared name in
// field to out. A field with no names yields a single anonymous
// entry (Name == "") — for parameter lists this preserves the
// `_ int` and `int` forms, for result lists it captures unnamed
// return types.
//
// requireNames is currently advisory: the helper always falls back
// to a single anonymous entry when Names is empty. The flag is kept
// so future callers can distinguish params (where an anonymous slot
// is intentional) from results (where it is the common case).
func appendFieldEntries(out []FuncParam, field *ast.Field, typeStr string, requireNames bool) []FuncParam {
	_ = requireNames
	if len(field.Names) == 0 {
		return append(out, FuncParam{Name: "", Type: typeStr})
	}
	for _, n := range field.Names {
		out = append(out, FuncParam{Name: n.Name, Type: typeStr})
	}
	return out
}
