//ff:type feature=gen-ir type=model
//ff:what CacheConfig -- cache 인프라 어댑터 설정

package ir

// CacheConfig holds the resolved cache backend configuration.
type CacheConfig struct {
	// Backend is one of "postgres" or "memory".
	Backend string
}
