//ff:type feature=tsx-parser type=model
//ff:what TSX 페이지 SSOT 에서 추출된 계약 구조체 (API 호출 / 폼 필드 / 컴포넌트 import)
package tsx

// PageSpec is the parsed contract extracted from a single React .tsx file.
// Represents the intersection between TSX (SSOT) and OpenAPI (SSOT) that
// downstream crosscheck rules (XOT-*) and self-consistency rules (T-*)
// operate against.
type PageSpec struct {
	File       string            // absolute or repo-relative path to the .tsx file
	Calls      []APICall         // apiClient.<op>(...) invocations
	FormFields []FormField       // useForm().register('x', ...) fields
	Imports    []ComponentImport // local component imports (@/components/..., ./components/...)
}

// APICall captures a single apiClient.<operationId>(...) invocation.
// Kind is a best-effort hint reserved for future heuristics (e.g. useQuery
// vs useMutation surrounding context). Phase001 always stores "raw".
type APICall struct {
	OperationID string       // "listWorkflows"
	Kind        string       // reserved: "query" | "mutation" | "raw"
	Args        []ArgBinding // keys of the first ObjectExpression argument
	Line        int          // 1-based source line of the callee
	Col         int          // 1-based source column of the callee
}

// ArgBinding is a single property key inside apiClient.<op>({ ... }).
// Value is the raw source snippet (best-effort) used purely for diagnostics;
// XOT-2 only compares Key names against OpenAPI parameter names.
type ArgBinding struct {
	Key   string
	Value string
}

// FormField captures a react-hook-form register('name', opts) invocation.
type FormField struct {
	Name     string
	Required bool
	Line     int
	Col      int
}

// ComponentImport captures a local component import (non-npm).
// Only imports whose source starts with "@/components/" or "./components/"
// are recorded; npm package imports are filtered out before emission.
type ComponentImport struct {
	Name string // imported symbol ("Button")
	Path string // module source ("@/components/ui/Button")
	Line int
}
