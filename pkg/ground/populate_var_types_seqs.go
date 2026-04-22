//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateVarTypesSeqs — 시퀀스에서 변수→원본 타입 스펙 매핑 수집
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateVarTypesSeqs registers variable→original type string mappings.
// Original string preserves wrapper/slice/pointer/package-prefix (e.g.,
// "[]Webhook", "*User", "billing.CheckCreditsResponse"). Downstream consumers
// (XOS-67, etc.) strip prefixes as needed via inferResponseValueType.
func populateVarTypesSeqs(g *rule.Ground, funcName string, seqs []ssac.Sequence) {
	for _, seq := range seqs {
		if seq.Result == nil || seq.Result.Var == "" || seq.Result.Type == "" {
			continue
		}
		g.Types["SSaC.var."+funcName+"."+seq.Result.Var] = seq.Result.Type
	}
}
