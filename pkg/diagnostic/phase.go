//ff:type feature=orchestrator type=model
//ff:what 진단 단계 열거 타입
package diagnostic

// Phase identifies which stage produced the diagnostic.
type Phase string

const (
	PhaseParse    Phase = "parse"
	PhaseValidate Phase = "validate"
)
