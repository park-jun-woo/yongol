//ff:type feature=gen-nestjs type=model
//ff:what tsElementType — 배열 element 의 TypeScript/Prisma 타입 쌍

package types

// tsElementType holds the resolved TypeScript and Prisma element types for
// an array column.
type tsElementType struct {
	prisma string // Prisma scalar type (e.g. "Int", "String")
	ts     string // TypeScript type (e.g. "number", "string")
}
