//ff:func feature=cli type=test control=iteration dimension=1
//ff:what TestGenerateCmd — 서브테스트 디스패치
package main

import "testing"

func TestGenerateCmd(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"MissingArgs", subtestTestGenerateCmdMissingArgs},
		{"OneArg", subtestTestGenerateCmdOneArg},
		{"MissingDir", subtestTestGenerateCmdMissingDir},
		{"ParseError", subtestTestGenerateCmdParseError},
		{"UnknownBackendNoManifest", subtestTestGenerateCmdUnknownBackendNoManifest},
		{"HasFlags", subtestTestGenerateCmdHasFlags},
		{"InvalidBackend", subtestTestGenerateCmdInvalidBackend},
		{"NestJSBackend", subtestTestGenerateCmdNestJSBackend},
		{"GoGinBackend", subtestTestGenerateCmdGoGinBackend},
		{"FastAPIBackend", subtestTestGenerateCmdFastAPIBackend},
		{"DefaultBackendFromManifest", subtestTestGenerateCmdDefaultBackendFromManifest},
	} {
		t.Run(st.name, st.fn)
	}
}
