//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @put 패키지 접두사 모델 파싱 검증 — cache.Cache.Set

package ssac

import "testing"

func TestParsePackagePrefixPut(t *testing.T) {
	src := `package service

// @put cache.Cache.Set({key: request.Key, value: request.Value})
func SetCache(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Package", seq.Package, "cache")
	assertEqual(t, "Model", seq.Model, "Cache.Set")
}
