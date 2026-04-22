//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what @response 블록 본문 줄을 처리
package ssac

// processResponseBody는 @response 블록 본문 줄을 처리한다.
func (cp *commentParser) processResponseBody(line string) {
	done, seq := handleResponseLine(line, cp.responseLines, cp.responseSuppressWarn)
	if done {
		cp.inResponse = false
		seq.Line = cp.responseStartLine
		cp.sequences = append(cp.sequences, seq)
		cp.responseLines = nil
		cp.responseStartLine = 0
		return
	}
	cp.responseLines = append(cp.responseLines, line)
}
