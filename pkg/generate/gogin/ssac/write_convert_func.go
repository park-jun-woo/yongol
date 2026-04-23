//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeConvertFunc — convert<Name>(row db.X) (*api.X, error) 함수 본문 기록

package ssac

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// writeConvertFunc generates: func convertWorkflow(row db.Workflow) (*api.Workflow, error)
//
// Two layer-crossings are handled:
//
//  1. Required vs optional fields (BUG-006) — oapi-codegen emits optional
//     properties as *T and required as T. Wrapping a required T with
//     ptrOf yields *T and fails to assign, so required fields go straight
//     through as values while optional ones get ptrOf.
//
//  2. JSONB ↔ map[string]interface{} boundary (BUG-005) — when an
//     OpenAPI property is `type: object, additionalProperties: true`,
//     oapi-codegen emits a map[string]interface{} while sqlc emits
//     json.RawMessage for the JSONB column. Direct assignment is a
//     Go-level type error. writeConvertFunc declares a local variable of
//     the map type and `json.Unmarshal`s the raw bytes into it before
//     the struct literal. Any unmarshal failure propagates as the
//     convert's second return value so callers surface the HTTP 500 /
//     trigger tx rollback at the transport layer.
//
// The convert signature always returns an error — even for schemas with
// no JSONB columns — so every caller uses a single pattern. The extra
// nil return is a no-op at runtime.
func writeConvertFunc(sb *strings.Builder, name string, schema *openapi3.Schema) {
	required := requiredSet(schema)

	propNames := make([]string, 0, len(schema.Properties))
	for k := range schema.Properties {
		propNames = append(propNames, k)
	}
	sort.Strings(propNames)

	// Collect JSONB-style fields so we can emit Unmarshal scaffolding
	// before the struct literal. The per-field local variable is named
	// <lowerCamel api field>Map to match the api-side field type.
	var jsonbs []jsonbFieldAlias
	for _, jsonName := range propNames {
		if isJSONBProperty(schema.Properties[jsonName]) {
			apiField := pascalCase(jsonName)
			dbField := sqlcPascalCase(jsonName)
			jsonbs = append(jsonbs, jsonbFieldAlias{
				jsonName: jsonName,
				apiField: apiField,
				dbField:  dbField,
				localVar: lowerFirst(apiField) + "Map",
			})
		}
	}

	sb.WriteString("func convert" + name + "(row db." + name + ") (*api." + name + ", error) {\n")
	for _, j := range jsonbs {
		sb.WriteString("\tvar " + j.localVar + " map[string]interface{}\n")
		sb.WriteString("\tif len(row." + j.dbField + ") > 0 {\n")
		sb.WriteString("\t\tif err := json.Unmarshal(row." + j.dbField + ", &" + j.localVar + "); err != nil {\n")
		sb.WriteString("\t\t\treturn nil, err\n")
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n")
	}
	sb.WriteString("\treturn &api." + name + "{\n")

	for _, jsonName := range propNames {
		apiField := pascalCase(jsonName)
		dbField := sqlcPascalCase(jsonName)
		apiType := apiCastFor(name, jsonName, schema.Properties[jsonName])
		rhs := pickConvertRHS(jsonName, apiField, dbField, required[jsonName], jsonbs, apiType)
		sb.WriteString("\t\t" + apiField + ": " + rhs + ",\n")
	}

	sb.WriteString("\t}, nil\n")
	sb.WriteString("}\n")
}

// pickConvertRHS chooses the right-hand side for one struct-literal line.
// JSONB fields read from the pre-unmarshalled local map variable;
// api-typed fields (enum or format-string wrappers like
// openapi_types.Email) cast through the api-side named type since sqlc
// keeps these columns as plain string; required scalars assign the sqlc
// row value directly; optional scalars wrap with ptrOf so the *T api
// slot accepts them.
func pickConvertRHS(jsonName, apiField, dbField string, isRequired bool, jsonbs []jsonbFieldAlias, apiCast string) string {
	for _, j := range jsonbs {
		if j.jsonName == jsonName {
			return j.localVar
		}
	}
	if apiCast != "" {
		if isRequired {
			return apiCast + "(row." + dbField + ")"
		}
		return "ptrOf(" + apiCast + "(row." + dbField + "))"
	}
	if isRequired {
		return "row." + dbField
	}
	return "ptrOf(row." + dbField + ")"
}

