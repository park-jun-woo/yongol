//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what XQS-19 — SSaC built-in 호출이 대응 sqlc 쿼리를 요구하는지 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// cacheInterface returns a minimal ssacmeta.PackageInterface fixture for the
// ssac cache package. Mirrors the real ssac/pkg/cache/interface.yaml so the
// unit test stays faithful to the authoritative catalog.
func cacheInterface() *ssacmeta.PackageInterface {
	return &ssacmeta.PackageInterface{
		Package: "cache",
		Ports: []ssacmeta.Port{
			{Name: "CacheSet", UsedBy: []string{"Set"}},
			{Name: "CacheGet", UsedBy: []string{"Get"}},
			{Name: "CacheDelete", UsedBy: []string{"Delete"}},
		},
	}
}

func queueInterface() *ssacmeta.PackageInterface {
	return &ssacmeta.PackageInterface{
		Package: "queue",
		Ports: []ssacmeta.Port{
			{Name: "QueuePublish", UsedBy: []string{"Publish", "PublishTx"}},
			{Name: "QueuePoll", UsedBy: []string{"Subscribe"}},
			{Name: "QueueAck", UsedBy: []string{"Subscribe"}},
		},
	}
}

func TestXqs19_Cache_Present_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "PutEntry", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "call", Package: "cache", Model: "Set"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"cache": cacheInterface()},
		SQLcQueries:    []sqlcparser.QuerySpec{{Name: "CacheSet"}},
	}
	if diags := xqs19SsacBuiltinQueryRequired(fs); len(diags) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", diags)
	}
}

func TestXqs19_Cache_Missing_Query_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "PutEntry", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "call", Package: "cache", Model: "Set"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"cache": cacheInterface()},
	}
	diags := xqs19SsacBuiltinQueryRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "CacheSet") {
		t.Errorf("diag missing CacheSet: %s", diags[0].Message)
	}
}

func TestXqs19_NonDBBuiltin_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Hash", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "call", Package: "crypto", Model: "Hash"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"cache": cacheInterface()},
	}
	if diags := xqs19SsacBuiltinQueryRequired(fs); len(diags) != 0 {
		t.Errorf("crypto.Hash is not DB-using, expected no diag: %+v", diags)
	}
}

func TestXqs19_Publish_Missing_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Enqueue", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "publish", Topic: "order.created"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"queue": queueInterface()},
	}
	diags := xqs19SsacBuiltinQueryRequired(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "QueuePublish") {
		t.Errorf("expected single XQS-19 naming QueuePublish, got: %+v", diags)
	}
}

func TestXqs19_Subscribe_Poll_And_Ack_Missing_Raises_Two(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "OnOrder", FileName: "svc.ssac",
			Subscribe: &ssacparser.SubscribeInfo{Topic: "order.created"},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"queue": queueInterface()},
	}
	diags := xqs19SsacBuiltinQueryRequired(fs)
	// Two ports (QueuePoll + QueueAck) both use Subscribe; both are missing.
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics (QueuePoll + QueueAck), got %d: %+v", len(diags), diags)
	}
}
