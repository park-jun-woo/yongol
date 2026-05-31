//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFileOp — 추출오류/LLM오류/빈응답/머지오류/성공(OpenAPI·Rego·Hurl)+desc lookup·msg fallback 분기 검증
package agent

func mockLLM(reply string, err error) LLMCallFunc {
	return func(b, m, s, u string) (string, error) { return reply, err }
}
