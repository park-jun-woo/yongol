//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what s71CollectRefs — 시퀀스에서 변수 참조 후보 값 수집 (Inputs, Target, EmailExpr, PasswordExpr, Fields)

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func s71CollectRefs(seq parsessac.Sequence) []string {
	var refs []string
	for _, v := range seq.Inputs {
		refs = append(refs, v)
	}
	if seq.Target != "" {
		refs = append(refs, seq.Target)
	}
	if seq.EmailExpr != "" {
		refs = append(refs, seq.EmailExpr)
	}
	if seq.PasswordExpr != "" {
		refs = append(refs, seq.PasswordExpr)
	}
	for _, v := range seq.Fields {
		refs = append(refs, v)
	}
	return refs
}
