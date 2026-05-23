//ff:func feature=agent type=helper control=selection
//ff:what layerName — 레이어의 사람이 읽기 쉬운 이름 반환

package agent

// layerName returns a human-readable name for a layer.
func layerName(l layer) string {
	switch l {
	case layerSSaC:
		return "SSaC"
	case layerDDL:
		return "DDL"
	case layerSQLcQuery:
		return "sqlc query"
	case layerOpenAPI:
		return "OpenAPI"
	case layerManifest:
		return "manifest"
	case layerRego:
		return "Rego"
	case layerStateDiagram:
		return "stateDiagram"
	case layerFuncSpec:
		return "func spec"
	case layerHurl:
		return "Hurl"
	default:
		return "unknown"
	}
}
