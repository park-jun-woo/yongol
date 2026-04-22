//ff:type feature=ssac-parse type=model
//ff:what StructField — type representing a struct field
package ssac

// StructField holds metadata for a single struct field.
type StructField struct {
	Name string // "OrderID"
	Type string // "int64"
}
