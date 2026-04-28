# pkg/parser/ssac

## 변경이력

- 2026-04-28: 초기 작성

## 역할

SSaC (`*.ssac`) 파일을 Go AST 로 파싱해 서비스 함수 시퀀스를 추출한다. 함수 선행 주석에서 `@get / @post / @put / @delete / @call / @auth / @state / @eval / @empty / @exists / @response / @subscribe / @publish / @verify-password` 등 어노테이션을 시퀀스로 변환하고, struct 선언과 import 도 함께 수집한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParseDir` | `(dir string) ([]ServiceFunc, []diagnostic.Diagnostic)` | 디렉토리 재귀 순회 + feature 폴더명 자동 부착 |
| `ParseFile` | `(path string) ([]ServiceFunc, []diagnostic.Diagnostic)` | 단일 `.ssac` 파일 파싱 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `ServiceFunc` | 서비스 함수 1건 (`Name / Feature / Sequences / Imports / Structs / Subscribe / Params / SuppressWarn / Line / File` 등) |
| `Sequence` | 단일 SSaC 시퀀스 라인 (annotation 종류별로 다양한 필드: Type / Model / Method / Result / Inputs / Target / Message / DiagramID / Transition / Topic / Options / ErrStatus / Line) |
| `Arg` | 함수 호출 인자 1건 |
| `Result` | 결과 바인딩 (`Type / Var / Wrapper`, e.g. `Cursor[T]`, `Page[T]`) |
| `StructInfo`, `StructField` | `.ssac` 파일에 선언된 Go struct |
| `SubscribeInfo` | 큐 구독 트리거 메타 (Topic / MessageType) |
| `ParamInfo` | 함수 파라미터 정보 |

## 비고

- `seq_const.go` 의 sequence 타입 상수 (`SeqGet / SeqPost / SeqAuth ...`) 는 후속 검증 / 코드젠에서 분기 키로 사용.
- 표준 라이브러리 `net/http` 등은 import 수집에서 제외 (`test_parse_imports_exclude_net_http_test`).
