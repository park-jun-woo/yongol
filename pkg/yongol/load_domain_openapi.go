//ff:func feature=orchestrator type=loader control=sequence
//ff:what 한 도메인의 OpenAPI contract 를 적재하고 presence(파일 존재) 를 반환 — doc 존재 시 전역 제약 병합
package yongol

// loadDomainOpenAPI loads one domain's OpenAPI contract into fs.DomainOpenAPIDocs/
// DomainOpenAPILines and returns its presence. Presence mirrors single-site
// detection (file existence); the doc is only parsed when the file is present, so
// the returned presence never drifts from DomainOpenAPIDocs membership.
func loadDomainOpenAPI(fs *Fullstack, name, oapiPath string) SSOTPresence {
	presence := probePresence(oapiPath)
	if presence != SSOTPopulated {
		return presence
	}
	doc, lines, diags := parseOpenAPI(oapiPath)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if doc != nil {
		fs.DomainOpenAPIDocs[name] = doc
		fs.DomainOpenAPILines[name] = lines
		mergeDomainConstraints(fs, doc, lines)
	}
	return presence
}
