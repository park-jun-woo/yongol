//ff:func feature=gen-gogin type=test control=sequence
//ff:what importIdentifier — import 라인에서 패키지 식별자 (마지막 경로 세그먼트) 추출

package boot

import "testing"

func TestImportIdentifier(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"stdlib", `"strconv"`, "strconv"},
		{"blank import", `_ "github.com/lib/pq"`, "pq"},
		{"explicit alias", `limiter "github.com/ulule/limiter/v3"`, "limiter"},
		{"semver major falls to parent", `"go.opentelemetry.io/otel/sdk/trace/v2"`, "trace"},
		{"plain path", `"github.com/gin-gonic/gin"`, "gin"},
		{"empty", ``, ""},
		{"no quote", `strconv`, ""},
		{"empty quotes", `""`, ""},
	}
	for _, c := range cases {
		if got := importIdentifier(c.in); got != c.want {
			t.Errorf("%s: importIdentifier(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
