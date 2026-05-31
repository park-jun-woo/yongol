//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.addParam 단위 테스트 (path → PathParams, query → QueryParams, 그 외 무시)
package ssac

func newParamGen() *methodGen {
	return &methodGen{
		PathParams:  map[string]bool{},
		QueryParams: map[string]queryParam{},
	}
}
