//ff:type feature=validate-contract type=model topic=preserve-safety
//ff:what unmarshalKind — PRV-12 내부에서 Unmarshal 호출 처리 상태 분류

package contract

// unmarshalKind classifies how an Unmarshal call's returned error is
// handled within the statement that wraps it.
type unmarshalKind int

const (
	// unmarshalKindUnknown — the call was not recognised as an
	// Unmarshal statement at all; callers skip it.
	unmarshalKindUnknown unmarshalKind = iota
	// unmarshalKindAssigned — result bound to a named err identifier
	// that must be checked later in the same block.
	unmarshalKindAssigned
	// unmarshalKindDiscarded — result discarded with `_ = ...` or
	// used as the init in `if err := ...; err != nil`, both accepted.
	unmarshalKindDiscarded
)
