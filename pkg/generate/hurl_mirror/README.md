# pkg/generate/hurl_mirror

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`specs/tests/**/*.hurl` 을 `arts/tests/` 로 1:1 디렉토리 미러링한다. 사용자 작성 Hurl 시나리오 파일을 내용 변경 없이 그대로 복사하고, 이전 실행 산출물 중 specs 에 더 이상 없는 .hurl (orphan) 은 자동 삭제.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `MirrorSpecsTests` | `(specsDir, artsDir string) (int, error)` | specs/tests/ → arts/tests/ 전체 복사. 복사된 파일 수 반환 |

## 동작

- `specsDir == ""` 또는 `specs/tests/` 부재 → no-op (0, nil)
- 하위 디렉토리 구조 보존
- 파일 내용 verbatim (변환 없음)
- arts/tests/ 에만 있는 orphan .hurl 은 prune
