//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what itemParamField — item.*/점 경로/비-item 소스 분기 검증

package stml_openapi

import "testing"

func TestItemParamField(t *testing.T) {
	if f, ok := itemParamField("item.id"); !ok || f != "id" {
		t.Errorf("item.id → %q, %v", f, ok)
	}
	if f, ok := itemParamField("item.photo.id"); !ok || f != "photo" {
		t.Errorf("item.photo.id → %q, %v", f, ok)
	}
	if f, ok := itemParamField("item."); !ok || f != "" {
		t.Errorf("item. → %q, %v", f, ok)
	}
	if _, ok := itemParamField("route.BuildingID"); ok {
		t.Error("route.* must not match")
	}
	if _, ok := itemParamField(""); ok {
		t.Error("empty source must not match")
	}
}
