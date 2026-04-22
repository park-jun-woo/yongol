//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 한 줄의 @annotation을 파싱하여 Sequence 반환
package ssac

import "strings"

// parseLine은 한 줄을 파싱하여 Sequence를 반환한다.
// @response { 의 경우 (nil, true, nil)를 반환하여 멀티라인 모드 시작을 알린다.
func parseLine(line string) (*Sequence, bool, error) {
	if strings.HasPrefix(line, "@response") {
		return parseResponseLine(line)
	}

	// @type! — ! 접미사 감지
	suppressWarn := false
	if idx := strings.IndexByte(line, '!'); idx > 0 {
		spaceIdx := strings.IndexByte(line, ' ')
		if spaceIdx < 0 || idx < spaceIdx {
			line = line[:idx] + line[idx+1:]
			suppressWarn = true
		}
	}

	seq, err := parseAnnotation(line)
	if err != nil {
		return nil, false, err
	}
	if seq != nil && suppressWarn {
		seq.SuppressWarn = true
	}
	return seq, false, nil
}
