//ff:func feature=stml-gen type=test control=sequence
//ff:what 로그인 페이지 TSX 생성을 검증
package stml

import ("strings"; "testing"; stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml")

func TestGenerateLoginPage(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main class="flex items-center justify-center min-h-screen">
  <div data-action="Login" class="space-y-4">
    <input data-field="Email" type="email" placeholder="이메일" class="w-full px-3 py-2 border rounded" />
    <input data-field="Password" type="password" placeholder="비밀번호" class="w-full px-3 py-2 border rounded" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "export default function LoginPage()")
	assertContains(t, code, "useMutation")
	assertContains(t, code, "api.Login")
	assertContains(t, code, `placeholder="이메일"`)
	assertContains(t, code, `placeholder="비밀번호"`)
	assertContains(t, code, `className="w-full px-3 py-2 border rounded"`)
	assertContains(t, code, `type="email"`)
	assertContains(t, code, `type="password"`)
	assertContains(t, code, ">로그인</")
	assertNotContains(t, code, ">제출</")
	assertNotContains(t, code, "useQuery(")
	assertNotContains(t, code, "useParams")
}
