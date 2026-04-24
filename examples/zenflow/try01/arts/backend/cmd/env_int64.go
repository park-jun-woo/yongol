//ff:func feature=main type=util control=sequence
//ff:what envInt64 — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=c991eb4c
package main

import (
	"os"
)

func envInt64(key string, def int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := parseSize(v)
	if err != nil {
		return def
	}
	return n
}
