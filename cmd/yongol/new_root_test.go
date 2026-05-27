//ff:func feature=cli type=test control=sequence
//ff:what newRoot test — root command 생성 및 subcommand 등록 검증

package main

import (
	"testing"
)

func TestNewRoot(t *testing.T) {
	t.Run("HasSubcommands", func(t *testing.T) {
		root := newRoot()
		if len(root.Commands()) < 5 {
			t.Errorf("expected at least 5 subcommands, got %d", len(root.Commands()))
		}
	})
	t.Run("UnknownFlag", func(t *testing.T) {
		_, _, err := runCmd(t, "--unknown-flag-xyz")
		if err == nil {
			t.Fatal("expected error for unknown flag, got nil")
		}
	})
	t.Run("Help", func(t *testing.T) {
		_, _, err := runCmd(t, "--help")
		if err != nil {
			t.Fatalf("help should not error: %v", err)
		}
	})
}
