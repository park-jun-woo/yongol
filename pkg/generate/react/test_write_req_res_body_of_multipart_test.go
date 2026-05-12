//ff:func feature=gen-react type=test control=sequence
//ff:what BodyOf 타입에 multipart/form-data fallback이 포함되는지 검증

package react

import (
	"strings"
	"testing"
)

func TestWriteReqResTypes_BodyOfMultipartFallback(t *testing.T) {
	var b strings.Builder
	writeReqResTypes(&b)
	out := b.String()

	assertContains(t, out, "'multipart/form-data': infer B")
}
