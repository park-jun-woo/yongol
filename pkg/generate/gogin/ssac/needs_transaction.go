//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what needsTransaction — mutating seq 존재 판정

package ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// needsTransaction returns true when the service func has at least one
// mutating sequence (@post/@put/@delete).
func needsTransaction(sf ssac.ServiceFunc) bool {
	for _, seq := range sf.Sequences {
		switch seq.Type {
		case "post", "put", "delete":
			return true
		}
	}
	return false
}
