//ff:type feature=gen-ir type=model
//ff:what CORSBootConfig -- CORS 부트 블록 설정 (허용 origin, credentials)

package ir

// CORSBootConfig carries CORS configuration from manifest.yaml to the boot
// renderers. When AllowOrigins is nil, the renderer emits a permissive
// enableCors() call.
type CORSBootConfig struct {
	AllowOrigins     []string
	AllowCredentials bool
}
