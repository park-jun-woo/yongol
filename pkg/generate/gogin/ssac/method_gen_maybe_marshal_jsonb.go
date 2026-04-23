//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.maybeMarshalJSONB — body JSONB 필드일 때 json.Marshal 프리앰블과 raw 변수명을 반환
package ssac

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
