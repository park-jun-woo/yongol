//ff:func feature=validate type=util control=iteration dimension=1 topic=config-check
//ff:what ServiceFunc 가 tx-bound @publish 조합을 포함하는지 판정

package ssac_manifest

import ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// hasTxBoundPublish returns true when the service func has both a mutating
// sequence (which triggers tx generation in yongol codegen) and a @publish
// sequence. This mirrors pkg/generate/gogin/ssac/needs_transaction.go.
func hasTxBoundPublish(fn ssacparser.ServiceFunc) bool {
	var hasMutation, hasPublish bool
	for _, seq := range fn.Sequences {
		switch seq.Type {
		case "post", "put", "delete":
			hasMutation = true
		case "publish":
			hasPublish = true
		}
	}
	return hasMutation && hasPublish
}
