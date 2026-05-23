//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-manifest
//ff:what runFileGroupCase — 파일별 entry 그룹 비교 테스트 공통 헬퍼

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func runFileGroupCase(t *testing.T, got map[string][]hurl.HurlEntry, wantFiles int, wantCounts map[string]int) {
	t.Helper()
	if len(got) != wantFiles {
		t.Fatalf("got %d files, want %d", len(got), wantFiles)
	}
	for file, wantN := range wantCounts {
		if len(got[file]) != wantN {
			t.Errorf("file %q: got %d entries, want %d", file, len(got[file]), wantN)
		}
	}
}
