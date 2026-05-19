//ff:func feature=validate type=util control=selection
//ff:what kindPresent — checks whether the parse result for a given SSOT kind is present in Fullstack
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
		// A Declared state (directory present + 0 files) must still run the validator so H-2 can fire
		return fs.HurlEntries != nil || fs.HurlFiles != nil || fs.PresenceOf(yongol.KindScenario) != yongol.SSOTAbsent
	case yongol.KindFunc:
		return fs.ProjectFuncSpecs != nil
	case yongol.KindSTML:
		return fs.STMLPages != nil
	case yongol.KindDesign:
		return fs.DesignSpec != nil
	case yongol.KindFeatures:
		return fs.Features != nil
	}
	return false
}
