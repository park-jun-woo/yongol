# pkg/validate/hurl_manifest

## 변경이력

- 2026-04-28: 초기 작성

## 역할

Hurl 테스트 entry 의 호출 흐름이 `manifest.yaml` `backend.auth` 모드 (cookie / bearer / none) 와 정합한지 교차 검증한다. 보호 구간 호출 전 인증 선행 여부 (XOH-06) 와 cookie 모드의 mutating 요청 CSRF 헤더 — manifest 해석값 `backend.auth.csrf.header_name`, 기본 `X-XSRF-TOKEN` — (XOH-07) 를 점검한다.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `IF-ELSE` = entry 순회 + auth state machine 휴리스틱

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XOH-06 | `xoh_06_auth_precondition` | 보호 구간 operation 호출 전에 인증 스텝이 선행되어야 함 (WARNING) | IF-ELSE | ✓ |
| XOH-07 | `xoh_07_csrf_on_mutation` | cookie 인증 모드에서 mutating 요청은 manifest 해석 CSRF 헤더 (기본 `X-XSRF-TOKEN`) 포함 (WARNING) | IF-ELSE | ✓ |

## 주요 함수

| 함수 | 설명 |
|---|---|
| `Run(fs)` | Hurl ↔ Manifest 교차 검증 실행 (XOH-06/07) |

## 보조 헬퍼

`groupByFile`, `processAuthEntry`, `checkFileAuth`, `carriesAuthHeader`, `hasCSRFHeader`, `isAuthPath`, `isMutating`, `is2xx`, `shouldCheckCSRF` 가 entry 단위 판정·auth state 갱신을 담당.
