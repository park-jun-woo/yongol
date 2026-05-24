//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-structural
//ff:what xss60CollectPublishTypes — 모든 @publish 시퀀스에서 topic → {fieldName → goType} 맵 수집

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss60CollectPublishTypes scans all @publish sequences and infers Go types for
// each payload field. Returns topic → {fieldName → goType}.
func xss60CollectPublishTypes(fs *yongol.Fullstack, tableMap map[string]*ddl.Table) map[string]map[string]string {
	out := map[string]map[string]string{}
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "publish" || seq.Topic == "" {
				continue
			}
			fields := out[seq.Topic]
			if fields == nil {
				fields = map[string]string{}
				out[seq.Topic] = fields
			}
			xss60CollectFieldTypes(fields, seq, fn, tableMap)
		}
	}
	return out
}
