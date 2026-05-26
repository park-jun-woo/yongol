//ff:type feature=gen-ir type=model
//ff:what InfraPlan -- 인프라 어댑터 활성화의 프레임워크 비의존 중간 표현

package ir

// InfraPlan is the framework-agnostic intermediate representation of the
// infrastructure adapter layer. Each field represents one subsystem;
// pointer fields are nil when the subsystem is inactive.
//
// Backend renderers consume this plan to produce framework-specific
// adapter code (e.g. sqlc-based postgres adapters for Go+Gin, TypeORM
// adapters for NestJS).
type InfraPlan struct {
	// Session configures the session backend adapter. nil when no session
	// subsystem is active.
	Session *SessionConfig

	// Cache configures the cache backend adapter. nil when no cache
	// subsystem is active.
	Cache *CacheConfig

	// Queue configures the queue backend adapter. nil when no queue
	// subsystem is active.
	Queue *QueueConfig

	// Auth configures the auth infrastructure adapter (refresh_tokens DDL,
	// RefreshStore). nil when no auth is declared.
	Auth *AuthInfraConfig
}
