//ff:type feature=validate-contract type=model topic=preserve-safety
//ff:what scanKind — PRV-13 내부 Scan 호출 처리 상태 분류

package contract

// scanKind classifies how a `.Scan(...)` call's returned error is
// handled within its enclosing statement.
type scanKind int

const (
	// scanKindUnknown — the statement is not a Scan call at all.
	scanKindUnknown scanKind = iota
	// scanKindAssigned — result bound to a named err identifier that
	// must be guarded later in the same block.
	scanKindAssigned
	// scanKindDiscarded — result explicitly dropped with `_ = ...` or
	// consumed by the init of an `if ...; err != nil` statement.
	scanKindDiscarded
)
