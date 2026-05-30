package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부

func TestHasActiveBlock_ZeroCov(t *testing.T) {
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true},
		{Name: "auth", Active: false},
	}}
	if !hasActiveBlock(plan, "cors") {
		t.Error("expected cors active")
	}
	if hasActiveBlock(plan, "auth") {
		t.Error("auth should not be active")
	}
}

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderConfig_ZeroCov — config.py 생성

func TestRenderConfig_ZeroCov(t *testing.T) {
	out, err := RenderConfig()
	if err != nil {
		t.Fatalf("RenderConfig error: %v", err)
	}
	for _, want := range []string{"BaseSettings", "database_url", "jwt_secret", "settings = Settings()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderConfig missing %q", want)
		}
	}
}

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderDatabase_ZeroCov — database.py 생성

func TestRenderDatabase_ZeroCov(t *testing.T) {
	out, err := RenderDatabase()
	if err != nil {
		t.Fatalf("RenderDatabase error: %v", err)
	}
	for _, want := range []string{"create_async_engine", "async_session", "class Base(DeclarativeBase)"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderDatabase missing %q", want)
		}
	}
}

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderMain_ZeroCov — main.py + nil plan 에러 + cors 분기

func TestRenderMain_ZeroCov(t *testing.T) {
	if _, err := RenderMain(nil, nil); err == nil {
		t.Error("expected error for nil plan")
	}
	plan := &ir.BootPlan{
		ProjectID:    "myapp",
		ActiveBlocks: []ir.BootBlock{{Name: "cors", Active: true}},
	}
	out, err := RenderMain(plan, []string{"users"})
	if err != nil {
		t.Fatalf("RenderMain error: %v", err)
	}
	for _, want := range []string{"FastAPI", "myapp", "CORSMiddleware", "users_router", "/health"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderMain missing %q", want)
		}
	}
}
