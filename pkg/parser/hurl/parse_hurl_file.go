//ff:func feature=crosscheck type=util control=iteration dimension=1 topic=scenario-check
//ff:what ParseFile — .hurl 파일에서 요청/응답 쌍 추출 (body / captures / asserts / headers 포함)

package hurl

import (
	"bufio"
	"os"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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
