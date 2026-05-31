//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestFeedLine_NonTableResetsTable(t *testing.T) {
	st := &rulebookParseState{inTable: true, sawHeader: true}
	st.feedLine("just a paragraph")
	if st.inTable || st.sawHeader {
		t.Errorf("non-table line should reset table flags")
	}
}
