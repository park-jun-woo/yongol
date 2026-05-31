//ff:func feature=cli type=test control=iteration dimension=1
//ff:what TestValidateCmd — 서브테스트 디스패치
package main

import "testing"

func TestValidateCmd(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"ValidSpecs", subtestTestValidateCmdValidSpecs},
		{"JSONFormat", subtestTestValidateCmdJSONFormat},
		{"SARIFFormat", subtestTestValidateCmdSARIFFormat},
		{"InvalidFormat", subtestTestValidateCmdInvalidFormat},
		{"InvalidSpecsDir", subtestTestValidateCmdInvalidSpecsDir},
		{"ParseErrorJSON", subtestTestValidateCmdParseErrorJSON},
		{"ParseErrorMD", subtestTestValidateCmdParseErrorMD},
		{"WithArtsDir", subtestTestValidateCmdWithArtsDir},
	} {
		t.Run(st.name, st.fn)
	}
}
