//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateResponseFields — @response 필드를 Ground.Schemas에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateResponseFields(g *rule.Ground, funcName string, seq ssac.Sequence) {
	if len(seq.Fields) == 0 {
		return
	}
	var fields []string
	for name := range seq.Fields {
		fields = append(fields, name)
	}
	g.Schemas["SSaC.response."+funcName] = fields
}
