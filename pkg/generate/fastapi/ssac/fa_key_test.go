//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what faKey — 주어진 Key 만 가진 ir.FieldArg 를 생성하는 테스트 헬퍼
package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// faKey returns a FieldArg carrying only the given Key.
func faKey(k string) ir.FieldArg { return ir.FieldArg{Key: k} }
