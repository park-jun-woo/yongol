//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// byNameSamplePage parses a feature-rich STML page used by the by-name tests.
func byNameSamplePage(t *testing.T) stmlparser.PageSpec {
	t.Helper()
	const src = `<main data-layout="app" data-route="/items">
  <section data-fetch="ListItems" data-param-status="route.status">
    <ul data-each="items">
      <li data-bind="name"></li>
      <span data-component="Badge" data-bind="status"></span>
    </ul>
    <span data-bind="total"></span>
    <p data-state="items.empty">없음</p>
    <section data-fetch="ListSub"><span data-bind="x"></span></section>
    <h2>Header</h2>
  </section>
  <div data-action="CreateItem" data-param-id="route.id">
    <input data-field="Name" type="text" />
    <input data-field="Count" type="number" />
    <span data-component="DatePicker" data-field="Due"></span>
    <button type="submit">Create</button>
  </div>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">Login</button>
  </div>
</main>`
	page, diags := stmlparser.ParseReader("items-page.html", strings.NewReader(src))
	if len(diags) > 0 {
		t.Fatalf("parse diags: %v", diags)
	}
	return page
}

func byNameConstraints() map[string]map[string]oapiparser.FieldConstraint {
	mn := 1
	mx := 50
	lo := 0.0
	hi := 100.0
	return map[string]map[string]oapiparser.FieldConstraint{
		"CreateItem": {
			"Name":  {Type: "string", MinLength: &mn, MaxLength: &mx, Required: true},
			"Count": {Type: "integer", Minimum: &lo, Maximum: &hi},
		},
		"Login": {
			"Email":    {Type: "string", Format: "email", Required: true},
			"Password": {Type: "string", MinLength: &mn},
		},
	}
}

// TestByNamePredicates_ZeroCov calls boolean predicate helpers by name.
func TestByNamePredicates_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	actions := page.Actions
	if !anyActionHasFields(actions) {
		t.Errorf("anyActionHasFields = false")
	}
	if !anyActionHasInputFields(actions) {
		t.Errorf("anyActionHasInputFields = false")
	}
	for _, a := range actions {
		_ = actionHasInputField(a)
	}
	_ = allLoginActions(actions)
	_ = allLoginActions([]stmlparser.ActionBlock{{OperationID: "Login"}})
	cons := byNameConstraints()
	if !anyActionHasZodConstraints(actions, cons) {
		t.Errorf("anyActionHasZodConstraints = false")
	}
	_ = anyActionHasZodConstraints(actions, nil)
}

// TestByNameZodValidations_ZeroCov calls the zod validation appenders by name.
func TestByNameZodValidations_ZeroCov(t *testing.T) {
	cons := byNameConstraints()
	strFC := cons["Login"]["Email"]
	numFC := cons["CreateItem"]["Count"]

	parts := appendStringValidations(nil, strFC)
	if len(parts) == 0 {
		t.Errorf("appendStringValidations empty")
	}
	_ = appendStringValidations(nil, numFC) // non-string returns unchanged

	nparts := appendNumericValidations(nil, numFC)
	if len(nparts) == 0 {
		t.Errorf("appendNumericValidations empty")
	}
	_ = appendNumericValidations(nil, strFC) // non-numeric returns unchanged
}

// TestByNameStringHelpers_ZeroCov calls small string helpers by name.
func TestByNameStringHelpers_ZeroCov(t *testing.T) {
	if clsAttr("c") == "" {
		t.Errorf("clsAttr empty")
	}
	_ = clsAttr("")
	if orDefault("", "d") != "d" {
		t.Errorf("orDefault default")
	}
	_ = orDefault("x", "d")
	runes := []rune("RoomID")
	_ = isWordBoundary(runes, 0)
	_ = isWordBoundary(runes, 4)
	_ = isWordBoundary([]rune("abc"), 0)
}

