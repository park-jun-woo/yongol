//ff:func feature=manifest type=accessor control=sequence
//ff:what KnownRefPaths — 지원되는 manifest.* 참조 경로 목록 반환

package manifest

// KnownRefPaths returns the set of all supported manifest.* reference paths.
func KnownRefPaths() []string {
	return []string{
		"auth.accessTokenTTL",
		"auth.refreshTokenTTL",
	}
}
