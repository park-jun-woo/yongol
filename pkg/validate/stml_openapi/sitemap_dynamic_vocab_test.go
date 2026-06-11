//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestSitemapDynamicVocab — 선언된 동적 속성명 수집 (없음=빈, 부분=해당 속성만, 전부=5종 고정 순서) 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapDynamicVocab(t *testing.T) {
	if got := sitemapDynamicVocab(stml.SitemapNode{Page: "dashboard"}); len(got) != 0 {
		t.Errorf("static node vocab = %v, want none", got)
	}
	partial := stml.SitemapNode{Each: "items", LabelField: "name"}
	if got := sitemapDynamicVocab(partial); !reflect.DeepEqual(got, []string{"data-each", "data-label-field"}) {
		t.Errorf("partial vocab = %v", got)
	}
	full := stml.SitemapNode{Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LinkParamsRaw: "item.id", LabelField: "name"}
	want := []string{"data-fetch", "data-each", "data-link", "data-link-params", "data-label-field"}
	if got := sitemapDynamicVocab(full); !reflect.DeepEqual(got, want) {
		t.Errorf("full vocab = %v, want %v", got, want)
	}
}
