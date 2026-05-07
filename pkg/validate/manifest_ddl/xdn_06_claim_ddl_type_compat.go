//ff:func feature=validate type=util control=iteration dimension=1 topic=manifest-infra
//ff:what claimDDLTypeCompatible — claim 타입이 DDL 컬럼 RawType 과 정합하는지 검사

package manifest_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func claimDDLTypeCompatible(claimType string, col ddl.Column) bool {
	allowed, ok := claimDDLTypeMatrix[claimType]
	if !ok {
		return false
	}
	normalized := normalizeDDLHead(col.RawType)
	for _, a := range allowed {
		if normalized == a {
			return true
		}
	}
	return false
}
