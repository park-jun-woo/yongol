//ff:func feature=generate type=util control=iteration dimension=2
//ff:what ssacUsesFileCalls — SSaC 서비스 함수의 @call file.* 사용 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// ssacUsesFileCalls mirrors pkg/validate/ssac_manifest.usesFile.
func ssacUsesFileCalls(fs *yongol.Fullstack) bool {
	if fs == nil {
		return false
	}
	for _, fn := range fs.ServiceFuncs {
		if sequencesCallPrefix(fn.Sequences, "file.") {
			return true
		}
	}
	return false
}
