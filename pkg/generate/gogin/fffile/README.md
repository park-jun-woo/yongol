# pkg/generate/gogin/fffile

## 변경이력

- 2026-04-28: 초기 작성

## 역할

"1 파일 1 함수" 네이밍 규칙에 따른 파일명 계산 + preserve 가드 적용 후 디스크 기록 유틸. PascalCase 함수명을 snake_case `.go` 로 변환하고 receiver 가 있으면 `<receiver>_<method>.go` 형태로 결합한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `FileNameForFunc` | `(funcName string) string` | PascalCase → `snake_case.go` |
| `FileNameForMethod` | `(receiverType, methodName string) string` | `<receiver>_<method>.go` |
| `EnsureUnique` | `(candidate string, used map[string]bool) string` | 충돌 시 `_2` / `_3` suffix 부여 |
| `WriteIfNotPreserved` | `(path string, content []byte) error` | preserve 상태이면 skip, 아니면 `os.WriteFile` |
