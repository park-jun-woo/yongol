//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what activePorts — ports 중 when 이 truthy 한 port 만 남긴 서브셋

package infra

import "github.com/park-jun-woo/yongol/pkg/ssacmeta"

// activePorts returns the subset of ports whose `when:` expression
// evaluates truthy under mctx. Inactive ports are silently skipped so the
// generated adapter only wires methods the manifest currently opts into.
func activePorts(ports []ssacmeta.Port, mctx map[string]any) []ssacmeta.Port {
	var out []ssacmeta.Port
	for _, p := range ports {
		if ssacmeta.EvaluateWhen(p.When, mctx) {
			out = append(out, p)
		}
	}
	return out
}
