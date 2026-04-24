//ff:func feature=main type=util control=sequence
//ff:what envBool — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=5b129301
package main

import (
	"os"
	"strings"
)

func envBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		return strings.EqualFold(v, "true") || v == "1"
	}
	return def
}
