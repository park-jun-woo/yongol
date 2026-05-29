//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.mapFields 단위 테스트 (Inputs map → 정렬된 "Key: value" 문자열)

package ssac

import "testing"

func TestMethodGenMapFields(t *testing.T) {
	g := &methodGen{
		PathParams: map[string]bool{"id": true},
	}
	t.Run("empty inputs → empty string", func(t *testing.T) {
		if got := g.mapFields(map[string]string{}); got != "" {
			t.Errorf("got %q, want empty", got)
		}
	})
	t.Run("sorted keys with mapped values", func(t *testing.T) {
		inputs := map[string]string{
			"Status": `"open"`,
			"ID":     "request.id",
		}
		got := g.mapFields(inputs)
		want := `ID: request.Id, Status: "open"`
		if got != want {
			t.Errorf("mapFields = %q, want %q", got, want)
		}
	})
}
