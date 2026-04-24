//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestDetectAuthOps — SSaC shape 기반 signup/login 감지 table-driven 검증 (BUG-023 회귀 방지)

package hurl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestDetectAuthOps exercises the shape-detection matrix described in
// plans/gen/hurl02/Phase003-AuthOpShapeDetection.md. Naming variance
// (Register / Signup / Join / EnrollStudent / SignIn) must not affect
// classification — only the SSaC body shape does.
func TestDetectAuthOps(t *testing.T) {
	type fixture struct {
		name         string
		opID         string
		public       bool
		hasPassword  bool
		funcBuilder  func(opID string) *ssac.ServiceFunc
		wantRole     string // "signup" | "login" | "" (neither)
		wantWarnSub  string // substring expected in warnings; "" = no specific warning
	}

	mkSignup := func(opID string) *ssac.ServiceFunc {
		fn := signupServiceFunc(opID)
		return &fn
	}
	mkLogin := func(opID string) *ssac.ServiceFunc {
		fn := loginServiceFunc(opID)
		return &fn
	}
	mkCombined := func(opID string) *ssac.ServiceFunc {
		return &ssac.ServiceFunc{
			Name: opID,
			Sequences: []ssac.Sequence{
				{Type: ssac.SeqCall, Model: "auth.HashPassword"},
				{Type: ssac.SeqPost, Model: "User.Create", Inputs: map[string]string{"PasswordHash": "hp.HashedPassword"}},
				{Type: ssac.SeqVerifyPassword, Model: "User"},
			},
		}
	}
	mkNoCreate := func(opID string) *ssac.ServiceFunc {
		return &ssac.ServiceFunc{
			Name: opID,
			Sequences: []ssac.Sequence{
				{Type: ssac.SeqCall, Model: "auth.HashPassword"},
			},
		}
	}
	mkEmpty := func(opID string) *ssac.ServiceFunc {
		return &ssac.ServiceFunc{Name: opID}
	}

	cases := []fixture{
		{name: "Register/signup", opID: "Register", public: true, hasPassword: true, funcBuilder: mkSignup, wantRole: "signup"},
		{name: "Signup/signup", opID: "Signup", public: true, hasPassword: true, funcBuilder: mkSignup, wantRole: "signup"},
		{name: "Join/signup", opID: "Join", public: true, hasPassword: true, funcBuilder: mkSignup, wantRole: "signup"},
		{name: "EnrollStudent/signup", opID: "EnrollStudent", public: true, hasPassword: true, funcBuilder: mkSignup, wantRole: "signup"},
		{name: "Login/login", opID: "Login", public: true, hasPassword: true, funcBuilder: mkLogin, wantRole: "login"},
		{name: "SignIn/login", opID: "SignIn", public: true, hasPassword: true, funcBuilder: mkLogin, wantRole: "login"},
		{name: "CheckEmail/query-not-auth", opID: "CheckEmail", public: true, hasPassword: false, funcBuilder: mkEmpty, wantRole: ""},
		{name: "UpdateProfile/non-public", opID: "UpdateProfile", public: false, hasPassword: false, funcBuilder: mkEmpty, wantRole: ""},
		{name: "Combined/signup+warn", opID: "ComboSignupLogin", public: true, hasPassword: true, funcBuilder: mkCombined, wantRole: "signup", wantWarnSub: "combined signup"},
		{name: "SignupWithoutCreate/signup+warn", opID: "LoneSignup", public: true, hasPassword: true, funcBuilder: mkNoCreate, wantRole: "signup", wantWarnSub: "no companion @post"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			op := newPasswordOp(tc.opID, tc.public, tc.hasPassword)
			doc := &openapi3.T{Paths: openapi3.NewPaths(
				openapi3.WithPath("/auth/"+tc.opID, &openapi3.PathItem{Post: op}),
			)}
			var funcs []ssac.ServiceFunc
			if fn := tc.funcBuilder(tc.opID); fn != nil {
				funcs = []ssac.ServiceFunc{*fn}
			}
			fs := &yongol.Fullstack{OpenAPIDoc: doc, ServiceFuncs: funcs}
			signup, login, warns := detectAuthOps(fs)
			got := ""
			switch {
			case signup != nil && signup.OpID == tc.opID:
				got = "signup"
			case login != nil && login.OpID == tc.opID:
				got = "login"
			}
			if got != tc.wantRole {
				t.Errorf("role: want %q got %q (warns=%v)", tc.wantRole, got, warns)
			}
			if tc.wantWarnSub != "" {
				found := false
				for _, w := range warns {
					if strings.Contains(w, tc.wantWarnSub) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected warning substring %q; got %v", tc.wantWarnSub, warns)
				}
			}
		})
	}
}

// TestDetectAuthOpsMultipleSignupCandidates pins the deterministic
// picker: two signup-shape ops → alphabetical-first wins + WARNING.
func TestDetectAuthOpsMultipleSignupCandidates(t *testing.T) {
	admin := newPasswordOp("AdminSignup", true, true)
	user := newPasswordOp("UserSignup", true, true)
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/admin-signup", &openapi3.PathItem{Post: admin}),
		openapi3.WithPath("/auth/user-signup", &openapi3.PathItem{Post: user}),
	)}
	adminFn := signupServiceFunc("AdminSignup")
	userFn := signupServiceFunc("UserSignup")
	fs := &yongol.Fullstack{OpenAPIDoc: doc, ServiceFuncs: []ssac.ServiceFunc{adminFn, userFn}}
	signup, _, warns := detectAuthOps(fs)
	if signup == nil {
		t.Fatalf("expected a signup pick, got nil (warns=%v)", warns)
	}
	if signup.OpID != "AdminSignup" {
		t.Errorf("signup pick: want AdminSignup (alphabetical first), got %q", signup.OpID)
	}
	sawMulti := false
	for _, w := range warns {
		if strings.Contains(w, "multiple signup candidates") {
			sawMulti = true
			break
		}
	}
	if !sawMulti {
		t.Errorf("expected 'multiple signup candidates' warning; got %v", warns)
	}
}

// newPasswordOp returns a minimal *openapi3.Operation for shape-
// detection tests. public=true sets Security to an empty non-nil slice
// (the OpenAPI "explicit public override"). hasPassword=true wires a
// JSON body with a "password" property.
func newPasswordOp(opID string, public, hasPassword bool) *openapi3.Operation {
	op := &openapi3.Operation{OperationID: opID}
	if public {
		emptySec := openapi3.SecurityRequirements{}
		op.Security = &emptySec
	}
	if hasPassword {
		body := &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: openapi3.Schemas{
				"email":    {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				"password": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			},
		}
		op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(body))}
	}
	op.Responses = newOKResponses()
	return op
}
