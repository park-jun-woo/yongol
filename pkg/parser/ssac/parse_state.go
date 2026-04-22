//ff:func feature=ssac-parse type=parser control=sequence topic=states
//ff:what @state 상태 전이 시퀀스 파싱
package ssac

import (
	"strconv"
	"strings"
)

// parseState는 @state를 파싱한다.
// diagramID {inputs} "transition" "message"
func parseState(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)

	// diagramID
	spaceIdx := strings.IndexByte(rest, ' ')
	if spaceIdx < 0 {
		return nil, nil
	}
	diagramID := rest[:spaceIdx]
	rest = strings.TrimSpace(rest[spaceIdx+1:])

	// {inputs}
	inputs, rest, err := extractInputs(rest)
	if err != nil {
		return nil, err
	}

	// "transition" "message" [STATUS]
	transition, msg, remainder := parseTwoQuoted(rest)

	seq := &Sequence{
		Type:       SeqState,
		DiagramID:  diagramID,
		Inputs:     inputs,
		Transition: transition,
		Message:    msg,
	}
	if remainder != "" {
		if code, err := strconv.Atoi(remainder); err == nil && code > 0 {
			seq.ErrStatus = code
		}
	}
	return seq, nil
}
