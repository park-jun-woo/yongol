//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildRequestValidatorConfig -- RequestValidatorConfig 생성 (항상 활성)

package ir

// buildRequestValidatorConfig always returns a config.
func buildRequestValidatorConfig() *RequestValidatorConfig {
	return &RequestValidatorConfig{Active: true}
}
