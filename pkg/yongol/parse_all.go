//ff:func feature=orchestrator type=command control=sequence
//ff:what 탐지된 모든 SSOT를 1회 파싱하여 Fullstack에 담아 반환. 수집형 fail.
package yongol

// ParseAll parses all detected SSOTs once and returns the results.
// Each SSOT is parsed independently; parser diagnostics from any SSOT are
// collected into fs.ParseDiagnostics (collect-and-continue fail mode). The CLI-level orchestrator
// gates validate execution on ParseDiagnostics being empty.
func ParseAll(root string, detected []DetectedSSOT) *Fullstack {
	fs := &Fullstack{SpecsDir: root, Presences: make(map[SSOTKind]SSOTPresence)}

	has := indexDetected(detected, fs)

	parseManifestIfPresent(fs, root, has)
	parseOpenAPIIfPresent(fs, has)
	parseSSaCIfPresent(fs, has)
	parseStatesIfPresent(fs, has)
	parseDDLIfPresent(fs, has)
	parsePolicyIfPresent(fs, has)
	parseFuncIfPresent(fs, has)
	parseScenarioIfPresent(fs, has)
	parseTSXIfPresent(fs, has)
	parseYongolPkgSpecs(fs)

	return fs
}
