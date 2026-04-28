# pkg/validate/ssac_manifest

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC 의 currentUser / `@publish` / `@subscribe` / JWT `@call` 및 session/cache/file/queue 빌트인 호출이 `manifest.yaml` 의 claims·backend 설정과 일치하는지 확인 (XNS-*, XSA-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XNS-48 | `CurrentUserClaims` | currentUser 사용 → manifest claims 활성화 (ERROR) | IF-ELSE | ✓ |
| XNS-49 | `CurrentUserField` | `currentUser.<field>` → manifest claims (ERROR) | IF-ELSE | ✓ |
| XNS-56 | `QueueRequired` | `@publish`/`@subscribe` → `queue.backend` (ERROR) | IF-ELSE | ✓ |
| XNS-57 | `MemoryTxPublish` | `queue.backend=memory` + tx-bound `@publish` 조합 (WARNING) | IF-ELSE | ✓ |
| XNS-73 | `JwtCallClaims` | JWT `@call` input → manifest claims fields (ERROR) | IF-ELSE | ✓ |
| XSA-70 | `SessionBackendRequired` | `@call session.*` → `manifest.session.backend` (ERROR) | IF-ELSE | ✓ |
| XSA-71 | `CacheBackendRequired` | `@call cache.*` → `manifest.cache.backend` (ERROR) | IF-ELSE | ✓ |
| XSA-72 | `FileBackendRequired` | `@call file.*`/`storage.*` → `manifest.file.backend` (ERROR) | IF-ELSE | ✓ |
| XSA-74 | `SessionBackendUnused` | session.backend 선언 + `@call` 없음 (WARNING) | IF-ELSE | ✓ |
| XSA-75 | `CacheBackendUnused` | cache.backend 선언 + `@call` 없음 (WARNING) | IF-ELSE | ✓ |
| XSA-76 | `FileBackendUnused` | file.backend 선언 + `@call` 없음 (WARNING) | IF-ELSE | ✓ |
| XSA-77 | `QueueBackendUnused` | queue.backend 선언 + `@publish`/`@subscribe` 없음 (WARNING) | IF-ELSE | ✓ |

## Defeater

없음.
