//ff:type feature=gen-ir type=model
//ff:what ErrorEnvelopeConfig -- 에러 응답 미들웨어 설정

package ir

// ErrorEnvelopeConfig holds the resolved error envelope middleware
// configuration.
type ErrorEnvelopeConfig struct {
	// ExposeInternalError controls whether internal error details are
	// included in responses (typically true only in dev environments).
	ExposeInternalError bool
}
