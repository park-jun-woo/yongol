//ff:func feature=stml-gen type=generator control=sequence
//ff:what renderTitleEffect — 페이지 마운트 시 document.title 설정 useEffect 블록 생성 (sitemap 라벨 공급원)

package stml

import "fmt"

// renderTitleEffect renders the mount useEffect that supplies
// document.title from the page's sitemap label (plans/stml/sitemap
// Phase004, DESIGN §4.6 — PageSpec has no title concept, so the sitemap
// label is the first title source). The empty dependency array pins the
// title to mount; title is the pre-joined "<label> · <app name>" string
// of GenerateOptions.DocumentTitles.
func renderTitleEffect(title string) string {
	return fmt.Sprintf("  useEffect(() => {\n    document.title = %s\n  }, [])\n", tsSingleQuote(title))
}
