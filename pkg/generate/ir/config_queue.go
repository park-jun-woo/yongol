//ff:type feature=gen-ir type=model
//ff:what QueueConfig -- queue 인프라 어댑터 설정

package ir

// QueueConfig holds the resolved queue backend configuration.
type QueueConfig struct {
	// Backend is one of "postgres" or "memory".
	Backend string
}
