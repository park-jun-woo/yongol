//ff:func feature=ssac-parse type=parser control=sequence
//ff:what processLine — processes a single comment line according to current parser state
package ssac

import "strings"

// processLine processes a single comment line according to the current parser state.
// commentLine is the 1-based line number of the comment (0 = unknown).
func (cp *commentParser) processLine(line string, commentLine int) error {
	if cp.inResponse {
		cp.processResponseBody(line)
		return nil
	}

	if !strings.HasPrefix(line, "@") {
		return nil
	}

	seq, isResponseStart, err := parseLine(line)
	if err != nil {
		return err
	}
	if isResponseStart {
		cp.inResponse = true
		cp.responseSuppressWarn = strings.HasPrefix(line, "@response!")
		cp.responseLines = nil
		cp.responseStartLine = commentLine
		return nil
	}
	if seq != nil {
		seq.Line = commentLine
		cp.sequences = append(cp.sequences, *seq)
	}
	return nil
}
