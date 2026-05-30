//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteTopic — Topic 비어있지 않을 때만 topic= 덧붙임 검증

package ffannot

import (
	"strings"
	"testing"
)

func TestWriteTopic(t *testing.T) {
	t.Run("EmptyNoOp", func(t *testing.T) {
		var sb strings.Builder
		writeTopic(&sb, FuncAnnot{Topic: ""})
		if sb.String() != "" {
			t.Errorf("expected no output for empty topic, got %q", sb.String())
		}
	})

	t.Run("NonEmpty", func(t *testing.T) {
		var sb strings.Builder
		writeTopic(&sb, FuncAnnot{Topic: "auth-check"})
		if sb.String() != " topic=auth-check" {
			t.Errorf("expected ' topic=auth-check', got %q", sb.String())
		}
	})
}
