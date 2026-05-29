//ff:func feature=stml-gen type=test control=sequence
//ff:what file input에 as any 캐스팅이 생성되는지 검증

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFieldFileInputAsAny(t *testing.T) {
	page, _ := stmlparser.ParseReader("upload.html", strings.NewReader(`<main>
  <div data-action="UploadFile">
    <input data-field="files" type="file" />
    <button type="submit">업로드</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `{...(uploadFileForm.register('files') as any)}`)
}
