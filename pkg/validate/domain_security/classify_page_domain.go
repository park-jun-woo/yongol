//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what classifyPageDomain — STML 페이지가 속한 도메인 결정
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// classifyPageDomain determines which domain a page belongs to based on its
// file path and the domain frontend directory configuration.
func classifyPageDomain(page stml.PageSpec, domainDirs map[string]string) string {
	for domain, dir := range domainDirs {
		pages := filterPagesByDomain([]stml.PageSpec{page}, dir)
		if len(pages) > 0 {
			return domain
		}
	}
	return ""
}
