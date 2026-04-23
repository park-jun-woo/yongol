//ff:type feature=gen-gogin type=model
//ff:what methodGen — SSaC method 코드젠 컨텍스트

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// methodGen holds all context needed to generate one StrictServerInterface method.
type methodGen struct {
	FuncName     string
	FileName     string // 원본 SSaC 파일명 (진단 메시지용)
	ModulePath   string
	PathParams   map[string]bool          // OpenAPI path param names (lowercase)
	QueryParams  map[string]queryParam    // OpenAPI query param name → rich metadata
	BodyFormats  map[string]string        // OpenAPI body field name → format ("email", "uuid", ...)
	RespFields   map[string]responseField // OpenAPI 200 response schema fields
	IsSubscribe  bool
	UseTx        bool
	FirstErr     bool // true = next err-only declaration uses :=
	ProjectFuncs []funcspec.FuncSpec
	BuiltinFuncs []funcspec.FuncSpec
	// WrapCalls, when true, wraps every @call emission with an explicit
	// otel.Tracer("ssac").Start(...) span. Sourced from
	// manifest.backend.observability.tracing.wrap_calls — off by default to
	// keep span volume manageable for high-traffic handlers. Phase009.
	WrapCalls bool
	// AccessTokenVar carries the variable name of the most recent
	// `@call auth.IssueToken`-bound result within this handler (Phase020).
	// build_call pairs it with the following `@call auth.RefreshToken`
	// assignment to emit auth.SetAuthCookies(ctx, <access>.AccessToken,
	// <refresh>.RefreshToken). Empty string means "no IssueToken seen";
	// the SetAuthCookies emission is skipped.
	AccessTokenVar string
	// DiagramSymbol maps each state diagram ID (lowercase filename stem
	// matched by SSaC `@state <id>`) to its exported PascalCase Go
	// symbol. Populated at methodGen construction from fs.StateDiagrams
	// so build_state can reference statemachine.<Symbol>CanTransition
	// even when the source .md file uses a lowercase name (BUG-002).
	DiagramSymbol map[string]string
}
