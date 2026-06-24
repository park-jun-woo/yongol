//ff:type feature=stml-gen type=model
//ff:what 코드 생성 옵션을 설정하는 구조체
package stml

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// GenerateOptions configures code generation behavior.
type GenerateOptions struct {
	APIImportPath string // import path for api module (default: "@/lib/api")
	UseClient     bool   // emit 'use client' directive (default: false)
	// BearerAuth indicates backend.auth is declared with
	// ResolvedMode() == "bearer". When true, data-capture declarations
	// commit response fields to the generated session store
	// (src/stores/auth.ts). In cookie mode (or without backend.auth)
	// captures are not emitted — only data-redirect/data-on-error apply.
	BearerAuth bool
	// RequestConstraints maps operationId → field name → FieldConstraint.
	// When non-nil, zod schemas are generated for actions that have fields.
	RequestConstraints map[string]map[string]oapiparser.FieldConstraint
	// ResponseArrayItemFields maps operationId → array field name → set of item field names.
	// Used to determine whether list items have an "id" field for React key.
	ResponseArrayItemFields map[string]map[string]map[string]bool
	// ResponseArrayItemTypes maps operationId → array field name → item field
	// name → OpenAPI type. Row actions inside data-each consult it: an
	// item.<Field> mutate argument bound to an integer path parameter is
	// wrapped with Number(...) only when the item field is not already
	// numeric in the response schema.
	ResponseArrayItemTypes map[string]map[string]map[string]string
	// ResponseBindTypes maps operationId → bind field path → FieldTypeInfo
	// (type + format). Type-aware data-bind rendering consults it: boolean →
	// Yes/No, date(-time) → locale format, number → toLocaleString, <img> →
	// src binding. A field absent from the map (or the option unwired) falls
	// back to the plain {value} emission, keeping output byte-identical
	// (plans/gen/frontend Phase037, BUG-126).
	ResponseBindTypes map[string]map[string]oapiparser.FieldTypeInfo
	// NoBodyOps is the set of operationIds whose OpenAPI definition has no
	// requestBody. Void mutations use mutate() instead of mutate({}).
	NoBodyOps map[string]bool
	// PathParamTypes maps operationId → paramName → OpenAPI type (e.g.
	// "integer"). When a path parameter is "integer", the generated code
	// wraps the useParams() value with Number() to satisfy TypeScript.
	PathParamTypes map[string]map[string]string
	// RoutePatterns maps STML page name (filename without .html) to the
	// page's resolved route pattern (stml.RoutePaths first pattern).
	// data-link emission substitutes the link's param sources into the
	// target page's pattern (page-flow Phase007).
	RoutePatterns map[string]string
	// DocumentTitles maps STML page name → the full document.title string
	// ("<sitemap label> · <app name>") emitted as a mount useEffect
	// (plans/stml/sitemap Phase004). Populated only when
	// frontend/sitemap.html exists and only for pages listed in it — a
	// page without an entry emits no title effect, so the sitemap-absent
	// output stays byte-identical.
	DocumentTitles map[string]string
	// CrumbFields maps STML page name → the sitemap data-crumb-field
	// declaration (plans/stml/sitemap Phase006). A listed page gets the
	// dynamic crumb-label wiring: useOutletContext + a useEffect that
	// feeds the first fetch's response field to the layout's setCrumbLabel
	// and updates document.title with the same value. Pages without an
	// entry stay byte-identical to the Phase004/005 emission.
	CrumbFields map[string]string
	// CrumbTitleSuffix is the " · <app name>" tail appended to the dynamic
	// crumb label when the Phase006 effect updates document.title — the
	// same join collectDocumentTitles uses for the static mount title (""
	// when the manifest carries no app name).
	CrumbTitleSuffix string
	// ErrorDisplayField is the ErrorResponse property a mutation onError
	// handler reads first when surfacing a thrown server error
	// (ExtractErrorDisplayField: "error" → "message" → default "error",
	// BUG-125/Phase036). renderOnErrorHandler normalizes "" to "error", so a
	// partially constructed GenerateOptions stays safe.
	ErrorDisplayField string
	// DesignSpec holds the parsed DESIGN.md tokens. When non-nil,
	// renderComponentJSX can consult component definitions to merge base
	// classes and validate variant/size references. Nil means no DESIGN.md
	// was declared — components emit only the STML-declared props.
	DesignSpec *design.DesignSpec
}
