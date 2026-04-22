//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what filterImportsUsed — body 에서 실제 참조된 import 라인만 남기기

package boot

// filterImportsUsed keeps only import lines whose package identifier (the
// last path segment) actually appears in the body. The matcher is a crude
// substring check which is sufficient because helper / main bodies stay
// small. keepBlank controls whether side-effect (`_ "…"`) imports survive
// filtering — main.go needs them (e.g. pq driver), helper files don't.
func filterImportsUsed(imports []string, body string, keepBlank bool) []string {
	var keep []string
	for _, imp := range imports {
		if shouldKeepImport(imp, body, keepBlank) {
			keep = append(keep, imp)
		}
	}
	return keep
}
