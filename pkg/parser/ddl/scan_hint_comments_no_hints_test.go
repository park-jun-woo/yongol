//ff:func feature=manifest type=test control=sequence
//ff:what scanHintComments — 파일 라인별 hint 수집 (block-above / trailing / pending drain)
package ddl

import (
	"strings"
	"testing"
)

func TestScanHintComments_NoHints(t *testing.T) {
	sql := "CREATE TABLE t (id BIGINT);\n-- ordinary comment\n"
	out, err := scanHintComments(strings.NewReader(sql), "/t.sql")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(out) != 0 {
		t.Errorf("expected no hints, got %+v", out)
	}
}
