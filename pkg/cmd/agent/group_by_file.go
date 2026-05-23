//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what groupByFile — 진단을 파일 경로별로 그룹화

package agent

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// groupByFile groups diagnostics by their File field, converting to specs-dir relative paths.
func groupByFile(diags []diagnostic.Diagnostic, absSpecs string) []fileGroup {
	m := map[string]*fileGroup{}
	var order []string
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			continue
		}
		rel := rebaseFile(d.File, absSpecs)
		if _, ok := m[rel]; !ok {
			m[rel] = &fileGroup{relFile: rel, layer: classifyFile(rel)}
			order = append(order, rel)
		}
		m[rel].diags = append(m[rel].diags, d)
	}
	result := make([]fileGroup, 0, len(order))
	for _, f := range order {
		result = append(result, *m[f])
	}
	return result
}
