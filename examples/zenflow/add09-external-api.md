# ZenFlow Add-on #09 — External API Integration

## Overview

Call an external API from SSaC using a `yongol import`-generated client. Exercises the external API integration pattern.

## Setup

```bash
yongol import specs/external/geocoding-api.yaml func/geocoding/
```

This generates a Go client package in `func/geocoding/` callable from SSaC via `@call geocoding.<Func>(...)`.

## New Endpoint

### POST /organizations/{id}/verify-address (`VerifyOrgAddress`)

Verify and geocode the organization's address using an external geocoding API.

1. `@auth` — admin only, same org.
2. `@get` the organization.
3. `@call geocoding.Geocode` with the org's address.
4. `@put` update org with verified lat/lng.
5. `@response` the updated organization.

## Key Pattern

```
1. yongol import <external-openapi.yaml> func/<pkg>/
2. SSaC: @call <pkg>.<Func>({...}) — same as any other @call
3. No net/http in func — the generated client handles HTTP internally
```

## DDL Change

Add `latitude NUMERIC`, `longitude NUMERIC`, `address_verified BOOLEAN DEFAULT false` columns to `organizations`.

## E2E Scenario

- Call VerifyOrgAddress, verify lat/lng are populated and address_verified is true.
