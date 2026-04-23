//ff:type feature=tsx-parser type=model
//ff:what PageSpec — 단일 .tsx 파일에서 추출한 API 호출 / 폼 필드 / 컴포넌트 import 계약

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
