//ff:func feature=orchestrator type=loader control=sequence
//ff:what 한 도메인의 STML 페이지를 적재하고 presence(glob 3-상태) 를 반환 — 디렉토리가 absent 가 아닐 때만 파싱
package yongol

// loadDomainSTML loads one domain's STML pages into fs.DomainSTMLPages and returns
// its three-state presence. Pages are only parsed when the directory is not absent
// and are stored only on a clean parse, mirroring the single-site STML loader.
func loadDomainSTML(fs *Fullstack, name, frontDir string) SSOTPresence {
	presence := probeSTMLPresence(frontDir)
	if presence == SSOTAbsent {
		return presence
	}
	pages, diags := parseSTMLDir(frontDir)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.DomainSTMLPages[name] = pages
	}
	return presence
}
