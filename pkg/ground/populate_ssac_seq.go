//ff:func feature=rule type=loader control=selection
//ff:what populateSSaCSeq — 개별 시퀀스에서 auth/call/model/publish 정보 추출
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateSSaCSeq(g *rule.Ground, funcName string, seq ssac.Sequence,
	authPairs, callRefs, modelRefs, pubTopics rule.StringSet) {

	switch seq.Type {
	case "auth":
		authPairs[seq.Action+":"+seq.Resource] = true
	case "call":
		registerSSaCCallRef(seq.Model, callRefs)
	case "publish":
		pubTopics[seq.Topic] = true
	case "get", "post", "put", "delete":
		registerSSaCModelRef(seq.Model, modelRefs)
	case "response":
		populateResponseFields(g, funcName, seq)
	}
}
