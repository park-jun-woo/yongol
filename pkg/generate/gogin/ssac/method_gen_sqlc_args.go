//ff:func feature=gen-gogin type=util control=selection
//ff:what methodGen.sqlcArgs — Inputs를 sqlc 호출 인자 문자열로 변환 (0/1/N 분기, JSONB hoist)

package ssac

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
	g.activeMethod = method
	defer func() { g.activeMethod = "" }()
	switch len(inputs) {
	case 0:
		return nil, "ctx", nil
	case 1:
		return g.sqlcArgsSingle(inputs)
	default:
		return g.sqlcArgsMulti(method, inputs)
	}
}
