package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부

func TestHasActiveBlock_ZeroCov(t *testing.T) {
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true},
		{Name: "session", Active: false},
	}}
	if !hasActiveBlock(plan, "cors") {
		t.Error("expected cors active")
	}
	if hasActiveBlock(plan, "session") {
		t.Error("session should not be active")
	}
	if hasActiveBlock(plan, "missing") {
		t.Error("missing should not be active")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestExtractCORSConfig_ZeroCov — CORS config 추출 + 기본값

func TestExtractCORSConfig_ZeroCov(t *testing.T) {
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true, Config: &ir.CORSBootConfig{
			AllowOrigins:     []string{"https://a.com"},
			AllowCredentials: true,
		}},
	}}
	origins, creds := extractCORSConfig(plan)
	if len(origins) != 1 || origins[0] != "https://a.com" || !creds {
		t.Errorf("extractCORSConfig = %v, %v", origins, creds)
	}
	// No cors block → default nil, true
	o2, c2 := extractCORSConfig(&ir.BootPlan{})
	if o2 != nil || !c2 {
		t.Errorf("default = %v, %v", o2, c2)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderCorsBlock_ZeroCov — origin 있음/없음 두 분기

func TestRenderCorsBlock_ZeroCov(t *testing.T) {
	// No origins → simple
	var b1 strings.Builder
	renderCorsBlock(&b1, &ir.BootPlan{})
	if !strings.Contains(b1.String(), "credentials: true") {
		t.Errorf("simple cors missing:\n%s", b1.String())
	}
	// With origins
	var b2 strings.Builder
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true, Config: &ir.CORSBootConfig{
			AllowOrigins: []string{"https://a.com"}, AllowCredentials: true,
		}},
	}}
	renderCorsBlock(&b2, plan)
	out := b2.String()
	if !strings.Contains(out, "origin: [") || !strings.Contains(out, "https://a.com") {
		t.Errorf("origin cors missing:\n%s", out)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderAppModule_ZeroCov — app.module.ts (infra + feature 모듈)

func TestRenderAppModule_ZeroCov(t *testing.T) {
	out, err := RenderAppModule([]string{"billing"}, []string{"queue"})
	if err != nil {
		t.Fatalf("RenderAppModule error: %v", err)
	}
	for _, want := range []string{"PrismaModule", "QueueModule", "BillingModule", "export class AppModule"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderAppModule missing %q", want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderMain_ZeroCov — main.ts + nil plan 에러 + cors 분기

func TestRenderMain_ZeroCov(t *testing.T) {
	if _, err := RenderMain(nil); err == nil {
		t.Error("expected error for nil plan")
	}
	plan := &ir.BootPlan{
		ProjectID: "myapp",
		ActiveBlocks: []ir.BootBlock{
			{Name: "cors", Active: true, Config: &ir.CORSBootConfig{AllowOrigins: []string{"https://a.com"}}},
		},
	}
	out, err := RenderMain(plan)
	if err != nil {
		t.Fatalf("RenderMain error: %v", err)
	}
	for _, want := range []string{"NestFactory", "bootstrap()", "myapp", "enableCors"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderMain missing %q", want)
		}
	}
}
