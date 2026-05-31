//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderModels — DDL Table → SQLAlchemy 모델 소스 생성(헤더·클래스·컬럼·FK·table_args) 검증
package models

import (
	"strings"
	"testing"
)

func TestRenderModelsEmpty(t *testing.T) {
	out, err := RenderModels(nil)
	if err != nil {
		t.Fatalf("RenderModels(nil) error: %v", err)
	}
	if !strings.Contains(out, "from app.database import Base") {
		t.Errorf("header missing for empty input:\n%s", out)
	}
}
