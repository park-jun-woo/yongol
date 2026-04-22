//ff:func feature=funcspec type=parser control=sequence dimension=1
//ff:what 단일 FuncSpec의 빈 Request/Response 필드를 typeMap에서 보충
package funcspec

func fillSpecFromTypeMap(spec *FuncSpec, typeMap map[string][]Field) {
	expectedReq := ucFirst(spec.Name) + "Request"
	expectedResp := ucFirst(spec.Name) + "Response"
	if len(spec.RequestFields) == 0 {
		if fields, ok := typeMap[expectedReq]; ok {
			spec.RequestFields = fields
		}
	}
	if len(spec.ResponseFields) == 0 {
		if fields, ok := typeMap[expectedResp]; ok {
			spec.ResponseFields = fields
		}
	}
}
