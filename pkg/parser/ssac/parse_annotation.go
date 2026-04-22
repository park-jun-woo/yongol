//ff:func feature=ssac-parse type=parser control=selection
//ff:what @get/@post/@put/@delete 등 어노테이션 종류별 분기
package ssac

import "strings"

// parseAnnotation은 @response 외의 어노테이션을 종류별로 분기하여 파싱한다.
func parseAnnotation(line string) (*Sequence, error) {
	switch {
	case strings.HasPrefix(line, "@get "):
		return parseCRUD(SeqGet, line[5:], true)
	case strings.HasPrefix(line, "@post "):
		return parseCRUD(SeqPost, line[6:], true)
	case strings.HasPrefix(line, "@put "):
		return parseCRUD(SeqPut, line[5:], false)
	case strings.HasPrefix(line, "@delete "):
		return parseCRUD(SeqDelete, line[8:], false)
	case strings.HasPrefix(line, "@empty "):
		return parseGuard(SeqEmpty, line[7:]), nil
	case strings.HasPrefix(line, "@exists "):
		return parseGuard(SeqExists, line[8:]), nil
	case strings.HasPrefix(line, "@state "):
		return parseState(line[7:])
	case strings.HasPrefix(line, "@auth "):
		return parseAuth(line[6:])
	case strings.HasPrefix(line, "@publish "):
		return parsePublish(line[9:])
	case strings.HasPrefix(line, "@subscribe "):
		topic, _ := extractQuoted(strings.TrimSpace(line[11:]))
		return &Sequence{Type: "subscribe", Topic: topic}, nil
	case strings.HasPrefix(line, "@call "):
		return parseCall(line[6:])
	case strings.HasPrefix(line, "@verify-password "):
		return parseVerifyPassword(line[len("@verify-password "):])
	default:
		return nil, nil
	}
}
