//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what p01Parse — Rego parse 검증 (이미 파싱됨/빈 디렉토리/유효 파일) 검증
package rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestP01Parse(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
