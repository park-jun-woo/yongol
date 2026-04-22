//ff:func feature=tsx-parser type=command control=iteration dimension=1
//ff:what swc AST JSON 을 순회하며 apiClient 호출 / register() / 로컬 component import 추출
package tsx

import (
	"encoding/json"
	"strings"
)

// astNode is a permissive view of a swc AST node. Only `type` and `span`
// are always present; all other fields are consumed on demand per node kind.
type astNode struct {
	Type string          `json:"type"`
	Span astSpan         `json:"span"`
	Raw  json.RawMessage `json:"-"`
}

type astSpan struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// visitor accumulates extraction results and maintains the byte-offset to
// (line, col) index derived from the source text.
type visitor struct {
	src        []byte
	lineOffset []int // byte offset of the start of each 1-based line
	page       *PageSpec
}

// newVisitor builds a line index from src so swc `span.start` byte offsets
// can be translated into (line, col). swc span starts at 1, not 0, per the
// swc documentation — adjusted here to standard 0-based offsets first.
func newVisitor(src []byte, page *PageSpec) *visitor {
	// lineOffset[0] = 0 (line 1 starts at byte 0)
	off := []int{0}
	for i, b := range src {
		if b == '\n' {
			off = append(off, i+1)
		}
	}
	return &visitor{src: src, lineOffset: off, page: page}
}

