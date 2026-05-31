//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderModule_EmptyFeature
package ssac

import (
	"testing"
)

func TestRenderModule_EmptyFeature(t *testing.T) {
	_, err := RenderModule("", nil)
	if err == nil {
		t.Error("empty feature should return error")
	}
}
