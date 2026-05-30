//ff:func feature=cli-init type=test control=sequence
//ff:what TestSkeletonFiles — 모든 srcEmbed가 embed.FS에서 읽히고 dest 비어있지 않음을 검증

package cliinit

import "testing"

func TestSkeletonFiles(t *testing.T) {
	files := skeletonFiles()
	if len(files) == 0 {
		t.Fatal("expected non-empty skeleton file list")
	}
	for _, f := range files {
		if f.srcEmbed == "" || f.destRel == "" {
			t.Errorf("skeleton file has empty field: %+v", f)
		}
		// Every srcEmbed must resolve in the embed.FS.
		if _, err := templateFiles.ReadFile(f.srcEmbed); err != nil {
			t.Errorf("srcEmbed %q not in embed.FS: %v", f.srcEmbed, err)
		}
	}
}
