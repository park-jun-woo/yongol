//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"testing"
)

func TestResolveBackendType(t *testing.T) {
	cases := []struct {
		lang, fw string
		want     BackendType
		err      bool
	}{
		{"go", "gin", GoGin, false},
		{"typescript", "nestjs", NestJS, false},
		{"python", "fastapi", FastAPI, false},
		{"ruby", "rails", "", true},
	}
	for _, c := range cases {
		assertResolveBackendType(t, c.lang, c.fw, c.want, c.err)
	}
}
