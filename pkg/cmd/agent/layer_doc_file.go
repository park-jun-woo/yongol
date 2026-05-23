//ff:func feature=agent type=helper control=selection
//ff:what layerDocFile — 레이어별 docs 파일명 매핑

package agent

// layerDocFile maps a layer to its docs filename.
func layerDocFile(l layer) string {
	switch l {
	case layerSSaC:
		return "ssac.md"
	case layerDDL:
		return "ddl.md"
	case layerSQLcQuery:
		return "sqlc.md"
	case layerOpenAPI:
		return "openapi.md"
	case layerRego:
		return "policy.md"
	case layerStateDiagram:
		return "states.md"
	case layerHurl:
		return "scenario.md"
	case layerManifest:
		return "manifest.md"
	case layerFuncSpec:
		return "func.md"
	default:
		return ""
	}
}
