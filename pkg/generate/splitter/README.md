# pkg/generate/splitter

## 변경이력

- 2026-04-28: 초기 작성

## 역할

외부 코드젠 도구 (oapi-codegen / sqlc) 의 단일 거대 산출물 (`server.gen.go`, `models.go`, `*.sql.go`) 을 AST 로 읽어 선언 단위로 분할한다. 분할 파일에 filefunc 어노테이션 (`//ff:func` / `//ff:type` / `//ff:what`) 자동 주입 후 원본 삭제.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `SplitDirectory` | `(dir string, tool Tool) error` | 디렉토리의 도구 산출 파일을 모두 분할 + 원본 정리 |
| `SplitFile` | `(srcPath, outDir string, tool Tool) ([]string, error)` | 단일 파일 분할 + 어노테이션 주입. 결과 파일 경로 반환 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Tool` | splitter 가 인식하는 도구 (oapi-codegen / sqlc) |

## 분할 규칙

- FuncDecl: `func_name.go` (method 는 `receiver_method.go`)
- GenDecl(type/const/var): 단일 spec 이름 → `snake_case.go`
- import 는 선언이 실제 사용하는 식별자만 살려 재구성 (blank/dot import 유지)
- 원본 generator 배너 (`Code generated ... DO NOT EDIT.`) 는 모든 분할 파일 선두에 복제
- doc comment 첫 줄 → `//ff:what` 요약, FuncDecl body depth-1 제어구조 → `control=` 자동 추론
