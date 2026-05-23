//ff:func feature=agent type=helper control=selection
//ff:what resolveDescFromFile — 파일 경로에서 feature desc 해석

package agent

import "github.com/park-jun-woo/yongol/pkg/parser/features"

func resolveDescFromFile(relPath string, l layer, lookup map[string]features.Feature) string {
	switch l {
	case layerSSaC:
		op := opFromSSaCFile(relPath)
		if f, ok := lookup[op]; ok {
			return f.Desc
		}
	case layerDDL:
		return descForTable(lookup, tableFromDDLFile(relPath))
	case layerSQLcQuery:
		return descForTable(lookup, tableFromSQLcFile(relPath))
	}
	return ""
}
