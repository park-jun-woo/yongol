//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what filterImportsUsed — body 에서 실제 참조된 import 라인만 남기기

package boot

import "testing"

func TestFilterImportsUsed(t *testing.T) {
	imports := []string{
		`"strconv"`,
		`"log/slog"`,
		`_ "github.com/lib/pq"`,
	}
	body := "n, _ := strconv.Atoi(x)\n"

	// keepBlank=false drops both the unused slog and the blank pq import.
	got := filterImportsUsed(imports, body, false)
	if !equalStrings(got, []string{`"strconv"`}) {
		t.Errorf("keepBlank=false: got %v, want [strconv]", got)
	}

	// keepBlank=true retains the side-effect pq import.
	got = filterImportsUsed(imports, body, true)
	if !equalStrings(got, []string{`"strconv"`, `_ "github.com/lib/pq"`}) {
		t.Errorf("keepBlank=true: got %v", got)
	}
}
