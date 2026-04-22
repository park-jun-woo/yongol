//ff:func feature=gen-gogin type=util control=sequence
//ff:what durationLiteral — time.Duration을 Go 표현식으로 렌더링 (0 → "0")

package boot

import (
	"fmt"
	"time"
)

// durationLiteral renders a time.Duration as a Go expression. A zero value
// is emitted as `0` (gin-contrib/cors treats 0 as no-cache). Non-zero
// durations are emitted as `time.Duration(N)` where N is the raw nanosecond
// count — round-trip safe regardless of unit the YAML used.
func durationLiteral(d time.Duration) string {
	if d == 0 {
		return "0"
	}
	return fmt.Sprintf("time.Duration(%d)", int64(d))
}
