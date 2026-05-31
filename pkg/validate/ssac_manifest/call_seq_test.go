//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func callSeq(model string) ssac.Sequence {
	return ssac.Sequence{Type: "call", Model: model}
}
