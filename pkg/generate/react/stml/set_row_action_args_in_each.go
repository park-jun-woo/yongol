//ff:func feature=stml-gen type=util control=sequence
//ff:what EachBlock의 item 스키마 타입을 해석해 행 액션 RowMutateArg를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setRowActionArgsInEach(e *stmlparser.EachBlock, opID string, itemTypes map[string]map[string]map[string]string, pathParamTypes map[string]map[string]string) {
	var itemFieldTypes map[string]string
	if fields, ok := itemTypes[opID]; ok {
		itemFieldTypes = fields[e.Field]
	}
	setRowActionArgsInEachChildren(e.Children, itemFieldTypes, pathParamTypes)
}
