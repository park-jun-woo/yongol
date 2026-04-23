//ff:func feature=tsx-parser type=test control=sequence
//ff:what ParseDir — testdata/ 전체 fixture 를 파싱해 6개 이상 페이지 반환

package tsx

import "testing"

func TestParseDir(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	pages, err := ParseDir("testdata")
	if err != nil {
		t.Fatalf("ParseDir: %v", err)
	}
	if len(pages) < 6 {
		t.Errorf("want >=6 pages (fixture count), got %d", len(pages))
	}
}
