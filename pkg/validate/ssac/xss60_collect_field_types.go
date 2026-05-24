//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what xss60CollectFieldTypes — @publish 시퀀스의 Inputs/Fields에서 필드별 Go 타입을 추론하여 맵에 수집

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xss60CollectFieldTypes infers Go types from a publish sequence's Inputs and
// Fields, storing results in the given fields map.
func xss60CollectFieldTypes(fields map[string]string, seq parsessac.Sequence, fn parsessac.ServiceFunc, tableMap map[string]*ddl.Table) {
	for fieldName, srcExpr := range seq.Inputs {
		goType := xss60InferType(srcExpr, fn, tableMap)
		if goType != "" {
			fields[fieldName] = goType
		}
	}
	for fieldName, srcExpr := range seq.Fields {
		goType := xss60InferType(srcExpr, fn, tableMap)
		if goType != "" {
			fields[fieldName] = goType
		}
	}
}
