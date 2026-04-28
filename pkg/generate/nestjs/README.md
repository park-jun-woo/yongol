# pkg/generate/nestjs

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

NestJS (TypeScript + Node.js) backend target 자리표시 패키지. **계획 단계, 미구현.** 첫 출시 backend target 은 `pkg/generate/gogin/`. 본 디렉토리는 yongol 이 framework-agnostic 임을 표시.

> 상위: [`pkg/generate/README.md`](../README.md).

## 공개 함수

(미구현) `manifest.yaml` 의 `backend.lang: typescript` / `backend.framework: nestjs` 가 들어오면 generate 파이프라인이 "not yet supported — track pkg/generate/nestjs/README.md" 를 명시적으로 반환.

## SSOT → NestJS 매핑 (계획)

| SSOT | NestJS 산출물 |
|---|---|
| OpenAPI | `*.controller.ts` (`@Controller`, `@Get/@Post/...`, `@ApiOperation`) |
| DDL + sqlc | Prisma schema + client (후보; TypeORM/Drizzle 평가 중) |
| SSaC | controller method body + `*.service.ts` + `class-validator` DTO |
| `@auth` / OPA Rego | `@UseGuards(AuthzGuard)` + OPA client |
| Mermaid stateDiagram | service 내부 state-transition guard |
| Func spec (`@func`) | `@Injectable()` service class |
| Model (`@dto`) | `class-validator` / `class-transformer` DTO |
| manifest.yaml | `AppModule` providers + `main.ts` bootstrap |
| Hurl | 변경 없음 — HTTP surface 동일 |

## Parity 계약

Go+Gin target 과 wire-byte 동등: 동일 JSON 필드/타입/HTTP status/error envelope (`{error, code, request_id}`)/cursor·offset pagination/`//yg:checked` preserve (TS 어노테이션 형식만 다름). 동일 rulebook (PRV-10~17 은 TS idiom 으로 재표현).

## 비범위 (v1)

Express/Koa/Fastify standalone, GraphQL, ORM 다양성 (v1 은 단일 ORM), custom module layout (NestJS convention 강제: `controllers/`, `services/`, `dto/`, `guards/`, `modules/`).

## Roadmap

| Phase | 내용 |
|---|---|
| 1 | Design RFC (current) — ORM / preserve format / DI 경계 결정 |
| 2 | Scaffolding — 빈 컴파일 가능 NestJS 프로젝트 emit (`arts/backend-nestjs/`) |
| 3 | OpenAPI → controllers + DTOs |
| 4 | SSaC → controller + service body |
| 5 | Authz + state machines |
| 6 | Preserve contract (`// yg:checked llm=... hash=...`) |
| 7 | Go+Gin parity — 동일 Hurl suite 통과 |

## 설계 원칙

1. SSOT 가 권위. 9 SSOT 에서 도출 안 되는 구조 추가 금지. 2. Behavior parity > idiom (TS convention 보다 Go+Gin 동작 우선). 3. Preserve = 동일 의미, 다른 구문. 4. Test surface 공유 (1 Hurl suite, 2 backend build).

## 참고

[`pkg/generate/gogin/`](../gogin) (shipping target), `docs/ssac.md`, `rulebook.md`, `manual-for-ai.md`.
