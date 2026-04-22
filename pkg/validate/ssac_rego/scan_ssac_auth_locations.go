//ff:func feature=validate type=util control=iteration dimension=1 topic=policy-check
//ff:what scanSSaCAuthLocations — 단일 ServiceFunc 의 @auth 시퀀스 위치를 locs 에 기록

package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// scanSSaCAuthLocations walks a single function's sequences and records the
// first occurrence location for each (action, resource) pair into locs.
func scanSSaCAuthLocations(fn ssac.ServiceFunc, locs map[[2]string]PairLocation) {
	for _, seq := range fn.Sequences {
		if seq.Type != "auth" {
			continue
		}
		key := [2]string{seq.Action, seq.Resource}
		if _, ok := locs[key]; ok {
			continue
		}
		locs[key] = PairLocation{File: fn.FileName, Line: seq.Line}
	}
}
