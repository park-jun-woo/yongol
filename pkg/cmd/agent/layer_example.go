//ff:func feature=agent type=helper control=selection
//ff:what layerExample — SSOT 레이어별 최소 동작 예시 반환

package agent

// layerExample returns a minimal working example for the given SSOT layer.
func layerExample(l layer) string {
	switch l {
	case layerSSaC:
		return ssacExample
	case layerDDL:
		return ddlExample
	case layerSQLcQuery:
		return sqlcExample
	case layerOpenAPI:
		return openapiExample
	case layerManifest:
		return manifestExample
	case layerRego:
		return regoExample
	case layerStateDiagram:
		return stateDiagramExample
	case layerFuncSpec:
		return funcSpecExample
	case layerHurl:
		return hurlExample
	default:
		return ""
	}
}
