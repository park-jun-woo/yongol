//ff:func feature=cli-init type=test control=selection
//ff:what TestLoadSkeletonFileContents — rendered/verbatim 성공 + render/read 에러 분기 검증

package cliinit

import (
	"strings"
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

func TestLoadSkeletonFileContents_Verbatim(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/sqlc.yaml", rendered: false}
	out, err := loadSkeletonFileContents(templateData{}, f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) == 0 {
		t.Error("expected verbatim content")
	}
}

func TestLoadSkeletonFileContents_RenderError(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/does-not-exist.tmpl", rendered: true}
	_, err := loadSkeletonFileContents(templateData{}, f)
	if err == nil || !strings.Contains(err.Error(), "render") {
		t.Fatalf("want render error, got %v", err)
	}
}

func TestLoadSkeletonFileContents_ReadError(t *testing.T) {
	f := skeletonFile{srcEmbed: "templates/does-not-exist", rendered: false}
	_, err := loadSkeletonFileContents(templateData{}, f)
	if err == nil || !strings.Contains(err.Error(), "read embedded") {
		t.Fatalf("want read error, got %v", err)
	}
}
