//ff:func feature=gen-ir type=util control=selection
//ff:what classifyFieldArgLocation -- 단일 FieldArg 의 Source/Field 에 따라 Location 분류

package ir

// classifyFieldArgLocation sets Location on a single FieldArg based on its
// Source and the OpenAPI param maps.
func classifyFieldArgLocation(fa *FieldArg, pathParams, queryParams map[string]bool) {
	if fa.IsQuoted || (fa.Literal != "" && fa.Source == "") {
		fa.Location = LocLiteral
		return
	}
	if fa.Source == "currentUser" {
		fa.Location = LocUser
		return
	}
	if fa.Source == "request" {
		if pathParams[fa.Field] {
			fa.Location = LocPath
			return
		}
		if queryParams[fa.Field] {
			fa.Location = LocQuery
			return
		}
		fa.Location = LocBody
		return
	}
	if fa.Source != "" {
		fa.Location = LocVar
	}
}
