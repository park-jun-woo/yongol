//ff:func feature=validate type=test control=sequence
//ff:what TestWithArtsDir — artsDir 옵션이 config에 주입되는지 검증

package validate

import "testing"

func TestWithArtsDir(t *testing.T) {
	c := &config{}
	WithArtsDir("/tmp/arts")(c)
	if c.artsDir != "/tmp/arts" {
		t.Fatalf("artsDir = %q, want /tmp/arts", c.artsDir)
	}
}