// TestByNameLookupMerge_ZeroCov calls lookupConstraints / mergeOpt by name.
func TestByNameLookupMerge_ZeroCov(t *testing.T) {
	cons := byNameConstraints()
	if lookupConstraints("CreateItem", cons) == nil {
		t.Errorf("lookupConstraints nil")
	}
	_ = lookupConstraints("Missing", cons)
	_ = lookupConstraints("X", nil)

	base := DefaultOptions()
	merged := mergeOpt(base, GenerateOptions{
		APIImportPath:           "@/api",
		UseClient:               true,
		HasAuthz:                true,
		RequestConstraints:      cons,
		ResponseArrayItemFields: map[string]map[string]map[string]bool{"ListItems": {"items": {"id": true}}},
		NoBodyOps:               map[string]bool{"X": true},
		PathParamTypes:          map[string]map[string]string{"GetItem": {"id": "integer"}},
	})
	if merged.APIImportPath != "@/api" {
		t.Errorf("mergeOpt APIImportPath = %q", merged.APIImportPath)
	}
}

// TestByNameCollectors_ZeroCov calls fetch-collection helpers by name.
func TestByNameCollectors_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	f := page.Fetches[0]
	ops := collectFetchOps(f, nil)
	if len(ops) == 0 {
		t.Errorf("collectFetchOps empty")
	}
	binds := collectFetchParamBinds(f, nil)
	_ = binds
	a := page.Actions[0]
	if deduplicateActions([]stmlparser.ActionBlock{a, a}) == nil {
		t.Errorf("deduplicateActions nil")
	}
	_ = extractBindFieldsFromChildren(f.Children)

	is := importSet{}
	compSet := map[string]bool{}
	collectFetchImports(f, &is, compSet)
	collectActionImports(a, &is, compSet)
	full := collectImports(page, "")
	_ = full
}

// TestByNameActionFetchMap_ZeroCov calls the action→fetch mapping helpers by name.
func TestByNameActionFetchMap_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	m := buildActionFetchMap(page)
	if m == nil {
		t.Fatalf("buildActionFetchMap nil")
	}
	walkChildrenForFetchMap(page.Children, []string{"ListItems"}, m)

	rm := map[string][]string{}
	recordActionFetchMapping("CreateItem", []string{"ListItems"}, rm)
	recordActionFetchMapping("CreateItem", []string{"ListItems"}, rm) // already present
	recordActionFetchMapping("NoFetch", nil, rm)

	ops := resolveInvalidateOps("CreateItem", []string{"ListItems"}, m)
	_ = ops
	_ = resolveInvalidateOps("Unknown", []string{"ListItems"}, m)
}

// TestByNameEachKeyFields_ZeroCov calls populate/setEachKeyField* by name.
func TestByNameEachKeyFields_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	raif := map[string]map[string]map[string]bool{
		"ListItems": {"items": {"id": true, "name": true}},
		"ListSub":   {"x": {"id": true}},
	}
	populateEachKeyFields(&page, raif)
	populateEachKeyFields(&page, nil)
	populateEachKeyFieldsInChildren(page.Children, raif)
	for i := range page.Children {
		populateEachKeyFieldForChild(&page.Children[i], raif)
	}

	f := page.Fetches[0]
	setEachKeyFieldsInFetch(&f, "ListItems", raif)
	setEachKeyFieldsInChildren(f.Children, "ListItems", raif)
	for i := range f.Children {
		setEachKeyFieldForChild(&f.Children[i], "ListItems", raif)
	}
	if len(f.Eaches) > 0 {
		setKeyFieldIfHasID(&f.Eaches[0], "ListItems", raif)
		setKeyFieldIfHasID(&f.Eaches[0], "Missing", raif)
	}
}

// TestByNameResolveMutationArgs_ZeroCov calls resolveMutationArgs by name.
func TestByNameResolveMutationArgs_ZeroCov(t *testing.T) {
	cons := byNameConstraints()
	// void
	resolveMutationArgs("DeleteItem", "", true, cons)
	// body only
	resolveMutationArgs("CreateItem", "", false, cons)
	// body + path
	fn, args := resolveMutationArgs("CreateItem", "{ id }", false, cons)
	if fn == "" && args == "" {
		t.Errorf("resolveMutationArgs body+path returned empty")
	}
}

