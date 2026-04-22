//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what processResponseBody — processes a body line of the @response block
package ssac

// processResponseBody processes a single body line of the @response block.
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
