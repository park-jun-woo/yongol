//ff:type feature=orchestrator type=model
//ff:what Diagnostic phase enum type
package diagnostic

// Phase identifies which stage produced the diagnostic.
type Phase string

const (
	PhaseParse    Phase = "parse"
	PhaseValidate Phase = "validate"
)
