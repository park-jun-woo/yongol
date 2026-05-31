//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHasDeferCloseAfter — defer name.Close() 존재/부재 + 각 스킵 분기 검증
package qcheck

import (
	"testing"
)

func TestHasDeferCloseAfter(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		list := blockStmts(t, "f := open()\ndefer f.Close()")
		if !hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected defer f.Close() to be found")
		}
	})

	t.Run("WrongName", func(t *testing.T) {
		list := blockStmts(t, "f := open()\ndefer g.Close()")
		if hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected false for mismatched receiver name")
		}
	})

	t.Run("NotDeferStmt", func(t *testing.T) {
		list := blockStmts(t, "f := open()\nf.Close()")
		if hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected false: Close() not deferred")
		}
	})

	t.Run("WrongMethod", func(t *testing.T) {
		list := blockStmts(t, "f := open()\ndefer f.Flush()")
		if hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected false for non-Close method")
		}
	})

	t.Run("NonSelectorDefer", func(t *testing.T) {
		list := blockStmts(t, "f := open()\ndefer cleanup()")
		if hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected false for non-selector defer")
		}
	})

	t.Run("NonIdentReceiver", func(t *testing.T) {
		list := blockStmts(t, "f := open()\ndefer obj.inner.Close()")
		if hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected false for non-ident receiver")
		}
	})

	t.Run("NoneAfter", func(t *testing.T) {
		list := blockStmts(t, "f := open()")
		if hasDeferCloseAfter(list, 0, "f") {
			t.Errorf("expected false when nothing follows")
		}
	})
}
