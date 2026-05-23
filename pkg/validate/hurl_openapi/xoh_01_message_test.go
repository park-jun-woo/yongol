//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what xoh01Message — path 존재 여부에 따른 XOH-01 진단 문구 선택 검증

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh01Message(t *testing.T) {
	routes := []apiRoute{
		{Path: "/users", Method: "GET", Segments: []string{"users"}},
		{Path: "/users", Method: "POST", Segments: []string{"users"}},
	}

	t.Run("path_not_found", func(t *testing.T) {
		e := hurl.HurlEntry{Method: "GET", Path: "/orders"}
		msg, advice := xoh01Message(e, []string{"orders"}, routes)
		if !strings.Contains(msg, "path not declared") {
			t.Errorf("expected 'path not declared' in msg, got %q", msg)
		}
		if !strings.Contains(advice, "Add a matching operation") {
			t.Errorf("expected matching advice, got %q", advice)
		}
	})

	t.Run("method_not_found_on_existing_path", func(t *testing.T) {
		e := hurl.HurlEntry{Method: "DELETE", Path: "/users"}
		msg, advice := xoh01Message(e, []string{"users"}, routes)
		if !strings.Contains(msg, "method not declared") {
			t.Errorf("expected 'method not declared' in msg, got %q", msg)
		}
		if !strings.Contains(msg, "GET, POST") {
			t.Errorf("expected available methods in msg, got %q", msg)
		}
		if !strings.Contains(advice, "DELETE") {
			t.Errorf("expected method in advice, got %q", advice)
		}
	})
}
