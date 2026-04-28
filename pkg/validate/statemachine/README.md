# pkg/validate/statemachine

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Mermaid stateDiagram 자체 정합성 검증 (ST-*). 파싱 단계에서 구조 오류만 검출.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| ST-1 | `StateDiagramParse` | states/*.md 재파싱으로 구조 오류 감지 (ERROR) | IF-ELSE | ✓ |

## Defeater

없음. diagram 이 이미 로드되어 있으면 규칙 침묵 — `test_st01_parse_skips_when_loaded_test.go`.
