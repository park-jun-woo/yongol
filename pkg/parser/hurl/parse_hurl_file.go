//ff:func feature=crosscheck type=util control=iteration dimension=1 topic=scenario-check
//ff:what .hurl 파일에서 요청/응답 쌍 추출 (body / captures / asserts / headers 포함)
package hurl

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

var (
	reHurlRequest  = regexp.MustCompile(`^(GET|POST|PUT|DELETE|PATCH)\s+(?:\{\{host\}\}|https?://[^/]*)(/\S*)`)
	reHurlResponse = regexp.MustCompile(`^HTTP\s+(\d+)`)
)

// ParseFile extracts request/response pairs from a .hurl file. In
// addition to the request line + HTTP status, the parser collects the
// JSON request body top-level field names, request headers, the
// [Captures] block, and jsonpath [Asserts] entries. These extensions
// feed XOH-01~09 cross-check rules; pre-existing consumers that only
// read Method / Path / StatusCode are unaffected.
func ParseFile(path string) ([]HurlEntry, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot open hurl file: " + err.Error(),
		}}
	}
	defer f.Close()

	st := &parseState{path: path}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		st.lineNum++
		st.feed(scanner.Text())
	}
	st.finish()
	return st.entries, nil
}

// parseState tracks the line-by-line walk across hurl sections. Sections
// matter because the meaning of `jsonpath "$..."` differs between
// [Asserts] and [Captures].
type parseState struct {
	path     string
	lineNum  int
	current  *HurlEntry
	entries  []HurlEntry
	section  string // "", "request-headers", "body", "captures", "asserts"
	bodyBuf  strings.Builder
}

func (s *parseState) feed(raw string) {
	line := strings.TrimSpace(raw)
	if m := reHurlRequest.FindStringSubmatch(line); m != nil {
		s.flushEntry()
		s.current = &HurlEntry{
			Method: m[1],
			Path:   trimQuery(m[2]),
			File:   s.path,
			Line:   s.lineNum,
		}
		s.section = "request-headers"
		s.bodyBuf.Reset()
		return
	}
	if s.current == nil {
		return
	}
	if m := reHurlResponse.FindStringSubmatch(line); m != nil {
		s.current.StatusCode = m[1]
		s.flushRequestBody()
		s.section = "response-headers"
		return
	}
	if processSectionHeader(s, line) {
		return
	}
	processContentLine(s, raw, line)
}

// trimQuery strips a trailing `?query` fragment from a URL path.
func trimQuery(p string) string {
	if idx := strings.Index(p, "?"); idx >= 0 {
		return p[:idx]
	}
	return p
}

// flushEntry appends the current entry to entries and resets state for
// the next request. Called when a new request line appears or at EOF.
func (s *parseState) flushEntry() {
	if s.current == nil {
		return
	}
	s.flushRequestBody()
	s.entries = append(s.entries, *s.current)
	s.current = nil
	s.section = ""
	s.bodyBuf.Reset()
}

// flushRequestBody parses the accumulated request-body buffer into the
// current entry's BodyFields if the buffer looks like a JSON object.
// Called when the HTTP status line terminates the request or when the
// entry is flushed at EOF.
func (s *parseState) flushRequestBody() {
	if s.current == nil {
		return
	}
	body := strings.TrimSpace(s.bodyBuf.String())
	s.bodyBuf.Reset()
	if body == "" {
		return
	}
	fields := extractJSONFieldNames(body)
	if len(fields) > 0 {
		s.current.BodyFields = append(s.current.BodyFields, fields...)
	}
}

func (s *parseState) finish() {
	s.flushEntry()
}
