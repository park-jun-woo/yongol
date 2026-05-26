//ff:type feature=gen-ir type=model
//ff:what RequestValidatorConfig -- OpenAPI 요청 검증 미들웨어 설정

package ir

// RequestValidatorConfig holds the resolved request validator middleware
// configuration. Currently a marker struct — the validator reads its
// OpenAPI spec at runtime. Future phases may add configurable options
// (strict mode, custom error format, etc.).
type RequestValidatorConfig struct {
	// Active is always true when the config is present.
	Active bool
}
