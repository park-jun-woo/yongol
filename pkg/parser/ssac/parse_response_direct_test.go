//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what @response 직접 변수 반환 파싱 검증 — Target 설정, Fields 비어있음

package ssac

import "testing"

func TestParseResponseDirect(t *testing.T) {
	src := `package service

// @get Page[Gig] gigPage = Gig.List({Query: query})
// @response gigPage
func ListGigs(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	var resp *Sequence
	for i := range sfs[0].Sequences {
		if sfs[0].Sequences[i].Type == SeqResponse {
			resp = &sfs[0].Sequences[i]
			break
		}
	}
	if resp == nil {
		t.Fatal("expected response sequence")
	}
	assertEqual(t, "Target", resp.Target, "gigPage")
	if len(resp.Fields) != 0 {
		t.Errorf("expected empty Fields for direct response, got %v", resp.Fields)
	}
}
