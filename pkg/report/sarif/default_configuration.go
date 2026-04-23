//ff:type feature=report type=model topic=sarif
//ff:what DefaultConfiguration — 규칙 기본 severity (rulebook Level → SARIF level)
package sarif

// DefaultConfiguration carries the rule's default severity, mapped from
// the rulebook `Level` column (ERROR → "error", WARNING → "warning").
type DefaultConfiguration struct {
	Level string `json:"level,omitempty"`
}
