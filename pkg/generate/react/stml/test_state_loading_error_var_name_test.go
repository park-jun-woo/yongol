//ff:func feature=stml-gen type=test control=sequence
//ff:what data-state loading/error 변수명 생성 규칙 검증
package stml

import ("strings"; "testing"; stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml")

func TestStateLoadingErrorVarName(t *testing.T) {
	page, _ := stmlparser.ParseReader("course-list-page.html", strings.NewReader(`<main>
  <section data-fetch="ListCourses">
    <ul data-each="courses">
      <li><span data-bind="name"></span></li>
    </ul>
    <div data-state="courses.loading" class="text-gray-500">로딩 중...</div>
    <div data-state="courses.error" class="text-red-500">불러오지 못했습니다</div>
  </section>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "listCoursesDataLoading")
	assertContains(t, code, "listCoursesDataError")
	assertNotContains(t, code, "coursesLoading")
	assertNotContains(t, code, "coursesError")
}
