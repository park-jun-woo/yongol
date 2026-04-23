//ff:type feature=migration type=model
//ff:what SafetyLevel — CheckSafety 가 Operation 에 붙이는 심각도 분류
package migration

// SafetyLevel is the severity Phase004 attaches to each Operation.
type SafetyLevel int

const (
	SafetySafe SafetyLevel = iota
	SafetyWarning
	SafetyError
)
