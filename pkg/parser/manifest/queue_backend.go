//ff:type feature=projectconfig type=model
//ff:what 큐 백엔드 설정 구조체
package manifest

// QueueBackend configures queue backend (postgres | memory).
type QueueBackend struct {
	Backend string `yaml:"backend"` // "postgres" or "memory"
}
