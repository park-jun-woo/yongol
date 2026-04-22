//ff:type feature=rule type=model
//ff:what CoverageCheckSpec — CoverageCheck 규칙의 판정 기준
package rule

// CoverageCheckSpec configures a CoverageCheck rule.
// LookupKey selects which Ground.Lookup set to check usage against.
type CoverageCheckSpec struct {
	BaseSpec
	LookupKey string
}
