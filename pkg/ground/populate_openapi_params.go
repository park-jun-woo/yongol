//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateOpenAPIParams — operationId별 path/query 파라미터 등록
package ground

import (

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateOpenAPIParams(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.OpenAPIDoc == nil {
		return
	}
	for _, item := range fs.OpenAPIDoc.Paths.Map() {
		populatePathOpsParams(g, item.Operations())
	}
}
