//ff:func feature=generate type=util control=iteration dimension=2
//ff:what ssacUsesSessionCalls — SSaC 서비스 함수의 @call session.* 사용 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// ssacUsesSessionCalls mirrors pkg/validate/ssac_manifest.usesSession
// without taking a cross-package dependency. Keeping the detection
// local lets generate stay above validate in the import graph.
func ssacUsesSessionCalls(fs *yongol.Fullstack) bool {
	if fs == nil {
		return false
	}
	for _, fn := range fs.ServiceFuncs {
		if sequencesCallPrefix(fn.Sequences, "session.") {
			return true
		}
	}
	return false
}
