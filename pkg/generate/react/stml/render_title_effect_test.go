//ff:func feature=stml-gen type=test control=sequence
//ff:what renderTitleEffect — document.title useEffect 블록 스냅샷 + 따옴표 이스케이프 검증

package stml

import "testing"

func TestRenderTitleEffect(t *testing.T) {
	want := "  useEffect(() => {\n" +
		"    document.title = '건물 상세 · zenflow'\n" +
		"  }, [])\n"
	if got := renderTitleEffect("건물 상세 · zenflow"); got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
	if got := renderTitleEffect("it's"); got != "  useEffect(() => {\n    document.title = 'it\\'s'\n  }, [])\n" {
		t.Errorf("quote escaping broken: %s", got)
	}
}
