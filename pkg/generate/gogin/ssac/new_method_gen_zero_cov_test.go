//ff:func feature=gen-gogin type=test control=sequence
//ff:what zz_zerocov_extract — 0% OpenAPI 추출 헬퍼(applyOperation/extractFromOpenAPI/tryExtractFromPathItem/extractBodyFormats/extractRespFields) 검증
package ssac

func newMethodGenZeroCov(name string) *methodGen {
	return &methodGen{
		FuncName:        name,
		SuccessStatus:   200,
		PathParams:      map[string]bool{},
		QueryParams:     map[string]queryParam{},
		BodyFormats:     map[string]string{},
		RespFields:      map[string]responseField{},
		BodyJSONBFields: map[string]bool{},
		DeclaredVars:    map[string]bool{},
	}
}
