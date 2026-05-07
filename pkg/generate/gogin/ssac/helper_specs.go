//ff:func feature=gen-gogin type=util control=sequence
//ff:what helperSpecs — returns the list of Server pointer/deref helper specs (filename/summary/body)
//ff:type feature=gen-gogin type=model

package ssac

// helperSpec describes one pointer/deref helper emitted as its own file.
// Kept local to the ssac package (no accessors) — callers consume the
// slice returned by helperSpecs.
type helperSpec struct {
	file string
	what string
	code string
}

// helperSpecs lists every pointer/deref helper that Server methods depend
// on. Order is stable so regeneration produces identical files.
func helperSpecs() []helperSpec {
	return []helperSpec{
		{
			file: "ptr_of.go",
			what: "ptrOf — wraps an arbitrary value T into *T (generic helper)",
			code: "func ptrOf[T any](v T) *T { return &v }\n",
		},
		{
			file: "deref_int.go",
			what: "derefInt — dereferences *int to int32 (nil→0)",
			code: "func derefInt(p *int) int32 { if p != nil { return int32(*p) }; return 0 }\n",
		},
		{
			file: "deref_str.go",
			what: "derefStr — dereferences *string to string (nil→\"\")",
			code: "func derefStr(p *string) string { if p != nil { return *p }; return \"\" }\n",
		},
		{
			file: "deref_int64.go",
			what: "derefInt64 — dereferences *int64 to int64 (nil→0)",
			code: "func derefInt64(p *int64) int64 { if p != nil { return *p }; return 0 }\n",
		},
		{
			file: "deref_bool.go",
			what: "derefBool — dereferences *bool to bool (nil→false)",
			code: "func derefBool(p *bool) bool { if p != nil { return *p }; return false }\n",
		},
		{
			file: "deref_enum.go",
			what: "derefEnum — dereferences a ~string-based enum pointer to string (nil→\"\")",
			code: "func derefEnum[T ~string](p *T) string { if p != nil { return string(*p) }; return \"\" }\n",
		},
	}
}
