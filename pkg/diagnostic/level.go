//ff:type feature=orchestrator type=model
//ff:what 진단 심각도 열거 타입
package diagnostic

// Level indicates the severity of a diagnostic.
type Level string

const (
	LevelError   Level = "ERROR"
	LevelWarning Level = "WARNING"
)
