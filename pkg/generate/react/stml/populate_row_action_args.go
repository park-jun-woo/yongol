//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 페이지의 data-each 내 행 액션에 mutate 호출 인자(RowMutateArg)를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// populateRowActionArgs walks all EachBlocks in the page and sets
// ActionBlock.RowMutateArg on every row action whose params reference the
// current row (item.<Field>). It mirrors populateEachKeyFields: the
// enclosing fetch operationId resolves the data-each item schema so integer
// wrapping can be decided per item field (page-flow Phase006).
func populateRowActionArgs(page *stmlparser.PageSpec, itemTypes map[string]map[string]map[string]string, pathParamTypes map[string]map[string]string) {
	for i := range page.Fetches {
		opID := page.Fetches[i].OperationID
		setRowActionArgsInFetch(&page.Fetches[i], opID, itemTypes, pathParamTypes)
	}
	setRowActionArgsInChildren(page.Children, "", itemTypes, pathParamTypes)
}
