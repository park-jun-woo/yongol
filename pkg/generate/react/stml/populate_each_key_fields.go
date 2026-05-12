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

	// Walk top-level fetches (flat)
	for i := range page.Fetches {
		opID := page.Fetches[i].OperationID
		setEachKeyFieldsInFetch(&page.Fetches[i], opID, responseArrayItemFields)
	}

	// Walk children tree for nested fetches
	populateEachKeyFieldsInChildren(page.Children, responseArrayItemFields)
}

func setEachKeyFieldsInFetch(f *stmlparser.FetchBlock, opID string, raif map[string]map[string]map[string]bool) {
	// Set key fields on flat Eaches
	for i := range f.Eaches {
		setKeyFieldIfHasID(&f.Eaches[i], opID, raif)
	}
	// Walk children tree
	setEachKeyFieldsInChildren(f.Children, opID, raif)
	// Walk nested fetches
	for i := range f.NestedFetches {
		nestedOpID := f.NestedFetches[i].OperationID
		setEachKeyFieldsInFetch(&f.NestedFetches[i], nestedOpID, raif)
	}
}

func setEachKeyFieldsInChildren(children []stmlparser.ChildNode, opID string, raif map[string]map[string]map[string]bool) {
	for i := range children {
		ch := &children[i]
		switch ch.Kind {
		case "each":
			setKeyFieldIfHasID(ch.Each, opID, raif)
		case "fetch":
			nestedOpID := ch.Fetch.OperationID
			setEachKeyFieldsInFetch(ch.Fetch, nestedOpID, raif)
		case "static":
			if ch.Static != nil {
				setEachKeyFieldsInChildren(ch.Static.Children, opID, raif)
			}
		case "state":
			if ch.State != nil {
				setEachKeyFieldsInChildren(ch.State.Children, opID, raif)
			}
		}
	}
}

func populateEachKeyFieldsInChildren(children []stmlparser.ChildNode, raif map[string]map[string]map[string]bool) {
	for i := range children {
		ch := &children[i]
		switch ch.Kind {
		case "fetch":
			opID := ch.Fetch.OperationID
			setEachKeyFieldsInFetch(ch.Fetch, opID, raif)
		case "static":
			if ch.Static != nil {
				populateEachKeyFieldsInChildren(ch.Static.Children, raif)
			}
		case "state":
			if ch.State != nil {
				populateEachKeyFieldsInChildren(ch.State.Children, raif)
			}
		}
	}
}

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
