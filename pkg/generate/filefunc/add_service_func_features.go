//ff:func feature=gen-filefunc type=util control=iteration dimension=1
//ff:what addServiceFuncFeatures — SSaC ServiceFunc 의 Feature 폴더명을 맵에 추가
package filefunc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// addServiceFuncFeatures inserts each SSaC ServiceFunc.Feature into dst as
// a key with empty description. Empty feature names are skipped.
func addServiceFuncFeatures(dst map[string]string, fs *yongol.Fullstack) {
	for i := range fs.ServiceFuncs {
		feat := strings.TrimSpace(fs.ServiceFuncs[i].Feature)
		insertFeatureIfNew(dst, feat, "")
	}
}
