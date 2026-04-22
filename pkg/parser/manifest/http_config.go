//ff:type feature=projectconfig type=model
//ff:what HTTPConfig 구조체 — backend.http (body/multipart/header limit + per-operationId override)
package manifest

// HTTPConfig mirrors the backend.http: section of manifest.yaml. It controls
// HTTP-level DoS guards injected as gin middleware in the generated backend.
//
// Values are human-readable size strings ("1MiB", "32MiB"). Parsing happens
// at generation time (pkg/generate/gogin/middleware.ParseSize) and at
// runtime via the same helper emitted into env helper files.
//
// Overrides map operationId → per-endpoint limits; typically used to relax
// the global body_limit for upload endpoints.
type HTTPConfig struct {
	BodyLimit      string                     `yaml:"body_limit"`
	MultipartLimit string                     `yaml:"multipart_limit"`
	HeaderLimit    string                     `yaml:"header_limit"`
	Overrides      map[string]HTTPOverride    `yaml:"overrides"`
}

// HTTPOverride is a per-operationId override. Either or both fields may be
// set; unset fields fall back to the global limits.
type HTTPOverride struct {
	BodyLimit      string `yaml:"body_limit"`
	MultipartLimit string `yaml:"multipart_limit"`
}
