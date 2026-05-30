//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what zz_zerocov — 0% 커버리지 STML 파서 함수(파일/디렉토리 I/O + 정적 래퍼 분기) 검증

package stml

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// richPageHTML exercises fetch/action/each blocks with nested static wrappers,
// fields, components, binds and submit buttons so that the static-dispatch and
// walk helpers (parseStaticInFetch/Action/Each, dispatchStatic*Child,
// handleStaticAction*, walkStaticAction*, handleFetchComponent, etc.) all run.
const richPageHTML = `<main data-route="/things/:id">
  <article data-fetch="GetThing" data-param-id="route.id">
    <div class="wrapper">
      <h2 data-field="Name"></h2>
      <div data-component="Avatar" data-field="Owner" class="av"></div>
      <span data-bind="Status"></span>
      <p class="plain">just static text</p>
    </div>
    <section data-each="Items">
      <div class="row">
        <div class="cell">
          <span data-bind="Title"></span>
          <em class="muted">label</em>
        </div>
      </div>
    </section>
    <div data-component="DirectCard" data-field="Card" class="dc"></div>
    <form data-action="UpdateThing">
      <fieldset>
        <input data-field="Title" type="text" placeholder="title" class="inp" />
        <div data-component="Picker" data-field="Tag" class="pk">
          <input data-field="Inner" type="text" />
        </div>
        <div class="static-action-wrap">
          <input data-field="Nested" type="text" />
          <div data-component="WalkCard" data-field="Walk" class="wc">
            <div data-component="InnerWalk" data-field="Inner2" class="iw">
              <input data-field="DeepNested" type="text" />
            </div>
          </div>
          <button type="submit">Save</button>
        </div>
      </fieldset>
    </form>
  </article>
  <footer class="page-footer">
    <p>static footer</p>
  </footer>
</main>`

//ff:what TestParseRichPage_ZeroCov — 정적 래퍼/필드/컴포넌트/each/action 분기 전부 도달
func TestParseRichPage_ZeroCov(t *testing.T) {
	page, diags := ParseReader("things.html", strings.NewReader(richPageHTML))
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if page.Route != "/things/:id" {
		t.Errorf("Route = %q, want /things/:id", page.Route)
	}
	if len(page.Fetches) == 0 {
		t.Fatalf("expected at least one fetch block")
	}
	fb := page.Fetches[0]
	if fb.OperationID != "GetThing" {
		t.Errorf("fetch OperationID = %q, want GetThing", fb.OperationID)
	}
	// Component collected via handleFetchComponent.
	foundComp := false
	for _, c := range fb.Components {
		if c.Name == "Avatar" {
			foundComp = true
		}
	}
	if !foundComp {
		t.Errorf("expected Avatar component collected, got %+v", fb.Components)
	}
	if len(fb.Eaches) == 0 {
		t.Errorf("expected each block collected")
	}
}

//ff:what TestParseFile_ZeroCov — ParseFile 정상/에러 경로
func TestParseFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "page.html")
	if err := os.WriteFile(good, []byte(richPageHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	page, diags := ParseFile(good)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if page.Route != "/things/:id" {
		t.Errorf("Route = %q", page.Route)
	}

	// Error path: missing file.
	_, diags = ParseFile(filepath.Join(dir, "nope.html"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing file")
	}
}

//ff:what TestParseDir_ZeroCov — ParseDir 정상/에러 경로
func TestParseDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.html"), []byte(richPageHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	// Non-html and a subdir should be skipped.
	if err := os.WriteFile(filepath.Join(dir, "ignore.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	pages, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(pages) != 1 {
		t.Errorf("pages = %d, want 1", len(pages))
	}

	// Error path: missing dir.
	_, diags = ParseDir(filepath.Join(dir, "missing"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing dir")
	}
}

const layoutHTML = `<div>
  <nav>
    <a data-nav="/home">Home</a>
  </nav>
  <slot data-outlet />
</div>`

//ff:what TestParseLayoutFile_ZeroCov — ParseLayoutFile 정상/에러 경로
func TestParseLayoutFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "app.html")
	if err := os.WriteFile(good, []byte(layoutHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	layout, diags := ParseLayoutFile(good)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if layout.Name != "app" {
		t.Errorf("Name = %q, want app", layout.Name)
	}

	_, diags = ParseLayoutFile(filepath.Join(dir, "nope.html"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing layout file")
	}
}

//ff:what TestParseLayoutDir_ZeroCov — ParseLayoutDir 정상/에러 경로
func TestParseLayoutDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "app.html"), []byte(layoutHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "skip.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	layouts, diags := ParseLayoutDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(layouts) != 1 {
		t.Errorf("layouts = %d, want 1", len(layouts))
	}

	_, diags = ParseLayoutDir(filepath.Join(dir, "missing"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing layout dir")
	}
}
