//ff:type feature=contract type=model
//ff:what PreservedState — generated artifact 의 preserve 상태 enum

package contract

// PreservedState classifies a generated artifact by comparing the
// `//ff:checked hash=<saved>` annotation against a freshly recomputed
// body hash.
type PreservedState int

const (
	// StateNotApplicable indicates the file has no `//ff:checked`
	// annotation (type-only file, user-authored func spec, external
	// artifact without yongol provenance).
	StateNotApplicable PreservedState = iota
	// StateUntouched indicates the saved hash matches the recomputed
	// one — the file is still in its generated form.
	StateUntouched
	// StatePreserved indicates the saved hash differs from the
	// recomputed one — the user has edited the generated body.
	StatePreserved
)
