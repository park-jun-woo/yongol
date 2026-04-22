//ff:func feature=ssac-parse type=parser control=sequence topic=currentuser
//ff:what @auth 권한 검사 시퀀스 파싱
package ssac

import (
	"strconv"
	"strings"
)

// parseAuth는 @auth를 파싱한다.
// "action" "resource" {inputs} "message"
func parseAuth(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)

	// "action"
	action, rest := extractQuoted(rest)
	rest = strings.TrimSpace(rest)

	// "resource"
	resource, rest := extractQuoted(rest)
	rest = strings.TrimSpace(rest)

	// {inputs}
	inputs, rest, err := extractInputs(rest)
	if err != nil {
		return nil, err
	}

	// "message" [STATUS]
	msg, remainder := extractQuoted(strings.TrimSpace(rest))

	seq := &Sequence{
		Type:     SeqAuth,
		Action:   action,
		Resource: resource,
		Inputs:   inputs,
		Message:  msg,
	}
	remainder = strings.TrimSpace(remainder)
	if remainder != "" {
		if code, err := strconv.Atoi(remainder); err == nil && code > 0 {
			seq.ErrStatus = code
		}
	}
	return seq, nil
}
