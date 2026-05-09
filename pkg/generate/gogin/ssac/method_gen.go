//ff:type feature=gen-gogin type=model
//ff:what methodGen — SSaC method 코드젠 컨텍스트

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// methodGen holds all context needed to generate one StrictServerInterface method.
type methodGen struct {
	FuncName     string
	FileName     string // 원본 SSaC 파일명 (진단 메시지용)
	ModulePath   string
	// ImportMap maps a package alias (path.Base of the import) to the full
	// Go import path declared in the SSaC file. Populated at newMethodGen
	// from sf.Imports so buildCallImports / buildEvalImports can look up
	// the correct import path without synthesising it (Phase006).
	ImportMap map[string]string
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
	// BodyJSONBFields lists request body property names whose OpenAPI
	// schema is the JSONB shape (`type: object, additionalProperties:
	// true`). oapi-codegen emits these as map[string]interface{} while
	// sqlc generates json.RawMessage for the matching JSONB column.
	// Populated from the request body schema at extract time so
	// sqlcArgs can wrap the source with json.Marshal(...) before
	// assigning into the params struct (BUG-005 request direction).
	BodyJSONBFields map[string]bool
	// Ownerships carries every Rego `@ownership` mapping parsed from the
	// project's policy files. Consumed by buildAuth (Phase003 ssac/purify)
	// to emit the corresponding OwnerLookup<Resource> sqlc call and
	// populate authz.CheckRequest.Owners. Nil / empty when the project has
	// no ownership annotations — buildAuth then emits an empty owners map.
	Ownerships []rego.OwnershipMapping
	// DDLTables is the parsed DDL cache shared with the convert / response
	// emitters. Plumbed into methodGen so INSERT-side helpers (sqlcArgs,
	// maybeMarshalJSONB) can resolve the target column's RawType via
	// types.MapPGType — needed to wrap literal JSONB values as []byte
	// (BUG-037 #1) and to feed the row → model rewiring of @post Model
	// = Model.Create(...) (BUG-037 #2).
	DDLTables []ddl.Table
	// SQLcQueries is the parsed sqlc query catalogue (`-- name:` entries)
	// shared with build_response so it can detect when a sqlc :one INSERT
	// returns the synthesised <Method>Row (RETURNING ... selects a subset
	// of columns) instead of the model row (RETURNING * or all columns).
	// In the Row case build_response wires through convert<Method>Row
	// which oapi-codegen emits with the right shape, instead of the
	// model-typed convert<Model> that fails the assignment (BUG-037 #2).
	SQLcQueries []sqlcparser.QuerySpec
	// activeMethod is the sqlc query name currently being emitted by
	// sqlcArgs / sqlcArgsSingle / sqlcArgsMulti. Set by the args helpers
	// before they invoke wrapJSONBLiteral / lookupSQLCMethodColumn so
	// per-column lookups can resolve the target DDL table from
	// SQLcQueries. Cleared back to empty after each emission so leakage
	// across sequences cannot misdirect a downstream literal wrap.
	activeMethod string
}
