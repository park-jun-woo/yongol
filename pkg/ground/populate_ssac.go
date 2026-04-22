//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSSaC — SSaC에서 funcName, auth 쌍, call 참조, pub/sub 토픽 추출
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateSSaC(g *rule.Ground, fs *yongol.Fullstack) {
	funcNames := make(rule.StringSet)
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)
	subTopics := make(rule.StringSet)

	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			subTopics[fn.Subscribe.Topic] = true
		} else {
			funcNames[fn.Name] = true
		}
		// @subscribe handler 의 sequences 도 처리 — @call / @publish / @model
		// 참조가 Ground 에 등록되도록 (BUG004 fix).
		for _, seq := range fn.Sequences {
			populateSSaCSeq(g, fn.Name, seq, authPairs, callRefs, modelRefs, pubTopics)
		}
	}
	g.Lookup["SSaC.funcName"] = funcNames
	g.Pairs["SSaC.auth"] = authPairs
	g.Lookup["SSaC.callRef"] = callRefs
	g.Lookup["SSaC.modelRef"] = modelRefs
	g.Pairs["SSaC.publish"] = pubTopics
	g.Pairs["SSaC.subscribe"] = subTopics
}