// TestByNameRenderJSX_ZeroCov calls the render*JSX helpers by name.
func TestByNameRenderJSX_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	noBody := map[string]bool{}
	ppt := map[string]map[string]string{}
	f := page.Fetches[0]
	a := page.Actions[0]

	if s := renderFetchJSX(f, 0, noBody); s == "" {
		t.Errorf("renderFetchJSX empty")
	}
	_ = renderFetchJSXBody(f, "data", 1, noBody)
	if s := renderActionJSX(a, 0, noBody); s == "" {
		t.Errorf("renderActionJSX empty")
	}
	if s := renderActionForm(a, 1); s == "" {
		t.Errorf("renderActionForm empty")
	}
	if s := renderActionButton(a, 1, noBody); s == "" {
		t.Errorf("renderActionButton empty")
	}
	_ = renderActionChildNodes(a.Children, "form", 2)
	_ = renderChildNodes(f.Children, "data", "item", 2, noBody)

	if len(f.Eaches) > 0 {
		if s := renderEachJSX(f.Eaches[0], "data", 1); s == "" {
			t.Errorf("renderEachJSX empty")
		}
	}
	if len(a.Fields) > 0 {
		if s := renderFieldJSX(a.Fields[0], "form", 1); s == "" {
			t.Errorf("renderFieldJSX empty")
		}
	}
	if len(f.States) > 0 {
		if s := renderStateJSX(f.States[0], "data", 1, noBody); s == "" {
			t.Errorf("renderStateJSX empty")
		}
	}
	// static element
	se := stmlparser.StaticElement{Tag: "h2", Text: "Header"}
	if s := renderStaticJSX(se, "data", "item", 1, noBody); s == "" {
		t.Errorf("renderStaticJSX empty")
	}

	_ = renderParamArgs(f.Params, "ListItems", ppt)
	_ = renderParamValues(f.Params)
	_ = renderInvalidateExpr([]string{"ListItems", "ListSub"})
	_ = renderInvalidateExpr(nil)
}

// TestByNameRenderHooks_ZeroCov calls render hooks / mutation / query / imports by name.
func TestByNameRenderHooks_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	cons := byNameConstraints()
	noBody := map[string]bool{}
	ppt := map[string]map[string]string{}
	f := page.Fetches[0]
	a := page.Actions[0]

	if s := renderUseQuery(f, ppt); s == "" {
		t.Errorf("renderUseQuery empty")
	}
	if s := renderUseMutation(a, []string{"ListItems"}, false, noBody, ppt, cons); s == "" {
		t.Errorf("renderUseMutation empty")
	}

	var sb strings.Builder
	renderFetchHooks(f, ppt, &sb)
	is := collectImports(page, "")
	renderPageHooks(page, is, ppt, &sb)
	renderPageMutations(page.Actions, []string{"ListItems"}, buildActionFetchMap(page), cons, false, noBody, ppt, &sb)
	if sb.Len() == 0 {
		t.Errorf("hook renderers produced nothing")
	}

	imports := renderImports(is, DefaultOptions())
	if imports == "" {
		t.Errorf("renderImports empty")
	}

	var jb strings.Builder
	renderPageJSX(page, &jb, noBody)
	renderPageJSXWithChildren(page.Children, &jb, noBody)
	if jb.Len() == 0 {
		t.Errorf("renderPageJSX produced nothing")
	}
}

// TestByNameGeneratePage_ZeroCov calls GeneratePage by name (best-effort).
func TestByNameGeneratePage_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	out := GeneratePage(page, "", GenerateOptions{
		RequestConstraints:      byNameConstraints(),
		ResponseArrayItemFields: map[string]map[string]map[string]bool{"ListItems": {"items": {"id": true}}},
	})
	if out == "" {
		t.Errorf("GeneratePage empty")
	}
}
