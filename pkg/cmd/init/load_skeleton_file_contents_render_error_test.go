//ff:func feature=cli-init type=test control=sequence
//ff:what TestLoadSkeletonFileContents — rendered/verbatim 성공 + render/read 에러 분기 검증
package cliinit

import (
	"strings"
	"testing"
)

func TestLoadSkeletonFileContents_RenderError(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/does-not-exist.tmpl", rendered: true}
	_, err := loadSkeletonFileContents(templateData{}, f)
	if err == nil || !strings.Contains(err.Error(), "render") {
		t.Fatalf("want render error, got %v", err)
	}
}
