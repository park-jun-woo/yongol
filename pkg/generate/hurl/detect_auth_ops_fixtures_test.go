//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what detectAuthOpsFixtures — TestDetectAuthOps 용 table-driven 케이스 빌더

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// detectAuthOpsFixtures returns the full shape-detection test matrix.
// Extracted from TestDetectAuthOps to keep the test function under the
// Q4 PURE line budget (range body <= 10 lines of non-control code).
func detectAuthOpsFixtures() []detectAuthOpsFixture {
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

	return []detectAuthOpsFixture{
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
}
