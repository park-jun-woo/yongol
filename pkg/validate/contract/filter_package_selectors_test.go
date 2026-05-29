//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestFilterPackageSelectors — import 패키지 셀렉터 제거, DDL 필드 후보만 남김 검증

package contract

import (
	"reflect"
	"testing"
)

func TestFilterPackageSelectors(t *testing.T) {
	t.Run("drops package selectors", func(t *testing.T) {
		in := []string{"sql.ErrNoRows", "u.Email", "json.Marshal"}
		pkgs := map[string]bool{"sql": true, "json": true}
		got := filterPackageSelectors(in, pkgs)
		want := []string{"u.Email"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("keeps non-dotted entries", func(t *testing.T) {
		in := []string{"Email", "u.Name"}
		got := filterPackageSelectors(in, map[string]bool{"u": true})
		want := []string{"Email"}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("empty pkgs returns input unchanged", func(t *testing.T) {
		in := []string{"sql.ErrNoRows"}
		got := filterPackageSelectors(in, nil)
		if !reflect.DeepEqual(got, in) {
			t.Fatalf("got %v, want %v", got, in)
		}
	})
}
