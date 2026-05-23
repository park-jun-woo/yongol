//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what descForTable — 테이블과 일치하는 feature의 desc 반환

package agent

import "github.com/park-jun-woo/yongol/pkg/parser/features"

func descForTable(lookup map[string]features.Feature, table string) string {
	for _, f := range lookup {
		if f.Table == table {
			return f.Desc
		}
	}
	return ""
}
