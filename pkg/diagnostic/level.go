//ff:type feature=orchestrator type=model
//ff:what Diagnostic severity enum type
package diagnostic

// Level indicates the severity of a diagnostic.
type Level string

const (
	LevelError   Level = "ERROR"
	LevelWarning Level = "WARNING"
)
