//ff:type feature=validate type=model topic=manifest-structural
//ff:what domainRuleCase — C-12~C-17 도메인 규칙 table-driven 공용 케이스 구조체

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

// domainRuleCase is the shared table-driven case for the domains-block rules
// (C-12~C-17): a Fullstack input and the expected diagnostic count.
type domainRuleCase struct {
	name      string
	fs        *yongol.Fullstack
	wantCount int
}
