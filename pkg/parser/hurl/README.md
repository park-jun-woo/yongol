# pkg/parser/hurl

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`.hurl` 통합 테스트 시나리오 파일 1건에서 요청/응답 쌍을 추출한다. Method / Path / StatusCode 외에도 JSON 요청 body 의 top-level 필드명, 요청 헤더, `[Captures]`, `[Asserts]` jsonpath 항목을 함께 수집해 XOH-01~09 교차 검증에 공급한다. 디렉토리 순회는 호출측 책임 (파일 단위 API 만 제공).

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParseFile` | `(path string) ([]HurlEntry, []diagnostic.Diagnostic)` | 단일 `.hurl` 파일 파싱 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `HurlEntry` | Hurl 요청/응답 쌍 1건 (`Method / Path / StatusCode / BodyFields / Headers / Captures / Asserts` 등) |
| `HurlHeader` | 요청 헤더 선언 |
| `HurlCapture` | `[Captures]` 라인 (`var := jsonpath\|header\|expr`) |
| `HurlAssert` | `[Asserts]` 의 `jsonpath "$.path" ...` 어서션 라인 |

## 비고

- `{{var}}` placeholder 는 JSON decode 가 가능한 형태로 정규화 후 body 필드명을 추출.
- 디렉토리 수집 헬퍼 `collectHurlFiles` 는 내부용. 외부에서는 `filepath.Walk` 등으로 직접 순회.
