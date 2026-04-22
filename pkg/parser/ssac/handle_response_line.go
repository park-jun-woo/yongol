//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what @response 블록 내부 줄을 처리하여 블록 종료 여부 반환
package ssac

// handleResponseLine은 @response 블록 내부 줄을 처리한다.
// 블록이 종료되면 (true, 완성된 Sequence)를 반환한다.
func handleResponseLine(line string, responseLines []string, suppressWarn bool) (bool, Sequence) {
	if line != "}" {
		return false, Sequence{}
	}
	return true, Sequence{
		Type:         SeqResponse,
		Fields:       parseResponseFields(responseLines),
		SuppressWarn: suppressWarn,
	}
}
