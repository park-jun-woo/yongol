//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestByNameOpDomainMap_ZeroCov(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "ListItems"}}),
	)}
	docs := []domainDoc{{Name: "core", Doc: doc, Cfg: manifest.DomainConfig{OpenAPI: "core.yaml"}}}

	m := buildOpDomainMap(docs)
	if m["ListItems"] != "core" {
		t.Errorf("buildOpDomainMap = %v", m)
	}

	result := map[string]string{}
	collectDocOpDomains(docs[0], result)
	if result["ListItems"] != "core" {
		t.Errorf("collectDocOpDomains = %v", result)
	}

	// checkUnconsumedOps: ListItems not consumed → diagnostic.
	consumed := map[string]struct{}{}
	if d := checkUnconsumedOps(docs[0], consumed, "XMO-21", "core"); len(d) == 0 {
		t.Errorf("checkUnconsumedOps expected diagnostic for unconsumed op")
	}
	// consumed → no diagnostic.
	consumed["ListItems"] = struct{}{}
	if d := checkUnconsumedOps(docs[0], consumed, "XMO-21", "core"); len(d) != 0 {
		t.Errorf("checkUnconsumedOps unexpected diagnostics: %v", d)
	}
}
