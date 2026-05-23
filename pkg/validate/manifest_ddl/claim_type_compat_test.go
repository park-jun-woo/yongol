//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what claimTypeCompatible — 빈 문자열 기본값 + 동일/비동일 타입 검증

package manifest_ddl

import "testing"

func TestClaimTypeCompatible(t *testing.T) {
	tests := []struct {
		name        string
		claimGoType string
		ddlGoType   string
		want        bool
	}{
		{
			name:        "empty claim defaults to string, matches string",
			claimGoType: "",
			ddlGoType:   "string",
			want:        true,
		},
		{
			name:        "empty claim defaults to string, mismatch with int64",
			claimGoType: "",
			ddlGoType:   "int64",
			want:        false,
		},
		{
			name:        "int64 matches int64",
			claimGoType: "int64",
			ddlGoType:   "int64",
			want:        true,
		},
		{
			name:        "string matches string",
			claimGoType: "string",
			ddlGoType:   "string",
			want:        true,
		},
		{
			name:        "bool matches bool",
			claimGoType: "bool",
			ddlGoType:   "bool",
			want:        true,
		},
		{
			name:        "int64 vs string mismatch",
			claimGoType: "int64",
			ddlGoType:   "string",
			want:        false,
		},
		{
			name:        "string vs bool mismatch",
			claimGoType: "string",
			ddlGoType:   "bool",
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := claimTypeCompatible(tt.claimGoType, tt.ddlGoType)
			if got != tt.want {
				t.Errorf("claimTypeCompatible(%q, %q) = %v, want %v",
					tt.claimGoType, tt.ddlGoType, got, tt.want)
			}
		})
	}
}
