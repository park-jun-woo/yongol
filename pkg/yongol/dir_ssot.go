//ff:type feature=orchestrator type=model
//ff:what dirSSOT — detect_ssots 전용 directory-backed SSOT 설명자
package yongol

// dirSSOT describes a directory-backed SSOT candidate: the kind, its expected
// directory under the specs root, and the glob patterns that, when matched,
// count as "populated" content.
type dirSSOT struct {
	kind  SSOTKind
	dir   string
	globs []string
}
