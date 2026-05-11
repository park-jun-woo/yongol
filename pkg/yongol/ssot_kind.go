//ff:type feature=orchestrator type=model
//ff:what SSOT 종류를 나타내는 열거 타입
package yongol

// SSOTKind identifies one of the SSOT types in a yongol project.
type SSOTKind string

const (
	KindOpenAPI  SSOTKind = "OpenAPI"
	KindDDL      SSOTKind = "DDL"
	KindSSaC     SSOTKind = "SSaC"
	KindStates   SSOTKind = "States"
	KindPolicy   SSOTKind = "Policy"
	KindScenario SSOTKind = "Scenario"
	KindFunc     SSOTKind = "Func"
	KindConfig   SSOTKind = "Config"
	KindSTML     SSOTKind = "STML"
)
