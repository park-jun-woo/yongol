//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestEmitAnnotationBlockEmpty — 빈 Block은 빈 문자열 반환

package ffannot

import "testing"

func TestEmitAnnotationBlockEmpty(t *testing.T) {
	got := EmitAnnotationBlock(Block{})
	if got != "" {
		t.Fatalf("empty block should produce \"\", got %q", got)
	}
}
