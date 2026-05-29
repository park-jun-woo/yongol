//ff:func feature=contract type=test control=sequence
//ff:what test: TestComputeBodyHashTypeOnlyFileReturnsEmpty — func 이 없는 파일은 빈 문자열 반환

package contract

import "testing"

func TestComputeBodyHashTypeOnlyFileReturnsEmpty(t *testing.T) {
	typeOnly := []byte("/" + "/ff:type feature=api type=model\n" +
		"/" + "/ff:what Foo — demo model\n" +
		"package api\n\n" +
		"type Foo struct {\n\tID int64\n}\n")
	h := ComputeBodyHash(typeOnly)
	if h != "" {
		t.Errorf("expected empty hash for type-only file, got %q", h)
	}
}
