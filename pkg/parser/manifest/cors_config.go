//ff:type feature=projectconfig type=model
//ff:what CORS 설정 구조체 (manifest backend.cors)
package manifest

import "time"

// CORSConfig mirrors the backend.cors: section in manifest.yaml. Fed to the
// generated main.go via blockCORS, which emits a gin-contrib/cors config
// literal. Environment variables (CORS_ALLOW_ORIGINS / CORS_ALLOW_METHODS /
// CORS_ALLOW_CREDENTIALS) override the static values at runtime.
type CORSConfig struct {
	Enabled          bool          `yaml:"enabled"`
	AllowOrigins     []string      `yaml:"allow_origins"`
	AllowMethods     []string      `yaml:"allow_methods"`
	AllowHeaders     []string      `yaml:"allow_headers"`
	ExposeHeaders    []string      `yaml:"expose_headers"`
	AllowCredentials bool          `yaml:"allow_credentials"`
	MaxAge           time.Duration `yaml:"max_age"`
}
