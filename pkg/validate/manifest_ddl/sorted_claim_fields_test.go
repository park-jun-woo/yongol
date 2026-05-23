//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what sortedClaimFields — nil/empty/복수 키 정렬 순서 검증

package manifest_ddl

import (
	"reflect"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSortedClaimFields(t *testing.T) {
	tests := []struct {
		name   string
		claims map[string]pmanifest.ClaimDef
		want   []string
	}{
		{
			name:   "nil map returns empty slice",
			claims: nil,
			want:   []string{},
		},
		{
			name:   "empty map returns empty slice",
			claims: map[string]pmanifest.ClaimDef{},
			want:   []string{},
		},
		{
			name: "single key",
			claims: map[string]pmanifest.ClaimDef{
				"user_id": {Key: "user_id"},
			},
			want: []string{"user_id"},
		},
		{
			name: "multiple keys sorted lexicographically",
			claims: map[string]pmanifest.ClaimDef{
				"role":    {Key: "role"},
				"user_id": {Key: "user_id"},
				"email":   {Key: "email"},
				"org_id":  {Key: "org_id"},
			},
			want: []string{"email", "org_id", "role", "user_id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedClaimFields(tt.claims)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedClaimFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
