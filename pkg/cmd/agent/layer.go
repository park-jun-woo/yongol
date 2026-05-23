//ff:type feature=agent type=helper
//ff:what layer — SSOT 파일 유형 식별자 (SSaC, DDL, OpenAPI 등)

package agent

// layer identifies the SSOT type of a file path.
type layer int

const (
	layerSSaC layer = iota
	layerDDL
	layerSQLcQuery
	layerOpenAPI
	layerManifest
	layerRego
	layerStateDiagram
	layerFuncSpec
	layerHurl
	layerUnknown
)

// layerPriority defines the fix order — lower index = fix first.
var layerPriority = []layer{
	layerDDL,
	layerSQLcQuery,
	layerOpenAPI,
	layerManifest,
	layerRego,
	layerStateDiagram,
	layerFuncSpec,
	layerHurl,
	layerSSaC,
}
