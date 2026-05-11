//ff:func feature=main type=util control=iteration dimension=1
//ff:what buildSensitiveKeys — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=8d454ea5
package main

import (
	"github.com/park-jun-woo/ssac/pkg/redact"
)

func buildSensitiveKeys(extras []string) map[string]bool {
	out := make(map[string]bool, len(redact.DefaultKeys)+len(extras))
	for k, v := range redact.DefaultKeys {
		out[k] = v
	}
	for _, k := range extras {
		out[k] = true
	}
	return out
}
