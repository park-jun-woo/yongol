//ff:func feature=validate-contract type=test control=selection topic=preserve-safety
//ff:what TestIsResourceAcquireCall — CallExpr 가 close 필요 리소스 반환 호출인지 판정 검증

package contract

import "testing"

func TestIsResourceAcquireCall(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"os.Open", "os.Open(p)", true},
		{"os.Create", "os.Create(p)", true},
		{"os.OpenFile", "os.OpenFile(p, 0, 0)", true},
		{"http.Get", "http.Get(u)", true},
		{"http.Post", "http.Post(u, ct, body)", true},
		{"client.Do", "client.Do(req)", true},
		{"db.Query", "db.Query(q)", true},
		{"tx.QueryContext", "tx.QueryContext(ctx, q)", true},
		{"open on wrong recv", "fs.Open(p)", false},
		{"get on wrong recv", "svc.Get(k)", false},
		{"queryrow excluded", "db.QueryRow(q)", false},
		{"free function", "Open(p)", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isResourceAcquireCall(mustCall(t, tt.src)); got != tt.want {
				t.Fatalf("isResourceAcquireCall(%q) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
	if isResourceAcquireCall(nil) {
		t.Fatal("nil call should return false")
	}
}
