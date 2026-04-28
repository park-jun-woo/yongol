# pkg/generate/gogin/ffannot

## 변경이력

- 2026-04-28: 초기 작성

## 역할

gogin 산출 Go 파일의 package 선언 직전에 들어갈 `//ff:func` / `//ff:type` / `//ff:what` 어노테이션 블록 문자열을 조립한다. 함수 body 라인을 분석해 `control=sequence|selection|iteration` 도 자동 추론.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `EmitAnnotationBlock` | `(b Block) string` | `//ff:func` 또는 `//ff:type` + `//ff:what` 블록 렌더링 |
| `BuildFuncAnnot` | `(a FuncAnnot) string` | `//ff:func feature=... type=... control=...` 1 줄 조립 |
| `BuildTypeAnnot` | `(a TypeAnnot) string` | `//ff:type feature=... type=...` 1 줄 조립 |
| `BuildWhat` | `(desc string) string` | `//ff:what <desc>` 1 줄 조립 |
| `DetectControl` | `(bodyLines []string) string` | depth-0 라인에서 지배적 제어구조 추론 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Block` | EmitAnnotationBlock 입력 (FuncAnnot / TypeAnnot / What) |
| `FuncAnnot` | feature / type / control / dimension / topic |
| `TypeAnnot` | feature / type |
