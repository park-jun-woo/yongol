//ff:func feature=gen-nestjs type=util control=selection
//ff:what arrayElementTSType — 배열 element head → TypeScript/Prisma element 타입 (ClassifyFamily 위임)

package types

import "github.com/park-jun-woo/yongol/pkg/generate/typemap"

// arrayElementTSType maps a PG array element head token to its TypeScript
// and Prisma element types by delegating to typemap.ClassifyFamily.
// Returns ok=false for elements outside the four natively-supported array
// families (Integer / Float / String / Boolean).
func arrayElementTSType(elementHead string) (tsElementType, bool) {
	fam := typemap.ClassifyFamily(elementHeadMeta{head: elementHead})
	switch fam {
	case typemap.FamilyInteger:
		return tsElementType{prisma: "Int", ts: "number"}, true
	case typemap.FamilyFloat:
		return tsElementType{prisma: "Float", ts: "number"}, true
	case typemap.FamilyString:
		return tsElementType{prisma: "String", ts: "string"}, true
	case typemap.FamilyBoolean:
		return tsElementType{prisma: "Boolean", ts: "boolean"}, true
	}
	return tsElementType{}, false
}
