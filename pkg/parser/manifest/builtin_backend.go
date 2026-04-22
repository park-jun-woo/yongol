//ff:type feature=projectconfig type=model
//ff:what 세션/캐시 백엔드 설정 구조체
package manifest

// BuiltinBackend configures session/cache backend (postgres | memory).
type BuiltinBackend struct {
	Backend string `yaml:"backend"` // "postgres" or "memory"
}
