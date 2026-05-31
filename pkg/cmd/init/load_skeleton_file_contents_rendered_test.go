//ff:func feature=cli-init type=test control=sequence
//ff:what TestLoadSkeletonFileContents — rendered/verbatim 성공 + render/read 에러 분기 검증
package cliinit

import (
	"testing"
)

func TestLoadSkeletonFileContents_Rendered(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/manifest.yaml.tmpl", rendered: true}
	data := templateData{ProjectID: "MyApp", ProjectIDNormalized: "myapp", Description: "d"}
	out, err := loadSkeletonFileContents(data, f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) == 0 {
		t.Error("expected rendered content")
	}
}
