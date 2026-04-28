//ff:func feature=crosscheck type=util control=selection topic=scenario-check
//ff:what processSectionHeader — `[Captures]` / `[Asserts]` 등 section 헤더 인식 및 state 전이

package hurl

import (
	"regexp"
	"strings"
)

var (
	reSectionHdr   = regexp.MustCompile(`^\[([A-Za-z][A-Za-z0-9]*)\]\s*$`)
	reHeaderLine   = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9\-]*):\s*(.*)$`)
	reJSONPath     = regexp.MustCompile(`jsonpath\s+"([^"]+)"`)
	reCaptureLine  = regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)\s*:\s*(.+)$`)
	reHeaderSource = regexp.MustCompile(`^header\s+"([^"]+)"`)
)

// processSectionHeader detects `[Captures]` / `[Asserts]` / other
// section markers and switches the parser state accordingly. Returns
// true when the line was consumed by the section-header branch.
func processSectionHeader(s *parseState, line string) bool {
	m := reSectionHdr.FindStringSubmatch(line)
	if m == nil {
		return false
	}
	switch strings.ToLower(m[1]) {
	case "captures":
		s.section = "captures"
	case "asserts":
		s.section = "asserts"
	case "options", "querystringparams", "formparams", "multipartformdata", "cookies":
		s.section = "other"
	default:
		s.section = "other"
	}
	return true
}
