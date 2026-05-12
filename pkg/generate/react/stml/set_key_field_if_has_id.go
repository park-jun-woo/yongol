//ff:func feature=stml-gen type=util control=sequence
//ff:what EachBlock의 응답 스키마에 id 필드가 있으면 KeyField를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setKeyFieldIfHasID(e *stmlparser.EachBlock, opID string, raif map[string]map[string]map[string]bool) {
	fields, ok := raif[opID]
	if !ok {
		return
	}
	itemFields, ok := fields[e.Field]
	if !ok {
		return
	}
	if itemFields["id"] {
		e.KeyField = "id"
	}
}
