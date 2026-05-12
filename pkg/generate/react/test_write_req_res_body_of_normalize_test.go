//ff:func feature=gen-react type=test control=sequence
//ff:what BodyOf 타입에서 Record<string, never>를 {}로 정규화하는 keyof B extends never 분기 검증

package react

import (
	"strings"
	"testing"
)

func TestWriteReqResTypes_BodyOfNormalize(t *testing.T) {
	var b strings.Builder
	writeReqResTypes(&b)
	out := b.String()

	assertContains(t, out, "keyof B extends never ? {} : B")
}
