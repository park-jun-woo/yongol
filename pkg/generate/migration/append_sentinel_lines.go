//ff:func feature=migration type=util control=iteration dimension=1
//ff:what appendSentinelLines — @sentinel INSERT 블록들을 builder 에 이어붙임

package migration

import (
	"strings"
)

// appendSentinelLines writes each sentinel INSERT block into b separated
// by blank lines. Each block is trimmed of trailing newlines before being
// emitted so the snapshot stays canonical regardless of source formatting.
// Extracted from renderTable to keep that func at control=sequence.
func appendSentinelLines(b *strings.Builder, sentinels []SentinelInsert) {
	for _, s := range sentinels {
		b.WriteByte('\n')
		body := strings.TrimRight(s.SQL, "\n")
		b.WriteString(body)
		b.WriteByte('\n')
	}
}
