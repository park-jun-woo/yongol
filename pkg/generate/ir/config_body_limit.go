//ff:type feature=gen-ir type=model
//ff:what BodyLimitConfig -- body/multipart 크기 제한 설정

package ir

// BodyLimitConfig holds the resolved HTTP body size limits.
type BodyLimitConfig struct {
	// BodyLimit is the maximum request body size in bytes.
	BodyLimit int64

	// MultipartLimit is the maximum multipart form size in bytes.
	MultipartLimit int64

	// BodyOverrides maps route paths to per-route body limit overrides.
	BodyOverrides map[string]int64

	// MultipartOverrides maps route paths to per-route multipart limit
	// overrides.
	MultipartOverrides map[string]int64
}
