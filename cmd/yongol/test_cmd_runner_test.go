//ff:func feature=cli type=test-helper control=sequence
//ff:what runCmd — cobra CLI integration test 실행 헬퍼 (stdout/stderr/err 캡처)

package main

import (
	"bytes"
	"testing"
)

// runCmd executes the yongol cobra tree built by newRoot() with the given
// argv and returns captured stdout, stderr, and the Execute() error. Tests
// assert on the error to infer the intended exit code:
//
//   - err == nil                  → exit 0
//   - err is *usageError          → exit 2 (mirrors main.go errors.As branch)
//   - err != nil (other)          → exit 1
//
// Nothing in the cobra tree writes to os.Stdout/os.Stderr when SetOut/SetErr
// are wired — except versionCmd which uses fmt.Printf directly. The version
// test therefore captures via the os.Stdout redirect helper below.
func runCmd(t *testing.T, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	root := newRoot()
	var sb, eb bytes.Buffer
	root.SetOut(&sb)
	root.SetErr(&eb)
	root.SetArgs(args)
	err = root.Execute()
	return sb.String(), eb.String(), err
}
