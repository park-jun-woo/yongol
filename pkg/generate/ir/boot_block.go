//ff:type feature=gen-ir type=model
//ff:what BootBlock -- 단일 부트스트랩 블록의 이름 + 활성 상태 + 설정

package ir

// BootBlock represents a single initialization block in the bootstrap
// sequence. Each block corresponds to one subsystem (db_init, cors,
// session, otel, etc.). The Active field is pre-evaluated by
// BuildBootPlan; renderers branch on it without re-checking manifest
// conditions.
//
// Config holds a block-specific typed struct (ObservabilityConfig,
// CORSConfig, etc.) accessed via type assertion. nil for always-active
// blocks that need no configuration (logger, env_helpers, router, etc.).
type BootBlock struct {
	// Name uniquely identifies the block (e.g. "db_init", "cors",
	// "session", "otel-init", "auth-init").
	Name string

	// Active is true when the block should emit code for the current
	// project configuration.
	Active bool

	// Config carries block-specific typed configuration. Renderers use
	// type assertion to extract the concrete struct:
	//
	//   if cfg, ok := block.Config.(*ir.CORSBootConfig); ok { ... }
	//
	// nil for blocks that are unconditionally active or need no config.
	Config any
}
