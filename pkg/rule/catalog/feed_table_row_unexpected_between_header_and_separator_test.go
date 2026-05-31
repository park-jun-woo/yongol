//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestFeedTableRow_UnexpectedBetweenHeaderAndSeparator(t *testing.T) {
	st := &rulebookParseState{sawHeader: true}
	st.feedTableRow("| not a separator | x |")
	if st.sawHeader {
		t.Errorf("should bail (reset sawHeader) on unexpected row")
	}
}
