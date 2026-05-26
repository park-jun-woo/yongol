//ff:type feature=gen-ir type=model
//ff:what SessionConfig -- session 인프라 어댑터 설정

package ir

// SessionConfig holds the resolved session backend configuration.
type SessionConfig struct {
	// Backend is one of "postgres" or "memory".
	Backend string
}
