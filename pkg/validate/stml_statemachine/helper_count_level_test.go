//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what countLevel — 특정 Level + 메시지 접두어에 해당하는 진단 개수를 센다

package stml_statemachine

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func countLevel(diags []diagnostic.Diagnostic, prefix string, level diagnostic.Level) int {
	n := 0
	for _, d := range diags {
		if d.Level == level && strings.Contains(d.Message, prefix) {
			n++
		}
	}
	return n
}
