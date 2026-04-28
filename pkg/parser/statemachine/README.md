# pkg/parser/statemachine

## 변경이력

- 2026-04-28: 초기 작성

## 역할

Mermaid `stateDiagram` 코드 블록이 들어있는 마크다운 파일 (`states/*.md`) 을 파싱해 초기 상태 / 전이 / 이벤트 / 유효 출발 상태 목록을 추출한다. 파일명에서 다이어그램 ID 를 도출하고, ID 의 PascalCase 변환을 `Symbol` 로 보관해 다운스트림 코드젠이 일관된 Go 식별자를 사용할 수 있도록 한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParseDir` | `(dir string) ([]*StateDiagram, []diagnostic.Diagnostic)` | 디렉토리 내 `*.md` 파일 일괄 파싱 (디렉토리 부재는 silent OK) |
| `ParseFile` | `(path string) (*StateDiagram, []diagnostic.Diagnostic)` | 단일 마크다운 파일 파싱, ID 는 파일명에서 추출 |
| `Parse` | `(id, content, file string) (*StateDiagram, []diagnostic.Diagnostic)` | mermaid 텍스트와 ID 를 받아 직접 파싱 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `StateDiagram` | mermaid stateDiagram 결과 (`ID / Symbol / File / InitialState / States / Transitions`) + 메서드 `Events()`, `ValidFromStates(event)` |
| `Transition` | 상태 전이 (`From / To / Event / Line`) |

## 비고

- 마크다운 ` ```mermaid ` 블록 한 개만 인식 (다중 블록은 미지원).
- 동일 ID 내 대소문자만 다른 상태명 (`draft` vs `Draft`) 은 ERROR 로 보고.
