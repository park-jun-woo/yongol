//ff:func feature=main type=util control=sequence
//ff:what envString — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=95dba27d
package main

import (
	"os"
)

func envString(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
