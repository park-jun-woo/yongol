//ff:func feature=tsx-parser type=command control=sequence
//ff:what 단일 .tsx 파일을 swc 로 파싱해 PageSpec 생성
package tsx

import (
	"encoding/json"
	"fmt"
	"os"
)

// Parse parses a single .tsx file and returns its PageSpec. The file is
// read once from disk to build the line index, and a second time indirectly
// by the swc wrapper (which reads the same path). The cost is negligible
// compared to parsing and keeps the runner free of STDIN plumbing.
func Parse(file string) (PageSpec, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return PageSpec{}, fmt.Errorf("read %s: %w", file, err)
	}
	astJSON, err := runSwcParse(file)
	if err != nil {
		return PageSpec{}, err
	}
	var root json.RawMessage
	if err := json.Unmarshal(astJSON, &root); err != nil {
		return PageSpec{}, fmt.Errorf("unmarshal swc ast for %s: %w", file, err)
	}
	page := PageSpec{File: file}
	v := newVisitor(src, &page)
	if err := v.walkRoot(root); err != nil {
		return PageSpec{}, fmt.Errorf("walk ast %s: %w", file, err)
	}
	return page, nil
}
