//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what matchBySchema — $ref 스키마 이름으로 feature op 매핑

package agent

import (
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

var reSchemaRef = regexp.MustCompile("#/components/schemas/(\\w+)")

func matchBySchema(msg string, offsets []pathOffset, feats []features.Feature) []string {
	refs := reSchemaRef.FindAllStringSubmatch(msg, -1)
	if len(refs) == 0 {
		return nil
	}
	var ops []string
	for _, ref := range refs {
		schema := strings.ToLower(ref[1])
		for _, feat := range feats {
			if schema == strings.ToLower(feat.Op) {
				ops = append(ops, feat.Op)
				break
			}
		}
	}
	return ops
}
