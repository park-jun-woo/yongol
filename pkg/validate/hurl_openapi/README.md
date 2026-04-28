# pkg/validate/hurl_openapi

## 변경이력

- 2026-04-28: 초기 작성

## 역할

Hurl 테스트 entry 가 OpenAPI 명세와 정합한지 교차 검증한다. URL/method/status/request body 필드/assert·capture jsonpath/사용되지 않은 capture 까지 6 개 규칙으로 점검한다. (구 XOH-35/36/37 은 2026-04-24 XOH-01/02 로 통합.)

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `IF-ELSE` = OpenAPI route/schema lookup + 세그먼트 비교

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XOH-01 | `xoh_01_url_method` | Hurl URL + method 가 OpenAPI operation 으로 선언됨 | IF-ELSE | ✓ |
| XOH-02 | `xoh_02_status_declared` | Hurl `HTTP <status>` 가 OpenAPI responses 에 선언됨 | IF-ELSE | ✓ |
| XOH-03 | `xoh_03_request_field_in_schema` | Hurl JSON body 필드가 OpenAPI request schema 에 존재 | IF-ELSE | ✓ |
| XOH-04 | `xoh_04_assert_path_in_schema` | Hurl `[Asserts]` jsonpath 가 OpenAPI response schema 에서 도달 가능 | IF-ELSE | ✓ |
| XOH-08 | `xoh_08_capture_path_in_schema` | Hurl `[Captures]` jsonpath 가 OpenAPI response schema 에서 도달 가능 | IF-ELSE | ✓ |
| XOH-09 | `xoh_09_unused_capture` | capture 한 변수가 같은 파일에서 참조됨 (WARNING) | IF-ELSE | ✓ |

## 주요 함수

| 함수 | 설명 |
|---|---|
| `Run(fs)` | Hurl ↔ OpenAPI 교차 검증 실행 (XOH-01/02/03/04/08/09) |

## 보조 헬퍼

OpenAPI route 수집/정규화 (`collectOpenAPIRoutes`, `normalizeOpenAPIPath`, `normalizeHurlPath`, `findExactRoute`, `findPathMatch`), schema 도달성 (`jsonPathReachable`, `descend`, `descendAllOf`, `schemaPropertyNames`, `requestBodyProps`, `responseSchemaForStatus`), jsonpath 파서 (`parseJSONPath`, `parseJSONPathChar`, `flushCurJSONPath`).
