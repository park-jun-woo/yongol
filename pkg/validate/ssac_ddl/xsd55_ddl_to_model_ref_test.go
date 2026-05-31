//ff:func feature=validate type=test control=iteration dimension=1
//ff:what TestXsd55DDLToModelRef — 서브테스트 디스패치
package ssac_ddl

import "testing"

func TestXsd55DDLToModelRef(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"orphan table errors", subtestTestXsd55DDLToModelRefOrphanTableErrors},
		{"func-managed exempt", subtestTestXsd55DDLToModelRefFuncManagedExempt},
		{"archived exempt", subtestTestXsd55DDLToModelRefArchivedExempt},
		{"func-managed does not exempt unrelated orphan", subtestTestXsd55DDLToModelRefFuncManagedDoesNotExemptUnrelatedOrphan},
		{"singular table matched by model", subtestTestXsd55DDLToModelRefSingularTableMatchedByModel},
		{"plural tables still matched", subtestTestXsd55DDLToModelRefPluralTablesStillMatched},
		{"genuine orphan still errors with references present", subtestTestXsd55DDLToModelRefGenuineOrphanStillErrorsWithReferencesPresent},
	} {
		t.Run(st.name, st.fn)
	}
}
