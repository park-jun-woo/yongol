//ff:func feature=gen-gogin type=util control=sequence
//ff:what helperSpecs — Server pointer/deref 헬퍼 사양(파일명/요약/본문) 리스트 반환
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
			file: "str_ptr.go",
			what: "strPtr — 문자열 리터럴을 *string 으로 래핑",
			code: "func strPtr(s string) *string { return &s }\n",
		},
		{
			file: "ptr_of.go",
			what: "ptrOf — 임의 값 T 를 *T 로 래핑 (제네릭 helper)",
			code: "func ptrOf[T any](v T) *T { return &v }\n",
		},
		{
			file: "deref_int.go",
			what: "derefInt — *int 을 int32 로 역참조 (nil→0)",
			code: "func derefInt(p *int) int32 { if p != nil { return int32(*p) }; return 0 }\n",
		},
		{
			file: "deref_str.go",
			what: "derefStr — *string 을 string 으로 역참조 (nil→\"\")",
			code: "func derefStr(p *string) string { if p != nil { return *p }; return \"\" }\n",
		},
		{
			file: "deref_int64.go",
			what: "derefInt64 — *int64 을 int64 로 역참조 (nil→0)",
			code: "func derefInt64(p *int64) int64 { if p != nil { return *p }; return 0 }\n",
		},
		{
			file: "deref_bool.go",
			what: "derefBool — *bool 을 bool 로 역참조 (nil→false)",
			code: "func derefBool(p *bool) bool { if p != nil { return *p }; return false }\n",
		},
		{
			file: "deref_enum.go",
			what: "derefEnum — ~string 파생 enum pointer 를 문자열로 역참조 (nil→\"\")",
			code: "func derefEnum[T ~string](p *T) string { if p != nil { return string(*p) }; return \"\" }\n",
		},
	}
}
