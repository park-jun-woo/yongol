//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.mapRequestValue — request.<field> 식을 Path/Query/Body 분기별로 Go 코드화

package ssac

// mapRequestValue converts a "request.<field>" SSaC expression into the
// matching oapi-codegen Go access expression. Path params route to
// `request.<Field>`, query params through queryAccessExpr, and everything
// else falls back to `request.Body.<Field>` with an optional wrapper-to-
// primitive cast (formatPrimitiveCast).
func (g *methodGen) mapRequestValue(field string) string {
	if g.PathParams[field] {
		return "request." + pascalCase(field)
	}
	if qp, isQuery := g.QueryParams[field]; isQuery {
		accessor := "request.Params." + pascalCase(field)
		return queryAccessExpr(qp, accessor)
	}
	expr := "request.Body." + pascalCase(field)
	if prim := formatPrimitiveCast(g.BodyFormats[field]); prim != "" {
		expr = prim + "(" + expr + ")"
	}
	return expr
}
