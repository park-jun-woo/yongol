package scaffold

import (
	"strings"
	"testing"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestDependencies_ZeroCov — 런타임 의존성 맵

func TestNestDependencies_ZeroCov(t *testing.T) {
	deps := nestDependencies()
	if deps["@nestjs/core"] == "" {
		t.Error("expected @nestjs/core dependency")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestDevDependencies_ZeroCov — 개발 의존성 맵

func TestNestDevDependencies_ZeroCov(t *testing.T) {
	deps := nestDevDependencies()
	if deps["@nestjs/cli"] == "" {
		t.Error("expected @nestjs/cli dev dependency")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderNestCLI_ZeroCov — nest-cli.json 생성

func TestRenderNestCLI_ZeroCov(t *testing.T) {
	out, err := RenderNestCLI()
	if err != nil {
		t.Fatalf("RenderNestCLI error: %v", err)
	}
	for _, want := range []string{"@nestjs/schematics", "sourceRoot", "deleteOutDir"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderNestCLI missing %q", want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderPackageJSON_ZeroCov — package.json 생성 + 빈 projectID 에러

func TestRenderPackageJSON_ZeroCov(t *testing.T) {
	if _, err := RenderPackageJSON(""); err == nil {
		t.Error("expected error for empty projectID")
	}
	out, err := RenderPackageJSON("myapp")
	if err != nil {
		t.Fatalf("RenderPackageJSON error: %v", err)
	}
	for _, want := range []string{`"name": "myapp"`, "nest build", "@nestjs/core", "prisma"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderPackageJSON missing %q", want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderTSConfig_ZeroCov — tsconfig.json 생성

func TestRenderTSConfig_ZeroCov(t *testing.T) {
	out, err := RenderTSConfig()
	if err != nil {
		t.Fatalf("RenderTSConfig error: %v", err)
	}
	for _, want := range []string{`"module": "commonjs"`, `"target": "ES2021"`, "experimentalDecorators"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderTSConfig missing %q", want)
		}
	}
}
