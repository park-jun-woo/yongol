//ff:func feature=orchestrator type=accessor control=sequence
//ff:what IsDomained — manifest 의 domains 선언 여부로 멀티 도메인 프로젝트 판정

package yongol

// IsDomained reports whether this is a multi-domain project. True when the
// manifest declares a non-empty domains block; false for single-site projects
// (which use the singular OpenAPIDoc/STMLPages/... fields).
func (fs *Fullstack) IsDomained() bool {
	return fs.Manifest != nil && len(fs.Manifest.Domains) > 0
}
