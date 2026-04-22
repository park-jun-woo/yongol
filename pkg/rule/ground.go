//ff:type feature=rule type=model
//ff:what Ground — 검증 규칙이 공유하는 조회 컨텍스트
package rule

// Ground holds lookup data for validation rules.
// Populated by the caller, stored in ctx via ctx.Set("ground", g).
type Ground struct {
	Lookup  map[string]StringSet // "target.kind" -> set of names
	Types   map[string]string    // "target.kind.name" -> type string
	Pairs   map[string]StringSet // "target.pairKind" -> set of "key:value"
	Config  map[string]bool      // config key -> present
	Vars    StringSet            // declared variable names
	Flags   StringSet            // flags for defeaters
	Schemas map[string][]string  // "target.schema" -> ordered field list
}
