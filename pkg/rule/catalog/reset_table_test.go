//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestResetTable(t *testing.T) {
	st := &rulebookParseState{inTable: true, sawHeader: true}
	st.resetTable()
	if st.inTable || st.sawHeader {
		t.Errorf("resetTable left flags set: %+v", st)
	}
}
