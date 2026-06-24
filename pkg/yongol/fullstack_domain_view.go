//ff:func feature=orchestrator type=accessor control=sequence
//ff:what DomainView — 단수 필드만 해당 도메인 데이터로 바꾼 fs 의 shallow copy 반환

package yongol

// DomainView returns a shallow copy of fs with the singular per-domain fields
// (OpenAPIDoc, OpenAPILines, STMLPages, Sitemap, Layouts) swapped to the named
// domain's data, so domain-mode codegen can reuse the existing single-site code
// paths once per domain (Decision A). RequestConstraints/ResponseConstraints and
// every other shared SSOT field stay shared by value — constraints are opID-keyed
// and globally unique under XDO-90. The receiver is not mutated.
func (fs *Fullstack) DomainView(name string) *Fullstack {
	view := *fs
	view.OpenAPIDoc = fs.DomainOpenAPIDocs[name]
	view.OpenAPILines = fs.DomainOpenAPILines[name]
	view.STMLPages = fs.DomainSTMLPages[name]
	view.Sitemap = fs.DomainSitemaps[name]
	view.Layouts = fs.DomainLayouts[name]
	return &view
}
