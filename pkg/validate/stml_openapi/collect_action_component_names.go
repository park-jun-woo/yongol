//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectActionComponentNames — ActionBlock의 Fields(data-component:)와 children에서 컴포넌트 이름 수집

package stml_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// collectActionComponentNames gathers component names from an ActionBlock. The
// parser stores form components as FieldBind{Tag: "data-component:"+name} in
// Fields (see handle_action_component.go), so component names are extracted
// from that prefix. Children are also walked recursively.
func collectActionComponentNames(a stml.ActionBlock, out map[string]struct{}) {
	for _, f := range a.Fields {
		if name, ok := strings.CutPrefix(f.Tag, "data-component:"); ok {
			out[name] = struct{}{}
		}
	}
	for _, ch := range a.Children {
		collectChildComponentNames(ch, out)
	}
}
