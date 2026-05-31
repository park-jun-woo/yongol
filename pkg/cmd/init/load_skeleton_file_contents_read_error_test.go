//ff:func feature=cli-init type=test control=sequence
//ff:what TestLoadSkeletonFileContents — rendered/verbatim 성공 + render/read 에러 분기 검증
package cliinit

import (
	"strings"
	"testing"
)

func TestLoadSkeletonFileContents_ReadError(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/does-not-exist", rendered: false}
	_, err := loadSkeletonFileContents(templateData{}, f)
	if err == nil || !strings.Contains(err.Error(), "read embedded") {
		t.Fatalf("want read error, got %v", err)
	}
}