// apiCastFor returns the Go expression (including package prefix) to
// cast a sqlc string column into the api-side named type, or "" when no
// cast is required. Two families are supported:
//
//   - String enums → `api.<Parent><Property>` (oapi-codegen's default
//     naming, e.g. WorkflowStatus for Workflow.status with enum values)
//   - String formats that oapi-codegen maps to a runtime wrapper type —
//     format: email → `openapi_types.Email` and format: uuid →
//     `openapi_types.UUID`. The runtime package is
//     github.com/oapi-codegen/runtime/types, already imported by
//     generated api files; convert files gain the import only when
//     at least one property requires it.
//
// Returns "" when the schema is missing, not a string, or has neither
// an enum nor a recognised format.
func apiCastFor(parentName, jsonName string, ref *openapi3.SchemaRef) string {
	if ref == nil || ref.Value == nil {
		return ""
	}
	s := ref.Value
	if s.Type != nil && !s.Type.Is("string") {
		return ""
	}
	if len(s.Enum) > 0 {
		return "api." + parentName + pascalCase(jsonName)
	}
	switch s.Format {
	case "email":
		return "openapi_types.Email"
	case "uuid":
		return "openapi_types.UUID"
	}
	return ""
}

// hasOpenAPITypesCast returns true when at least one of schema's string
// properties has a format that apiCastFor maps to openapi_types.*.
// Drives the conditional import of github.com/oapi-codegen/runtime/types
// in the generated convert file.
func hasOpenAPITypesCast(schema *openapi3.Schema) bool {
	if schema == nil {
		return false
	}
	for jsonName, ref := range schema.Properties {
		cast := apiCastFor("", jsonName, ref)
		if cast == "openapi_types.Email" || cast == "openapi_types.UUID" {
			return true
		}
	}
	return false
}

// jsonbFieldAlias carries the bookkeeping for one JSONB property within
// writeConvertFunc — the OpenAPI json name, its PascalCase api-side
// field, the sqlc-side PascalCase row field, and the local variable
// name that holds the unmarshalled map before the struct literal.
type jsonbFieldAlias struct {
	jsonName string
	apiField string
	dbField  string
	localVar string
}

// requiredSet flattens schema.Required into a lookup set. Returns an
// empty (non-nil) map when schema is nil or has no required list so
// callers can index without a nil check.
func requiredSet(schema *openapi3.Schema) map[string]bool {
	out := make(map[string]bool, len(schema.Required))
	if schema == nil {
		return out
	}
	for _, r := range schema.Required {
		out[r] = true
	}
	return out
}

// hasJSONBProperty returns true when any property of schema is a JSONB
// shape (see isJSONBProperty). Used by the convert-file emitter to
// decide whether to import encoding/json; schemas without JSONB fields
// skip the import to avoid the Go "imported and not used" error.
func hasJSONBProperty(schema *openapi3.Schema) bool {
	if schema == nil {
		return false
	}
	for _, p := range schema.Properties {
		if isJSONBProperty(p) {
			return true
		}
	}
	return false
}

// isJSONBProperty returns true for OpenAPI property schemas that
// oapi-codegen emits as map[string]interface{} — namely `type: object`
// with no fixed properties and additionalProperties open. The sqlc side
// stores these as json.RawMessage for a JSONB column; writeConvertFunc
// bridges them via json.Unmarshal.
func isJSONBProperty(ref *openapi3.SchemaRef) bool {
	if ref == nil || ref.Value == nil {
		return false
	}
	s := ref.Value
	if s.Type == nil || !s.Type.Is("object") {
		return false
	}
	if len(s.Properties) > 0 {
		return false
	}
	// Accept either additionalProperties: true (Has) or a permissive
	// {} schema (Schema present but no constraints). oapi-codegen maps
	// both to map[string]interface{}.
	if s.AdditionalProperties.Has != nil && *s.AdditionalProperties.Has {
		return true
	}
	if s.AdditionalProperties.Schema != nil {
		return true
	}
	return false
}

// lowerFirst lower-cases the leading rune. Used to derive a local
// variable name from an exported api field name ("PayloadTemplate" →
// "payloadTemplate"). The result is suffixed with "Map" at the call
// site so it clearly signals the decoded map.
func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
