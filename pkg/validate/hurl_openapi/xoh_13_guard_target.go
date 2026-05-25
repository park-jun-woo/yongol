//ff:func feature=validate type=util control=selection topic=hurl-openapi
//ff:what xoh13GuardTarget — guard 시퀀스 타입별 대상 설명 문자열 반환

package hurl_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

func xoh13GuardTarget(seq ssac.Sequence) string {
	switch seq.Type {
	case "empty", "exists":
		return seq.Target
	case "auth":
		return seq.Action + " " + seq.Resource
	case "state":
		return seq.DiagramID + "." + seq.Transition
	case "eval":
		if seq.Package != "" {
			return seq.Package + "." + seq.Model
		}
		return seq.Model
	default:
		return ""
	}
}
