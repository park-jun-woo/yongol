//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what ddlTableSingular 단위 테스트 (복수형 → 단수형)

package ssac

import "testing"

func TestDdlTableSingular(t *testing.T) {
	cases := map[string]string{
		"users":          "user",
		"organizations":  "organization",
		"workflows":      "workflow",
		"actions":        "action",
		"execution_logs": "execution_log",
		"categories":     "category",
		"classes":        "class",
		"boxes":          "box",
		"address":        "address", // ends with ss → unchanged
		"status":         "statu",   // simple s strip (no ss)
	}
	for in, want := range cases {
		if got := ddlTableSingular(in); got != want {
			t.Errorf("ddlTableSingular(%q) = %q, want %q", in, got, want)
		}
	}
}
