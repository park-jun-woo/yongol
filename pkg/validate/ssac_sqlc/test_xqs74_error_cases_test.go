//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS74_ErrorCases — non-integer PK (TEXT/UUID) 모델에 @empty/@exists 시 에러 발생 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS74_ErrorCases(t *testing.T) {
	t.Run("TextPKRaisesError", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "RevokeToken", FileName: "revoketoken.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "RefreshToken.FindByHash", Result: &ssacparser.Result{Type: "RefreshToken", Var: "rt"}, Line: 5},
					{Type: "empty", Target: "rt", Message: "Invalid token", Line: 8},
				},
			}},
			DDLTables: []ddl.Table{{
				Name:        "refresh_tokens",
				Columns:     map[string]ddl.Column{"token_hash": {Name: "token_hash", RawType: "TEXT", NotNull: true}, "claims": {Name: "claims", RawType: "JSONB"}, "expires_at": {Name: "expires_at", RawType: "TIMESTAMPTZ"}},
				ColumnOrder: []string{"token_hash", "claims", "expires_at"},
				PrimaryKey:  []string{"token_hash"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		d := diags[0]
		if !strings.Contains(d.Message, "XQS-74") {
			t.Errorf("expected XQS-74 rule ID, got: %s", d.Message)
		}
		if !strings.Contains(d.Message, "RefreshToken") {
			t.Errorf("expected model name, got: %s", d.Message)
		}
		if !strings.Contains(d.Message, "TEXT") {
			t.Errorf("expected RawType TEXT, got: %s", d.Message)
		}
		if d.Line != 8 {
			t.Errorf("expected line 8, got %d", d.Line)
		}
	})

	t.Run("UUIDPKRaisesError", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "CreateDevice", FileName: "createdevice.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "Device.FindBySerial", Result: &ssacparser.Result{Type: "Device", Var: "dev"}, Line: 3},
					{Type: "exists", Target: "dev", Message: "Device already exists", ErrStatus: 409, Line: 6},
				},
			}},
			DDLTables: []ddl.Table{{
				Name:        "devices",
				Columns:     map[string]ddl.Column{"device_id": {Name: "device_id", RawType: "UUID", NotNull: true}, "serial": {Name: "serial", RawType: "TEXT"}},
				ColumnOrder: []string{"device_id", "serial"},
				PrimaryKey:  []string{"device_id"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "UUID") {
			t.Errorf("expected UUID in message, got: %s", diags[0].Message)
		}
	})
}
