//ff:func feature=validate type=util control=selection
//ff:what kindPresent — 주어진 SSOT kind의 파싱 결과가 Fullstack에 존재하는지 검사
package validate

import "github.com/park-jun-woo/yongol/pkg/yongol"

func kindPresent(fs *yongol.Fullstack, k yongol.SSOTKind) bool {
	switch k {
	case yongol.KindConfig:
		return fs.Manifest != nil
	case yongol.KindOpenAPI:
		return fs.OpenAPIDoc != nil
	case yongol.KindDDL:
		return fs.DDLTables != nil
	case yongol.KindSSaC:
		return fs.ServiceFuncs != nil
	case yongol.KindStates:
		return fs.StateDiagrams != nil
	case yongol.KindPolicy:
		return fs.ParsedPolicies != nil
	case yongol.KindScenario:
		// Declared 상태 (디렉토리 있음 + 파일 0) 도 validator 가 돌아야 H-2 발행 가능
		return fs.HurlEntries != nil || fs.HurlFiles != nil || fs.PresenceOf(yongol.KindScenario) != yongol.SSOTAbsent
	case yongol.KindFunc:
		return fs.ProjectFuncSpecs != nil
	case yongol.KindTSX:
		return fs.TSXPages != nil
	}
	return false
}
