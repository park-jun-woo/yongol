//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what FetchBlock 내 EachBlock에 KeyField를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setEachKeyFieldsInFetch(f *stmlparser.FetchBlock, opID string, raif map[string]map[string]map[string]bool) {
	for i := range f.Eaches {
		setKeyFieldIfHasID(&f.Eaches[i], opID, raif)
	}
	setEachKeyFieldsInChildren(f.Children, opID, raif)
	for i := range f.NestedFetches {
		nestedOpID := f.NestedFetches[i].OperationID
		setEachKeyFieldsInFetch(&f.NestedFetches[i], nestedOpID, raif)
	}
}
