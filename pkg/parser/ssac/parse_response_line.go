//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what parseResponseLine — parses an @response line, returning a Sequence or signalling multi-line start
package ssac

import "strings"

// parseResponseLine parses an @response line.
func parseResponseLine(line string) (*Sequence, bool, error) {
	tag := "@response"
	suppressWarn := false
	if strings.HasPrefix(line, "@response!") {
		tag = "@response!"
		suppressWarn = true
	}
	trimmed := strings.TrimSpace(strings.TrimPrefix(line, tag))
	if trimmed == "{" {
		return nil, true, nil
	}
	// single-line struct: @response { field: var, ... }
	if strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}") {
		inner := trimmed[1 : len(trimmed)-1]
		lines := strings.Split(inner, ",")
		return &Sequence{
			Type:         SeqResponse,
			Fields:       parseResponseFields(lines),
			SuppressWarn: suppressWarn,
		}, false, nil
	}
	// shorthand: @response varName
	if trimmed != "" {
		return &Sequence{
			Type:         SeqResponse,
			Target:       trimmed,
			SuppressWarn: suppressWarn,
		}, false, nil
	}
	// @response 또는 @response! — 빈 응답 (204 등)
	return &Sequence{Type: SeqResponse, SuppressWarn: suppressWarn}, false, nil
}
