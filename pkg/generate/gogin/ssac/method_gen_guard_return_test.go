//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.guardReturn 단위 테스트 (HTTP JSONResponse vs subscribe fmt.Errorf)

package ssac

import "testing"

func TestMethodGenGuardReturn(t *testing.T) {
	t.Run("subscribe returns fmt.Errorf", func(t *testing.T) {
		g := &methodGen{IsSubscribe: true}
		got := g.guardReturn("not found", 404)
		want := `return fmt.Errorf("not found")`
		if got != want {
			t.Errorf("guardReturn = %q, want %q", got, want)
		}
	})
	t.Run("http returns JSONResponse with neutral code", func(t *testing.T) {
		g := &methodGen{FuncName: "GetWorkflow"}
		got := g.guardReturn("Not found", 404)
		want := `return api.GetWorkflow404JSONResponse{Error: "Not found", Code: "not_found"}, nil`
		if got != want {
			t.Errorf("guardReturn = %q, want %q", got, want)
		}
	})
}
