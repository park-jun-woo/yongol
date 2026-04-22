//ff:func feature=gen-gogin type=util control=sequence
//ff:what WriteIfNotPreserved — 대상 파일이 preserve 상태이면 skip, 아니면 os.WriteFile 로 기록

package fffile

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

// WriteIfNotPreserved writes content to path with 0644 permissions unless
// the existing file at path is in StatePreserved. A preserved file carries
// a `//ff:checked llm=yongol-gen hash=<saved>` whose saved hash no longer
// matches the recomputed body hash — i.e. the user has edited the
// generated function body. In that case the write is skipped and an
// slog.Info event is emitted so operators can trace which files were
// retained across a regenerate.
//
// The helper lives in the fffile leaf package so both the parent gogin
// package (WriteManyFiles) and child emitters (ssac.writeMethodFile,
// ssac.emitConvertFuncFile, …) can route their writes through it without
// creating an import cycle.
func WriteIfNotPreserved(path string, content []byte) error {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return os.WriteFile(path, content, 0o644)
		}
		return err
	}
	state, derr := contract.DetectPreserved(path)
	if derr == nil && state == contract.StatePreserved {
		slog.Info("skipping preserved file", "path", path)
		return nil
	}
	return os.WriteFile(path, content, 0o644)
}
