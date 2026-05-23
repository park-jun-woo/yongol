//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what filterImmutable — immutable 파일(features.yaml, .hurl, .yongol)의 diagnostic 제외

package agent

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// filterImmutable removes diagnostics whose File refers to an immutable source:
//   - features.yaml
//   - tests/*.hurl or any file ending in .hurl
//   - .yongol
func filterImmutable(diags []diagnostic.Diagnostic) []diagnostic.Diagnostic {
	out := make([]diagnostic.Diagnostic, 0, len(diags))
	for _, d := range diags {
		if isImmutable(d.File) {
			continue
		}
		out = append(out, d)
	}
	return out
}
