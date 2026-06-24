//ff:func feature=orchestrator type=accessor control=sequence
//ff:what AllSTMLPages — 모든 도메인 STML 페이지를 DomainNames 순서로 평탄화

package yongol

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// AllSTMLPages returns every STML page across all domains. In domain mode the
// pages are flattened in DomainNames() (sorted) order for deterministic output;
// in single-site mode the singular STMLPages slice is returned as-is.
func (fs *Fullstack) AllSTMLPages() []stml.PageSpec {
	if !fs.IsDomained() {
		return fs.STMLPages
	}
	var pages []stml.PageSpec
	for _, name := range fs.DomainNames() {
		pages = append(pages, fs.DomainSTMLPages[name]...)
	}
	return pages
}
