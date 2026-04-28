//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what collectFuncCallDiags — ServiceFunc 의 @call / @publish 시퀀스에 대한 XQS-19 진단 수집

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// collectFuncCallDiags walks f.Sequences and returns XQS-19 diagnostics
// for every @call / @publish sequence whose required interface.yaml
// ports are missing from have.
func collectFuncCallDiags(f ssacparser.ServiceFunc, interfaces map[string]*ssacmeta.PackageInterface, have map[string]bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range f.Sequences {
		diags = append(diags, collectSeqPortDiags(f, seq, interfaces, have)...)
	}
	return diags
}
