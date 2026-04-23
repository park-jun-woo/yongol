//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.sqlcArgs — Inputs를 sqlc 호출 인자 문자열로 변환 (0/1/N 분기, JSONB hoist)

package ssac

import (
	"sort"
	"strings"
)

// sqlcArgs converts Inputs → sqlc call arguments string. When an input's
// source value is a body JSONB field (recorded in BodyJSONBFields), the
// returned preamble emits a `<field>Raw, err := json.Marshal(<src>); if
// err != nil { return nil, err }` pair and the corresponding struct
// field references <field>Raw instead of the raw map expression.
// Needed because sqlc params expect json.RawMessage while oapi-codegen
// hands us map[string]interface{} (BUG-005 request direction).
//
// Keys are converted to sqlc PascalCase (e.g. "URL" → "Url",
// "OrgID" → "OrgID") by the caller via the validate layer.
func (g *methodGen) sqlcArgs(method string, inputs map[string]string) (preamble []string, args string, imports []string) {
	if len(inputs) == 0 {
		return nil, "ctx", nil
	}
	if len(inputs) == 1 {
		for k, v := range inputs {
			if raw, pre, needsJSON := g.maybeMarshalJSONB(k, v); needsJSON {
				preamble = append(preamble, pre...)
				imports = append(imports, `"encoding/json"`)
				return preamble, "ctx, " + raw, imports
			}
			return nil, "ctx, " + g.mapValue(v), nil
		}
	}
	// Validate (XQS-14/16) guarantees key matches sqlc field name. Use as-is.
	var fields []string
	for k, v := range inputs {
		if raw, pre, needsJSON := g.maybeMarshalJSONB(k, v); needsJSON {
			preamble = append(preamble, pre...)
			imports = append(imports, `"encoding/json"`)
			fields = append(fields, k+": "+raw)
			continue
		}
		fields = append(fields, k+": "+g.mapValue(v))
	}
	sort.Strings(fields)
	return preamble, "ctx, db." + method + "Params{" + strings.Join(fields, ", ") + "}", imports
}

// maybeMarshalJSONB returns the local raw variable name and the preamble
// lines needed to populate it, along with a flag indicating whether the
// JSONB marshal path applies. The decision is driven by the source
// expression — when it references a body property in
// BodyJSONBFields (e.g. `request.body.payload_template` or
// `request.payload_template`), the raw value is a Go map that must be
// serialised via json.Marshal before assignment to the sqlc params
// field (json.RawMessage). Non-JSONB sources fall through to mapValue.
func (g *methodGen) maybeMarshalJSONB(inputKey, srcExpr string) (rawVar string, pre []string, ok bool) {
	if len(g.BodyJSONBFields) == 0 {
		return "", nil, false
	}
	bodyProp := bodyPropertyFromExpr(srcExpr)
	if bodyProp == "" || !g.BodyJSONBFields[bodyProp] {
		return "", nil, false
	}
	local := lowerFirst(pascalCase(inputKey)) + "Raw"
	pre = []string{
		local + ", err := json.Marshal(" + g.mapValue(srcExpr) + ")",
		"if err != nil { return nil, err }",
	}
	return local, pre, true
}

// bodyPropertyFromExpr extracts the OpenAPI body property name from an
// SSaC source expression of the form `request.<prop>` or
// `request.body.<prop>`. Returns "" for other shapes so
// maybeMarshalJSONB only triggers on known body-sourced values.
func bodyPropertyFromExpr(expr string) string {
	expr = strings.TrimSpace(expr)
	const reqPrefix = "request."
	if !strings.HasPrefix(expr, reqPrefix) {
		return ""
	}
	rest := strings.TrimPrefix(expr, reqPrefix)
	rest = strings.TrimPrefix(rest, "body.")
	// Strip any further member access — we only care about the top
	// property which maps to the OpenAPI body schema field.
	if i := strings.IndexAny(rest, ". "); i >= 0 {
		rest = rest[:i]
	}
	return rest
}
