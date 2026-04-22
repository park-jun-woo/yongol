# pkg/validate/ssac_manifest

SSaC 의 currentUser/@publish/@subscribe/JWT @call 이 `manifest.yaml` 의 claims/queue 설정을 올바로 참조하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XNS-49 | `Manifest.claims` | currentUser.field → claims | IF-ELSE |
| XNS-73 | `Manifest.claims.fields` | JWT @call input → claims | IF-ELSE |

## ConfigRequired (IF-ELSE)

| 규칙 ID | ConfigKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XNS-48 | `backend.auth.claims` | currentUser 사용 → claims 필수 | IF-ELSE |
| XNS-56 | `queue.backend` | @publish/@subscribe → queue 설정 필수 | IF-ELSE |

## Defeater

없음.
