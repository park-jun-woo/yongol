//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what xss60CheckSubscriber — 하나의 subscriber 함수의 필드 타입을 publish 타입과 비교

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xss60CheckSubscriber checks one subscriber function against publish field types.
func xss60CheckSubscriber(fn parsessac.ServiceFunc, publishFieldTypes map[string]map[string]string) []diagnostic.Diagnostic {
	if fn.Subscribe == nil || fn.Subscribe.Topic == "" || fn.Param == nil {
		return nil
	}
	msg := xss60FindMsgStruct(fn)
	if msg == nil {
		return nil
	}
	pubFields := publishFieldTypes[fn.Subscribe.Topic]
	if pubFields == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, field := range msg.Fields {
		if d, ok := xss60CompareFieldType(fn, field, pubFields); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
