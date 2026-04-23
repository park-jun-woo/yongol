//ff:type feature=projectconfig type=model
//ff:what ErrorConfig — backend.error 섹션 모델 (envelope 포맷 + request_id 정책)

package manifest

// ErrorConfig mirrors the backend.error: section of manifest.yaml. Drives the
// Phase004 error-envelope + request_id middlewares emitted into the generated
// backend. When the block is absent yongol applies sensible defaults — every
// field has a documented zero-value meaning so projects can leave the block
// out entirely and still receive the envelope + ULID request ids.
//
// Env overrides (resolved in generated main.go via envBool / envString):
//
//	BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM=true|false
//	BACKEND_ERROR_REQUEST_ID_HEADER="X-Request-Id"
//	BACKEND_ERROR_DEFAULT_LOCALE="ko"
//	BACKEND_ERROR_EXPOSE_INTERNAL_ERROR=true|false
type ErrorConfig struct {
	RequestID           *RequestIDConfig `yaml:"request_id,omitempty"`
	DefaultLocale       string           `yaml:"default_locale,omitempty"`
	ExposeInternalError *bool            `yaml:"expose_internal_error,omitempty"`
}
