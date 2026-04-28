//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what collectFuncSubscribeDiags — ServiceFunc.Subscribe 에 대한 XQS-19 진단 수집

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// collectFuncSubscribeDiags returns XQS-19 diagnostics for every queue
// port whose UsedBy contains "Subscribe" and whose sqlc query is
// missing. Returns nothing when f is not a subscribe handler.
func collectFuncSubscribeDiags(f ssacparser.ServiceFunc, interfaces map[string]*ssacmeta.PackageInterface, have map[string]bool) []diagnostic.Diagnostic {
	if f.Subscribe == nil {
		return nil
	}
	iface := interfaces["queue"]
	if iface == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, port := range iface.Ports {
		if !containsUsedBy(port.UsedBy, "Subscribe") {
			continue
		}
		if have[port.Name] {
			continue
		}
		diags = append(diags, buildXqs19DiagSubscribe(f, port.Name))
	}
	return diags
}
