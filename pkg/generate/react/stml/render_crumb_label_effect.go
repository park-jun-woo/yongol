//ff:func feature=stml-gen type=generator control=sequence
//ff:what renderCrumbLabelEffect — data-crumb-field 페이지의 setCrumbLabel 배선 + fetch 도착 시 라벨·document.title 갱신 useEffect 생성

package stml

import "fmt"

// renderCrumbLabelEffect renders the dynamic crumb-label wiring of a
// data-crumb-field page (plans/stml/sitemap Phase006): useOutletContext
// receives the layout's setCrumbLabel and a useEffect feeds it the first
// fetch's response field once the data arrives, updating document.title
// with the same value (the static mount title of Phase004 stays as the
// pre-arrival fallback). The `?? {}` + optional call is the mandatory
// null guard — react-router's outlet context defaults to null, so a bare
// page (no layout, hence no provider) must not throw on destructuring or
// calling. The `v != null` guard keeps the static label until the field
// actually exists (fetch in flight, failed, or absent field).
func renderCrumbLabelEffect(field, dataVar, titleSuffix string) string {
	title := "String(v)"
	if titleSuffix != "" {
		title = "String(v) + " + tsSingleQuote(titleSuffix)
	}
	return fmt.Sprintf(`  const { setCrumbLabel } = useOutletContext<{ setCrumbLabel?: (label: string) => void }>() ?? {}
  useEffect(() => {
    const v = %s?.%s
    if (v != null) {
      setCrumbLabel?.(String(v))
      document.title = %s
    }
  }, [%s])

`, dataVar, field, title, dataVar)
}
