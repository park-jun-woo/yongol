//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what resolveCallErrStatus — @call func 의 @error status 조회

package ssac

import "github.com/park-jun-woo/yongol/pkg/parser/funcspec"

// resolveCallErrStatus looks up the @error annotation status for a @call
// sequence. Priority: seq explicit > FuncSpec @error > default 500.
func resolveCallErrStatus(seqStatus int, pkgName, funcName string, project, builtin []funcspec.FuncSpec) int {
	if seqStatus != 0 {
		return seqStatus
	}
	key := pkgName + "." + lcFirst(funcName)
	for _, sp := range project {
		if sp.Package+"."+sp.Name == key && sp.ErrStatus != 0 {
			return sp.ErrStatus
		}
	}
	for _, sp := range builtin {
		if sp.Package+"."+sp.Name == key && sp.ErrStatus != 0 {
			return sp.ErrStatus
		}
	}
	return 500
}
