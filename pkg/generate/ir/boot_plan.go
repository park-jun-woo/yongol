//ff:type feature=gen-ir type=model
//ff:what BootPlan -- 부트스트랩 블록 활성화의 프레임워크 비의존 중간 표현

package ir

// BootPlan is the framework-agnostic intermediate representation of the
// main.go bootstrap sequence. It captures which initialization blocks are
// active and their configuration, independent of the target backend
// framework (Go+Gin, NestJS, FastAPI).
//
// Backend renderers consume the pre-evaluated ActiveBlocks list. The
// activation logic (Pattern A/B/C/D) is already resolved by
// BuildBootPlan so renderers only need to emit framework-specific code
// for active blocks.
type BootPlan struct {
	// ProjectID is the project name from manifest.metadata.name.
	ProjectID string

	// Module is the Go module path from manifest.backend.module.
	Module string

	// BackendType identifies the target backend framework as a string
	// (e.g. "go-gin", "nestjs", "fastapi"). Typed as string to avoid
	// an import cycle with the generate package.
	BackendType string

	// ActiveBlocks lists every bootstrap block with its activation state
	// already evaluated. Blocks where Active==false may still appear for
	// informational purposes (e.g. diagnostics) but renderers skip them.
	ActiveBlocks []BootBlock

	// ErrorEnvelope specifies the standard error response shape that all
	// backend renderers must produce. nil when no error envelope is
	// configured (every project gets one by default).
	ErrorEnvelope *ErrorEnvelopeSpec
}
