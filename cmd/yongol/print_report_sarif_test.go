//ff:func feature=cli type=test control=iteration dimension=1
//ff:what TestPrintReportSARIF — 서브테스트 디스패치
package main

import "testing"

func TestPrintReportSARIF(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"WithErrors", subtestTestPrintReportSARIFWithErrors},
		{"NoErrors", subtestTestPrintReportSARIFNoErrors},
		{"WarningsOnly", subtestTestPrintReportSARIFWarningsOnly},
		{"WriteError", subtestTestPrintReportSARIFWriteError},
		{"CatalogRulesEmbedded", subtestTestPrintReportSARIFCatalogRulesEmbedded},
	} {
		t.Run(st.name, st.fn)
	}
}
