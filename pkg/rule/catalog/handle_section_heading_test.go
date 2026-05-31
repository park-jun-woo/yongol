//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestHandleSectionHeading(t *testing.T) {
	st := &rulebookParseState{inTable: true, sawHeader: true}
	st.handleSectionHeading("## SSaC Rules")
	if st.sectionTitle != "SSaC Rules" || st.sectionSkip {
		t.Errorf("section state = %+v", st)
	}
	if st.inTable || st.sawHeader {
		t.Errorf("flags should reset on heading: %+v", st)
	}
	// Deprecated section sets skip
	st.handleSectionHeading("## Deprecated")
	if !st.sectionSkip {
		t.Errorf("Deprecated should set sectionSkip")
	}
}
