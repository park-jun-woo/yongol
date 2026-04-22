//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what handleResponseLine — processes a line inside a @response block and reports whether the block has ended
package ssac

// handleResponseLine processes a line inside a @response block.
// It returns (true, completed Sequence) when the block ends.
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
