//ff:func feature=orchestrator type=accessor control=sequence
//ff:what DomainNames — manifest 도메인 이름을 정렬해 반환 (단일 사이트 시 빈 목록)

package yongol

import "sort"

// DomainNames returns the manifest's domain names in sorted order for
// deterministic iteration. Returns an empty (non-nil) slice for single-site
// projects (no domains declared).
func (fs *Fullstack) DomainNames() []string {
	if !fs.IsDomained() {
		return []string{}
	}
	names := make([]string, 0, len(fs.Manifest.Domains))
	for name := range fs.Manifest.Domains {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
