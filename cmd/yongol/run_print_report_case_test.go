//ff:func feature=cli type=test control=iteration dimension=1
//ff:what TestRunPrintReportCase — 서브테스트 디스패치
package main

import "testing"

func TestRunPrintReportCase(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"AllPass", subtestTestRunPrintReportCaseAllPass},
		{"OneError", subtestTestRunPrintReportCaseOneError},
		{"OneWarning", subtestTestRunPrintReportCaseOneWarning},
		{"ErrorWithMsgCheck", subtestTestRunPrintReportCaseErrorWithMsgCheck},
		{"MultipleSteps", subtestTestRunPrintReportCaseMultipleSteps},
		{"ErrorNoMsgSubstring", subtestTestRunPrintReportCaseErrorNoMsgSubstring},
	} {
		t.Run(st.name, st.fn)
	}
}