// resolve converts a swc span.start (1-based byte offset) into 1-based
// (line, col). A start of 0 or out-of-range yields (0, 0) so callers can
// detect missing positions.
func (v *visitor) resolve(spanStart int) (line, col int) {
	if spanStart <= 0 {
		return 0, 0
	}
	// swc spans are 1-based; convert to 0-based byte offset.
	off := spanStart - 1
	if off > len(v.src) {
		off = len(v.src)
	}
	// Binary search over lineOffset.
	lo, hi := 0, len(v.lineOffset)-1
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if v.lineOffset[mid] <= off {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	line = lo + 1
	col = off - v.lineOffset[lo] + 1
	return
}

// snippet returns the raw source bytes for a swc span (best-effort; used
// only for diagnostic values like ArgBinding.Value).
func (v *visitor) snippet(span astSpan) string {
	if span.Start <= 0 || span.End <= span.Start {
		return ""
	}
	s := span.Start - 1
	e := span.End - 1
	if s < 0 {
		s = 0
	}
	if e > len(v.src) {
		e = len(v.src)
	}
	if s >= e {
		return ""
	}
	return strings.TrimSpace(string(v.src[s:e]))
}

// walkRoot decodes the root Module and visits its body.
func (v *visitor) walkRoot(root json.RawMessage) error {
	var m struct {
		Body []json.RawMessage `json:"body"`
	}
	if err := json.Unmarshal(root, &m); err != nil {
		return err
	}
	for _, node := range m.Body {
		v.walk(node)
	}
	return nil
}

// walk dispatches on node `type` and recurses into child containers. The
// walker is permissive: unknown node types are still recursed into using
// the generic value-based descent so newly-added swc node kinds never drop
// matches (e.g. inside decorators, JSX expressions).
func (v *visitor) walk(raw json.RawMessage) {
	if len(raw) == 0 || string(raw) == "null" {
		return
	}
	var head astNode
	if err := json.Unmarshal(raw, &head); err != nil {
		return
	}
	switch head.Type {
	case "ImportDeclaration":
		v.handleImport(raw)
	case "CallExpression":
		v.handleCall(raw)
	case "KeyValueProperty":
		v.handleKeyValueProperty(raw)
	}
	// Always descend — a CallExpression may host child CallExpressions (chained
	// then()s), and ImportDeclarations are leaves.
	v.descend(raw)
}

// descend recurses into every child value of an arbitrary JSON object or
// array, invoking walk on every encountered object. This is simpler than
// enumerating swc's ~120 node shapes and the overhead is negligible for
// per-page ASTs (typically a few hundred KB).
func (v *visitor) descend(raw json.RawMessage) {
	// Fast reject: only objects / arrays contain children.
	switch raw[0] {
	case '{':
		var obj map[string]json.RawMessage
		if err := json.Unmarshal(raw, &obj); err != nil {
			return
		}
		for k, child := range obj {
			if k == "type" || k == "span" {
				continue
			}
			v.walkAny(child)
		}
	case '[':
		var arr []json.RawMessage
		if err := json.Unmarshal(raw, &arr); err != nil {
			return
		}
		for _, child := range arr {
			v.walkAny(child)
		}
	}
}

// walkAny routes into walk for object-shaped AST nodes and into descend
// for arrays. Scalar leaves are skipped.
func (v *visitor) walkAny(raw json.RawMessage) {
	if len(raw) == 0 {
		return
	}
	switch raw[0] {
	case '{':
		v.walk(raw)
	case '[':
		v.descend(raw)
	}
}

// handleImport extracts ComponentImports from local paths only. npm package
// imports (react, @tanstack/*, etc) are intentionally skipped so T-1 only
// triggers on imports yongol is responsible for.
func (v *visitor) handleImport(raw json.RawMessage) {
	var d struct {
		Span       astSpan           `json:"span"`
		Specifiers []json.RawMessage `json:"specifiers"`
		Source     struct {
			Value string `json:"value"`
		} `json:"source"`
	}
	if err := json.Unmarshal(raw, &d); err != nil {
		return
	}
	if !isLocalComponentImport(d.Source.Value) {
		return
	}
	line, _ := v.resolve(d.Span.Start)
	for _, spec := range d.Specifiers {
		var s struct {
			Type  string `json:"type"`
			Local struct {
				Value string `json:"value"`
			} `json:"local"`
			Imported *struct {
				Value string `json:"value"`
			} `json:"imported"`
		}
		if err := json.Unmarshal(spec, &s); err != nil {
			continue
		}
		name := s.Local.Value
		if s.Imported != nil && s.Imported.Value != "" {
			name = s.Imported.Value
		}
		if name == "" {
			continue
		}
		v.page.Imports = append(v.page.Imports, ComponentImport{
			Name: name,
			Path: d.Source.Value,
			Line: line,
		})
	}
}

// isLocalComponentImport matches the two conventions emitted by yongol's
// React scaffold: path-alias `@/components/...` and relative `./components/...`
// (or `../components/...`). Anything else — npm packages, absolute HTTP URLs,
// deep path-aliases like `@/lib/...` — is skipped.
func isLocalComponentImport(src string) bool {
	if src == "" {
		return false
	}
	if strings.HasPrefix(src, "@/components/") {
		return true
	}
	if strings.HasPrefix(src, "./components/") || strings.HasPrefix(src, "../components/") {
		return true
	}
	// Relative sibling paths inside pages/ that still live under components/.
	if strings.HasPrefix(src, "./") || strings.HasPrefix(src, "../") {
		if strings.Contains(src, "/components/") {
			return true
		}
	}
	return false
}

// handleCall extracts apiClient.<op>(...) and register('name', ...) calls.
// A single CallExpression can match neither, one, or both (rare) patterns;
// each path returns independently.
func (v *visitor) handleCall(raw json.RawMessage) {
	var c struct {
		Span      astSpan           `json:"span"`
		Callee    json.RawMessage   `json:"callee"`
		Arguments []json.RawMessage `json:"arguments"`
	}
	if err := json.Unmarshal(raw, &c); err != nil {
		return
	}

	// Pattern 1: apiClient.<op>(...)
	if opID, ok := matchApiClientCallee(c.Callee); ok {
		line, col := v.resolve(calleePropertySpan(c.Callee).Start)
		call := APICall{OperationID: opID, Kind: "raw", Line: line, Col: col}
		if len(c.Arguments) > 0 {
			call.Args = v.extractArgBindings(c.Arguments[0])
		}
		v.page.Calls = append(v.page.Calls, call)
	}

	// Pattern 2: register('name', opts?)
	if name, required, ok := matchRegisterCall(c.Callee, c.Arguments); ok {
		line, col := v.resolve(c.Span.Start)
		v.page.FormFields = append(v.page.FormFields, FormField{
			Name: name, Required: required, Line: line, Col: col,
		})
	}
}

// matchApiClientCallee checks that the callee is MemberExpression with
// object=Identifier("apiClient") and property=Identifier(<op>). Returns
// the operationId and true on success. apiClient.foo.bar() is rejected.
func matchApiClientCallee(callee json.RawMessage) (string, bool) {
	var m struct {
		Type     string `json:"type"`
		Object   struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"object"`
		Property struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"property"`
	}
	if err := json.Unmarshal(callee, &m); err != nil {
		return "", false
	}
	if m.Type != "MemberExpression" {
		return "", false
	}
	if m.Object.Type != "Identifier" || m.Object.Value != "apiClient" {
		return "", false
	}
	if m.Property.Type != "Identifier" || m.Property.Value == "" {
		return "", false
	}
	return m.Property.Value, true
}

// calleePropertySpan returns the span of the property identifier in a
// MemberExpression callee. Falls back to a zero span on error so the
// caller's resolve(0) produces (0,0).
func calleePropertySpan(callee json.RawMessage) astSpan {
	var m struct {
		Property struct {
			Span astSpan `json:"span"`
		} `json:"property"`
	}
	_ = json.Unmarshal(callee, &m)
	return m.Property.Span
}

// matchRegisterCall recognizes `register('field'[, { required: true|false }])`.
// Works whether called standalone or as part of a {...register('x')} spread.
// required is best-effort from a BooleanLiteral inside the options object.
func matchRegisterCall(callee json.RawMessage, args []json.RawMessage) (string, bool, bool) {
	var ident struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(callee, &ident); err != nil {
		return "", false, false
	}
	if ident.Type != "Identifier" || ident.Value != "register" {
		return "", false, false
	}
	if len(args) == 0 {
		return "", false, false
	}
	var firstArg struct {
		Expression struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"expression"`
	}
	if err := json.Unmarshal(args[0], &firstArg); err != nil {
		return "", false, false
	}
	if firstArg.Expression.Type != "StringLiteral" || firstArg.Expression.Value == "" {
		return "", false, false
	}
	name := firstArg.Expression.Value
	required := false
	if len(args) >= 2 {
		required = parseRequiredFromOptions(args[1])
	}
	return name, required, true
}

// parseRequiredFromOptions extracts the boolean value of the `required`
// property from a register() options literal. Returns false on any shape
// mismatch; non-literal expressions (variables, ternaries) are treated as
// "undetermined" → false.
func parseRequiredFromOptions(arg json.RawMessage) bool {
	var w struct {
		Expression struct {
			Type       string            `json:"type"`
			Properties []json.RawMessage `json:"properties"`
		} `json:"expression"`
	}
	if err := json.Unmarshal(arg, &w); err != nil {
		return false
	}
	if w.Expression.Type != "ObjectExpression" {
		return false
	}
	for _, p := range w.Expression.Properties {
		var kv struct {
			Type string `json:"type"`
			Key  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"key"`
			Value struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"value"`
		}
		if err := json.Unmarshal(p, &kv); err != nil {
			continue
		}
		if kv.Type != "KeyValueProperty" {
			continue
		}
		if kv.Key.Value == "required" && kv.Value.Type == "BooleanLiteral" {
			return kv.Value.Value
		}
	}
	return false
}

// handleKeyValueProperty recognises `{ mutationFn: apiClient.X }` and
// `{ queryFn: apiClient.X }` patterns where the value is a bare member
// expression (not a CallExpression). Without this branch, TanStack Query
// style registrations slip past XOT-* rules because there is no literal
// `apiClient.X()` site for the visitor to hook onto.
func (v *visitor) handleKeyValueProperty(raw json.RawMessage) {
	var kv struct {
		Key struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"key"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(raw, &kv); err != nil {
		return
	}
	switch kv.Key.Value {
	case "mutationFn", "queryFn":
		// continue
	default:
		return
	}
	opID, ok := matchApiClientCallee(kv.Value)
	if !ok {
		return
	}
	line, col := v.resolve(calleePropertySpan(kv.Value).Start)
	// Already recorded by a concrete CallExpression? Skip to avoid doubles.
	for _, existing := range v.page.Calls {
		if existing.OperationID == opID && existing.Line == line {
			return
		}
	}
	v.page.Calls = append(v.page.Calls, APICall{
		OperationID: opID, Kind: "raw", Line: line, Col: col,
	})
}

// extractArgBindings pulls keys out of the first argument of an apiClient
// call. Only the top-level ObjectExpression's keys are captured — nested
// objects and spreads are skipped (XOT-2 compares flat parameter names only).
func (v *visitor) extractArgBindings(arg json.RawMessage) []ArgBinding {
	var w struct {
		Expression struct {
			Type       string            `json:"type"`
			Properties []json.RawMessage `json:"properties"`
		} `json:"expression"`
	}
	if err := json.Unmarshal(arg, &w); err != nil {
		return nil
	}
	if w.Expression.Type != "ObjectExpression" {
		return nil
	}
	out := make([]ArgBinding, 0, len(w.Expression.Properties))
	for _, p := range w.Expression.Properties {
		var kv struct {
			Type string `json:"type"`
			Key  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"key"`
			Value struct {
				Span astSpan `json:"span"`
			} `json:"value"`
		}
		if err := json.Unmarshal(p, &kv); err != nil {
			continue
		}
		// Shorthand: { id } → Type="Identifier" with same Value as key.
		if kv.Type == "Identifier" {
			var ident struct {
				Value string  `json:"value"`
				Span  astSpan `json:"span"`
			}
			if err := json.Unmarshal(p, &ident); err == nil && ident.Value != "" {
				out = append(out, ArgBinding{Key: ident.Value, Value: ident.Value})
			}
			continue
		}
		if kv.Type != "KeyValueProperty" || kv.Key.Value == "" {
			continue
		}
		out = append(out, ArgBinding{
			Key:   kv.Key.Value,
			Value: v.snippet(kv.Value.Span),
		})
	}
	return out
}
