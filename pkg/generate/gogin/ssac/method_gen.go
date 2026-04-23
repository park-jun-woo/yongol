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
	// SuccessStatus is the 2xx status code used for the generated success
	// response. Derived from the OpenAPI operation's HTTP method + declared
	// responses via openapi.DeriveSuccessStatus (BUG-004). Defaults to 200
	// so callers that never invoked extractFromOpenAPI keep emitting the
	// historical status, but real operations populate this during
	// newMethodGen → extractFromOpenAPI so build_response and
	// build_field_response produce `api.<Op><Code>JSONResponse` with the
	// correct code.
	SuccessStatus int
	// Method is the HTTP verb (GET/POST/PUT/PATCH/DELETE) of the operation,
	// populated alongside SuccessStatus. Used by diagnostics that reference
	// the source-of-truth HTTP method.
	Method string
	// VarTypes maps an SSaC result variable name (e.g. "updated", "act")
	// to the declared sqlc row type (e.g. "Workflow", "Action"). Populated
	// when each sequence with a Result binding is visited. buildResponse
	// consults this map so it can route `@response <var>` through
	// convert<Model>(<var>) instead of casting the sqlc row to the api
	// DTO — the two types diverge on acronym casing and JSONB encoding,
	// so the conversion helper is mandatory (BUG-003 / BUG-005).
	VarTypes map[string]string
}
