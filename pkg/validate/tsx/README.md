# pkg/validate/tsx

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

React `.tsx` SSOT 자체 정합성 검증 (T-*). 교차 검증은 `pkg/validate/tsx_openapi/` (XOT-*) 에 있다.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·경로 해석 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| T-1 | `ComponentFile` | `@/components/` / `./components/` import 의 실제 파일 (.tsx/.ts/.jsx/.js/index) 존재 (WARNING) | IF-ELSE | ✓ |

## Defeater

없음. npm 패키지 import 는 파서 단계에서 이미 제외.
