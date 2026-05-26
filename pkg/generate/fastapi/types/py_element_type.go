//ff:type feature=gen-fastapi type=model
//ff:what pyElementType — 배열 element 의 Python/SQLAlchemy 타입 쌍

package types

// pyElementType holds the resolved Python and SQLAlchemy element types for
// an array column.
type pyElementType struct {
	sa string // SQLAlchemy scalar type (e.g. "Integer", "String")
	py string // Python type (e.g. "int", "str")
}
