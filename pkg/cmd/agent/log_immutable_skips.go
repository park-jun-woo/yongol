//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what logImmutableSkips — immutable 파일 건너뛰기 요약 출력

package agent

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// logImmutableSkips prints a single summary line per immutable file.
func logImmutableSkips(w io.Writer, diags []diagnostic.Diagnostic, absSpecs string) {
	seen := map[string]bool{}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			continue
		}
		if !isImmutable(d.File) {
			continue
		}
		rel := rebaseFile(d.File, absSpecs)
		if seen[rel] {
			continue
		}
		seen[rel] = true
		fmt.Fprintf(w, "  skipped: %s (immutable)\n", rel)
	}
}
