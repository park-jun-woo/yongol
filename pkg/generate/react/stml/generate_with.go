//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 지정된 Target으로 페이지 목록을 순회하며 파일을 생성한다
package stml

import (
	"fmt"
	"os"
	"path/filepath"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// GenerateWith produces files using the given Target.
func GenerateWith(t Target, pages []stmlparser.PageSpec, specsDir, outDir string, opts ...GenerateOptions) (*GenerateResult, error) {
	opt := DefaultOptions()
	if len(opts) > 0 {
		opt = mergeOpt(opt, opts[0])
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return nil, fmt.Errorf("mkdir %s: %w", outDir, err)
	}

	for _, page := range pages {
		code := t.GeneratePage(page, specsDir, opt)
		path := filepath.Join(outDir, page.Name+t.FileExtension())
		if err := os.WriteFile(path, []byte(code), 0o644); err != nil {
			return nil, fmt.Errorf("write %s: %w", path, err)
		}
	}

	return &GenerateResult{
		Pages:        len(pages),
		Dependencies: t.Dependencies(pages),
	}, nil
}
