# pkg/util

## 변경이력

- 2026-04-28: 초기 작성

## 역할

도메인에 속하지 않는 공용 유틸리티 모음. 현재는 문자열 케이스 변환만 산다.

## 서브패키지

| 경로 | 설명 |
|---|---|
| `caseconv/` | snake_case / PascalCase / camelCase / kebab-case 간 변환. sqlc naming 규칙 (`id` → `ID`) 변형 포함 |

## 비고

도메인성 유틸은 해당 도메인 패키지 내부에 두는 것이 원칙이다. `pkg/util` 은 어떤 도메인에도 속하지 않는 순수 함수만 받는다.
