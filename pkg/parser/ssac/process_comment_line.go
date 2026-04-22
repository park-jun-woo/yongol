//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 주석 한 줄을 상태에 따라 처리
package ssac

import "strings"

// processLine은 주석 한 줄을 상태에 따라 처리한다.
// commentLine은 주석의 1-based 라인 번호 (0 = 미상).
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
