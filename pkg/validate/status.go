//ff:type feature=validate type=model
//ff:what Status — validation step 결과 상태 열거 타입
package validate

// Status represents the outcome of a single validation step.
type Status int

const (
	StatusPass Status = iota
	StatusFail
	StatusSkip
	StatusMissing
)
