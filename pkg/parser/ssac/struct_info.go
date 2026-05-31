//ff:type feature=ssac-parse type=model
//ff:what StructInfo — type representing a Go struct declared in a .ssac file
package ssac

// StructInfo holds metadata for a Go struct declared in a .ssac file.
type StructInfo struct {
	Name   string // "OnOrderCompletedMessage"
	Fields []StructField
}
