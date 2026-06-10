//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what FetchBlock 내 EachBlock의 행 액션에 RowMutateArg를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func setRowActionArgsInFetch(f *stmlparser.FetchBlock, opID string, itemTypes map[string]map[string]map[string]string, pathParamTypes map[string]map[string]string) {
	for i := range f.Eaches {
		setRowActionArgsInEach(&f.Eaches[i], opID, itemTypes, pathParamTypes)
	}
	setRowActionArgsInChildren(f.Children, opID, itemTypes, pathParamTypes)
	for i := range f.NestedFetches {
		nestedOpID := f.NestedFetches[i].OperationID
		setRowActionArgsInFetch(&f.NestedFetches[i], nestedOpID, itemTypes, pathParamTypes)
	}
}
