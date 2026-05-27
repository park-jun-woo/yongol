//ff:func feature=cli type=test control=sequence
//ff:what TestUsageArgs — usageArgs 래핑 (usageError) 검증

package main

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
)

func TestUsageArgs(t *testing.T) {
	t.Run("ValidArgs", func(t *testing.T) {
		validator := usageArgs(cobra.ExactArgs(1))
		cmd := &cobra.Command{}
		err := validator(cmd, []string{"arg1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("InvalidArgs", func(t *testing.T) {
		validator := usageArgs(cobra.ExactArgs(1))
		cmd := &cobra.Command{}
		err := validator(cmd, []string{})
		if err == nil {
			t.Fatal("expected error for wrong arg count")
		}
		var ue *usageError
		if !errors.As(err, &ue) {
			t.Errorf("expected *usageError, got %T", err)
		}
	})
}
