//ff:func feature=agent type=helper control=sequence
//ff:what formatDuration — 사람이 읽기 쉬운 시간 문자열 반환

package agent

import (
	"fmt"
	"time"
)

// formatDuration returns a human-readable duration string.
func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%dm %02ds", m, s)
}
