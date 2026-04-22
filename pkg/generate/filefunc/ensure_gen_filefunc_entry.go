//ff:func feature=gen-filefunc type=util control=iteration dimension=1
//ff:what ensureGenFilefuncEntry — 생성기가 반드시 필요로 하는 고정 feature 엔트리를 추가
package filefunc

// anchorFeatures are feature keys always present in the generated codebook
// because specific generator outputs hardcode them on //ff:func headers.
var anchorFeatures = map[string]string{
	"gen-filefunc":       "filefunc codebook.yaml 자동 생성",
	"runtime-middleware": "런타임 요청 검증 미들웨어 (kin-openapi)",
	"main":               "애플리케이션 엔트리포인트 (cmd/main.go)",
}

// ensureGenFilefuncEntry guarantees that the codebook contains the anchor
// feature entries so downstream validators accept generator-hardcoded
// feature names even when no on-disk package carries them.
func ensureGenFilefuncEntry(dst map[string]string) {
	for name, desc := range anchorFeatures {
		insertFeatureIfNew(dst, name, desc)
	}
}
