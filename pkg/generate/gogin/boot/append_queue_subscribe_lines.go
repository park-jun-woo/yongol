//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what appendQueueSubscribeLines — ServiceFunc @subscribe에서 queue.Subscribe 라인 추가

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// appendQueueSubscribeLines appends a queue.Subscribe call for each
// ServiceFunc with an @subscribe annotation, preserving declaration
// order.
func appendQueueSubscribeLines(lines []string, serviceFuncs []ssac.ServiceFunc) []string {
	for _, fn := range serviceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		lines = append(lines,
			fmt.Sprintf(`queue.Subscribe(%q, srv.%s)`, fn.Subscribe.Topic, fn.Name),
		)
	}
	return lines
}
