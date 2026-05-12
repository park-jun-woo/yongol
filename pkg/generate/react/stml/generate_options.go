//ff:type feature=stml-gen type=model
//ff:what 코드 생성 옵션을 설정하는 구조체
package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// GenerateOptions configures code generation behavior.
type GenerateOptions struct {
	APIImportPath string // import path for api module (default: "@/lib/api")
	UseClient     bool   // emit 'use client' directive (default: false)
	// RequestConstraints maps operationId → field name → FieldConstraint.
	// When non-nil, zod schemas are generated for actions that have fields.
	RequestConstraints map[string]map[string]oapiparser.FieldConstraint
	// ResponseArrayItemFields maps operationId → array field name → set of item field names.
	// Used to determine whether list items have an "id" field for React key.
	ResponseArrayItemFields map[string]map[string]map[string]bool
}
