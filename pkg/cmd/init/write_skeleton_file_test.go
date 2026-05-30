//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteSkeletonFile — 정상 기록 / load 에러 / write 에러(dest 디렉 부재) 분기 검증

package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSkeletonFile_Success(t *testing.T) {
	target := t.TempDir()
	f := skeletonFile{srcEmbed: "templates/sqlc.yaml", destRel: "sqlc.yaml", rendered: false}
	if err := writeSkeletonFile(target, templateData{}, f); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(target, "sqlc.yaml")); err != nil {
		t.Errorf("file not written: %v", err)
	}
}

func TestWriteSkeletonFile_LoadError(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/missing", destRel: "x", rendered: false}
	if err := writeSkeletonFile(t.TempDir(), templateData{}, f); err == nil {
		t.Fatal("want load error for missing embed")
	}
}

func TestWriteSkeletonFile_WriteError(t *testing.T) {
	// destRel points into a non-existent subdir -> WriteFile fails.
	f := skeletonFile{srcEmbed: "templates/sqlc.yaml", destRel: "no/such/dir/sqlc.yaml", rendered: false}
	if err := writeSkeletonFile(t.TempDir(), templateData{}, f); err == nil {
		t.Fatal("want write error when dest dir missing")
	}
}
