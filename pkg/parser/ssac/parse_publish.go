//ff:func feature=ssac-parse type=parser control=sequence topic=publish
//ff:what @publish 이벤트 발행 시퀀스 파싱
package ssac

import "strings"

// parsePublish는 @publish를 파싱한다.
// "topic" {payload} [{options}]
func parsePublish(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)
	topic, rest := extractQuoted(rest)
	rest = strings.TrimSpace(rest)
	payload, rest, err := extractInputs(rest)
	if err != nil {
		return nil, err
	}
	rest = strings.TrimSpace(rest)
	var options map[string]string
	if strings.HasPrefix(rest, "{") {
		options, _, err = extractInputs(rest)
		if err != nil {
			return nil, err
		}
	}
	return &Sequence{
		Type:    SeqPublish,
		Topic:   topic,
		Inputs:  payload,
		Options: options,
	}, nil
}
