//ff:func feature=stml-gen type=test control=sequence
//ff:what GeneratePage — DocumentTitles 등재 페이지만 document.title useEffect 방출 / useState 와 react import 병합 / 미등재 무방출 검증

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_TitleEffect(t *testing.T) {
	const src = `<main>
  <section data-fetch="ListBuildings">
    <span data-bind="total"></span>
  </section>
</main>`

	t.Run("listed page emits the mount title effect", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("building-list.html", strings.NewReader(src))
		code := GeneratePage(page, "", GenerateOptions{
			DocumentTitles: map[string]string{"building-list": "건물 목록 · zenflow"},
		})
		assertContains(t, code, "import { useEffect } from 'react'")
		assertContains(t, code, "export default function BuildingList() {\n  useEffect(() => {\n    document.title = '건물 목록 · zenflow'\n  }, [])\n")
	})

	t.Run("useEffect merges into the existing useState react import", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("building-new.html", strings.NewReader(`<main>
  <div data-action="CreateBuilding">
    <input data-field="Name" type="text" />
    <button type="submit">등록</button>
  </div>
</main>`))
		code := GeneratePage(page, "", GenerateOptions{
			DocumentTitles: map[string]string{"building-new": "건물 등록 · zenflow"},
		})
		assertContains(t, code, "import { useEffect, useState } from 'react'")
	})

	t.Run("unlisted page emits no effect — sitemap-absent output stays byte-identical", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("building-list.html", strings.NewReader(src))
		code := GeneratePage(page, "", GenerateOptions{})
		assertNotContains(t, code, "useEffect")
		assertNotContains(t, code, "document.title")
	})
}
