# pkg/contract

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`//ff:checked hash=<hex>` preserve 해시 / `//ff:preserve reason="..."` 어노테이션 / 외부 심볼 추출을 담당하는 인프라. 코드젠 산출물 위에 사용자가 작성한 본문을 바디 해시로 식별하고, 정합성 검증(PRV-* 룰)이 참조할 시그니처·외부 심볼을 추출한다.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ComputeBodyHash` | `ComputeBodyHash(src []byte) string` | 첫 non-init FuncDecl 본문의 SHA-256 앞 4바이트 (8 hex). filefunc A7 호환. CRLF/BOM 정규화 후 계산 |
| `NormalizeBody` | `NormalizeBody(src []byte) []byte` | BOM 제거, CRLF/CR → LF, 마지막 newline 보장 |
| `StripAnnotationBlock` | `StripAnnotationBlock(src []byte) []byte` | 파일 상단 `//ff:` 어노테이션 블록 + generator 배너 제거 후 본문 반환 |
| `DetectPreserved` | `DetectPreserved(filePath string) (PreservedState, error)` | `//ff:checked` 의 saved hash 와 재계산 hash 비교로 preserve 상태 판정 |
| `CollectPreserved` | `CollectPreserved(rootDir string) ([]string, error)` | 디렉토리를 walk 하며 `StatePreserved` 인 `.go` 파일 경로 수집 (`.git`, `vendor`, `node_modules` 스킵) |
| `ParsePreserveReason` | `ParsePreserveReason(filePath string) (string, error)` | `//ff:preserve reason="..."` 어노테이션의 reason 값 추출 |
| `ExtractSignature` | `ExtractSignature(filePath string) (FuncSignature, error)` | 첫 non-init FuncDecl 의 시그니처 (이름·파라미터·반환·error 여부) 추출 |
| `ExtractExternalSymbols` | `ExtractExternalSymbols(filePath string) (ExternalSymbols, error)` | 첫 func body 를 AST walk 해 sqlc 쿼리 / `<pkg>.<Func>` 호출 / struct field 접근을 분류 수집 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `PreservedState` | `StateNotApplicable` / `StateUntouched` / `StatePreserved` 3-상태 enum |
| `FuncSignature` | `Name`, `Params []FuncParam`, `Returns []string`, `HasErr bool` |
| `FuncParam` | `Name`, `Type` (다중 이름 그룹은 분해, 익명 파라미터는 `Name == ""`) |
| `ExternalSymbols` | `SqlcQueries` / `CallTargets` / `DDLFields` 세 버킷 |
