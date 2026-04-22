//ff:func feature=contract type=test control=sequence
//ff:what test: TestComputeBodyHashIgnoresAnnotationChanges — 어노테이션 블록 변경은 hash 에 영향 없음

package contract

import "testing"

func TestComputeBodyHashIgnoresAnnotationChanges(t *testing.T) {
	body := "package demo\n\nfunc Demo(x int) int {\n\treturn x + 1\n}\n"
	without := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" + body)
	withChecked := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"/" + "/ff:checked llm=yongol-gen hash=00000000\n" + body)
	h1 := ComputeBodyHash(without)
	h2 := ComputeBodyHash(withChecked)
	if h1 != h2 {
		t.Errorf("hash should be stable across annotation-block edits: %s vs %s", h1, h2)
	}
}
