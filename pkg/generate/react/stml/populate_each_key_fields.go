//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 페이지의 EachBlock에 OpenAPI 응답 스키마 기반 KeyField를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// populateEachKeyFields walks all EachBlocks in the page and sets KeyField
// to "id" when the OpenAPI response schema indicates the array items have
// an "id" field. When responseArrayItemFields is nil, no key fields are set.
func populateEachKeyFields(page *stmlparser.PageSpec, responseArrayItemFields map[string]map[string]map[string]bool) {
	if responseArrayItemFields == nil {
		return
	}

	for i := range page.Fetches {
		opID := page.Fetches[i].OperationID
		setEachKeyFieldsInFetch(&page.Fetches[i], opID, responseArrayItemFields)
	}

	populateEachKeyFieldsInChildren(page.Children, responseArrayItemFields)
}
