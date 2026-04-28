//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what collectSeqPortDiags — 한 시퀀스에 대해 interface.yaml ports 요구 검증

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// collectSeqPortDiags resolves the (pkg, method) pair for seq and
// returns XQS-19 diagnostics for each interface.yaml port whose sqlc
// query is missing. Non-DB-facing sequences contribute nothing.
func collectSeqPortDiags(f ssacparser.ServiceFunc, seq ssacparser.Sequence, interfaces map[string]*ssacmeta.PackageInterface, have map[string]bool) []diagnostic.Diagnostic {
	pkg, method := resolveBuiltinCall(seq, f.Subscribe != nil)
	if pkg == "" {
		return nil
	}
	iface := interfaces[pkg]
	if iface == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, port := range iface.Ports {
		if !containsUsedBy(port.UsedBy, method) {
			continue
		}
		if have[port.Name] {
			continue
		}
		diags = append(diags, buildXqs19Diag(f, seq, pkg, method, port.Name))
	}
	return diags
}
