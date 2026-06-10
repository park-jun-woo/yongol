//ff:type feature=stml-gen type=model
//ff:what 코드 생성 옵션을 설정하는 구조체
package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

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
	// NoBodyOps is the set of operationIds whose OpenAPI definition has no
	// requestBody. Void mutations use mutate() instead of mutate({}).
	NoBodyOps map[string]bool
	// PathParamTypes maps operationId → paramName → OpenAPI type (e.g.
	// "integer"). When a path parameter is "integer", the generated code
	// wraps the useParams() value with Number() to satisfy TypeScript.
	PathParamTypes map[string]map[string]string
}
