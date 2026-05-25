//ff:type feature=manifest type=model
//ff:what RefValue — manifest.* 참조 해석 결과 구조체

package manifest

// RefValue holds the resolved value from a manifest.* reference and its
// Go-level type so the code generator can emit the correct literal.
type RefValue struct {
	Raw    string // original YAML value (e.g. "15m")
	GoLit  string // Go literal for codegen (e.g. "900")
	GoType string // Go type name (e.g. "int64")
}
