//ff:func feature=main type=util control=iteration dimension=1
//ff:what parseSize — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=a7f1c56b
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseSize(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty size")
	}
	up := strings.ToUpper(s)
	var mult int64 = 1
	for _, suf := range []struct {
		tag string
		v   int64
	}{
		{"KIB", 1 << 10}, {"MIB", 1 << 20}, {"GIB", 1 << 30}, {"TIB", 1 << 40},
		{"KB", 1000}, {"MB", 1000 * 1000}, {"GB", 1000 * 1000 * 1000}, {"TB", 1000 * 1000 * 1000 * 1000},
		{"B", 1},
	} {
		if strings.HasSuffix(up, suf.tag) {
			up = strings.TrimSuffix(up, suf.tag)
			mult = suf.v
			break
		}
	}
	up = strings.TrimSpace(up)
	n, err := strconv.ParseInt(up, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse size %q: %w", s, err)
	}
	if n < 0 {
		return 0, fmt.Errorf("negative size %q", s)
	}
	return n * mult, nil
}
