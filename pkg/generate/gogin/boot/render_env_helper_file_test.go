//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderEnvHelperFile — env helper source 조립 + 제어흐름/import 필터 분기 검증

package boot

import (
	"strings"
	"testing"
)

func TestRenderEnvHelperFile(t *testing.T) {
	t.Run("SequenceWithUsedImport", func(t *testing.T) {
		body := "func envInt(key string, def int) int {\n" +
			"\tv := os.Getenv(key)\n" +
			"\tn, err := strconv.Atoi(v)\n" +
			"\tif err != nil {\n\t\treturn def\n\t}\n" +
			"\treturn n\n}\n"
		got := renderEnvHelperFile("envInt", body, []string{`"os"`, `"strconv"`, `"time"`})

		if !strings.Contains(got, "package main") {
			t.Errorf("missing package decl:\n%s", got)
		}
		// os and strconv are referenced -> included; time is not -> excluded.
		if !strings.Contains(got, `"os"`) || !strings.Contains(got, `"strconv"`) {
			t.Errorf("expected used imports present:\n%s", got)
		}
		if strings.Contains(got, `"time"`) {
			t.Errorf("unused import 'time' should be filtered:\n%s", got)
		}
		if !strings.Contains(got, "//ff:func") {
			t.Errorf("expected ff annotation:\n%s", got)
		}
	})

	t.Run("IterationGetsDimension", func(t *testing.T) {
		// A for-loop at depth 1 -> control=iteration -> dimension=1 annotation.
		body := "func sumEnv(keys []string) int {\n" +
			"\ttotal := 0\n" +
			"\tfor _, k := range keys {\n" +
			"\t\ttotal += len(k)\n" +
			"\t}\n" +
			"\treturn total\n}\n"
		got := renderEnvHelperFile("sumEnv", body, nil)
		if !strings.Contains(got, "control=iteration") {
			t.Errorf("expected iteration control for for-loop body:\n%s", got)
		}
		if !strings.Contains(got, "dimension=1") {
			t.Errorf("expected dimension=1 for iteration:\n%s", got)
		}
	})

	t.Run("AppendsTrailingNewline", func(t *testing.T) {
		body := "func noop() {}" // no trailing newline
		got := renderEnvHelperFile("noop", body, nil)
		if !strings.HasSuffix(got, "\n") {
			t.Errorf("expected trailing newline appended, got:\n%q", got)
		}
	})
}
