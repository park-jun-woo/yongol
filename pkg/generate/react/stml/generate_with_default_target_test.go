//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what Generate와 GenerateWith(DefaultTarget) 결과 동일성 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateWithDefaultTarget(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	outDir1 := t.TempDir()
	outDir2 := t.TempDir()
	r1, err := Generate([]stmlparser.PageSpec{page}, "", outDir1)
	if err != nil {
		t.Fatal(err)
	}
	r2, err := GenerateWith(DefaultTarget(), []stmlparser.PageSpec{page}, "", outDir2)
	if err != nil {
		t.Fatal(err)
	}
	if r1.Pages != r2.Pages {
		t.Errorf("Pages mismatch: %d vs %d", r1.Pages, r2.Pages)
	}
	for k, v := range r1.Dependencies {
		if r2.Dependencies[k] != v {
			t.Errorf("Dependency %s: %q vs %q", k, v, r2.Dependencies[k])
		}
	}
}
