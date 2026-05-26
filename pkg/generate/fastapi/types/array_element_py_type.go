//ff:func feature=gen-fastapi type=util control=selection
//ff:what arrayElementPyType — 배열 element head → Python/SQLAlchemy element 타입 (ClassifyFamily 위임)

package types

import "github.com/park-jun-woo/yongol/pkg/generate/typemap"

// arrayElementPyType maps a PG array element head token to its Python
// and SQLAlchemy element types by delegating to typemap.ClassifyFamily.
// Returns ok=false for elements outside the four natively-supported array
// families (Integer / Float / String / Boolean).
func arrayElementPyType(elementHead string) (pyElementType, bool) {
	fam := typemap.ClassifyFamily(elementHeadMeta{head: elementHead})
	switch fam {
	case typemap.FamilyInteger:
		return pyElementType{sa: "Integer", py: "int"}, true
	case typemap.FamilyFloat:
		return pyElementType{sa: "Float", py: "float"}, true
	case typemap.FamilyString:
		return pyElementType{sa: "String", py: "str"}, true
	case typemap.FamilyBoolean:
		return pyElementType{sa: "Boolean", py: "bool"}, true
	}
	return pyElementType{}, false
}
