//ff:type feature=projectconfig type=model
//ff:what HTTPConfig — backend.http 섹션 모델 (body/multipart/header limit + 오버라이드 맵)

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
	BodyLimit      string                  `yaml:"body_limit"`
	MultipartLimit string                  `yaml:"multipart_limit"`
	HeaderLimit    string                  `yaml:"header_limit"`
	Overrides      map[string]HTTPOverride `yaml:"overrides"`
}
