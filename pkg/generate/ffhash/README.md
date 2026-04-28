# pkg/generate/ffhash

## 변경이력

- 2026-04-28: 초기 작성

## 역할

생성된 Go 산출물의 `//ff:` 어노테이션 블록 끝에 `//ff:checked llm=yongol-gen hash=<sha>` 라인을 삽입/갱신한다. 사용자 편집 보존(preserve) 검사용 baseline hash 를 코드 자체에 박아 두기 위한 후처리.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `WalkAndInject` | `(root string, skipRelPrefixes []string) error` | 디렉토리를 재귀 순회하며 `.go` 파일에 checked 라인 삽입 (skip prefix 제외) |
| `InjectCheckedLine` | `(src []byte) []byte` | 단일 소스 바이트에 checked 라인 삽입/갱신 (idempotent) |

## 동작

- `//ff:func` 또는 `//ff:type` 블록이 없는 파일은 건드리지 않음
- `func` 선언이 없는 type-only 파일도 스킵 (hash 계산 불가)
- 동일 입력 두 번 호출해도 출력 동일 (idempotency)
- 변동이 있는 파일만 다시 기록
