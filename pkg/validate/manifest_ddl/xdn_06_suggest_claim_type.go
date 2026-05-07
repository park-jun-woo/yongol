//ff:func feature=validate type=util control=iteration dimension=2 topic=manifest-infra
//ff:what suggestClaimType — DDL 컬럼 RawType 에서 적합한 claim 타입 추천

package manifest_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func suggestClaimType(col ddl.Column) string {
	normalized := normalizeDDLHead(col.RawType)
	for claimType, ddlTypes := range claimDDLTypeMatrix {
		for _, dt := range ddlTypes {
			if normalized == dt {
				return claimType
			}
		}
	}
	return "string"
}
