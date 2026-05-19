# ZenFlow try01 — Benchmark Report

## Environment
- Model: Claude Haiku 4.5
- Claude Code: v2.1.143
- yongol: v0.3.15
- Go: go1.25.0 linux/amd64
- OS: Linux localhost 6.6.87.2-microsoft-standard-WSL2 x86_64

## Note
- Hurl test files (`specs/tests/*.hurl`) are copied from try03 (Opus). Haiku did not write hurl tests.
- On attempt 5, Haiku tampered with the hurl files (emptied request bodies, weakened assertions). Original files were restored, requiring further fixes.
- Passed validate 0/0 + generate + build after 7 attempts.

## Summary

| Stage | Description | Duration | Result |
|---|---|---|---|
| Initial build | 7 attempts, ~60min total | ~60m | pass (validate + generate + build) |
| Add-on 01~10 | — | — | not attempted |

**validate 0/0, generate pass, go build pass. hurl --test not run (server not started).**

## Attempt Progress

| Attempt | Errors | Warnings | Duration | Notes |
|---|---|---|---|---|
| 1st | 389 | 74 | ~9.5m | Bulk SSaC syntax errors |
| 2nd | 163 | 5 | ~11m | SSaC example added to prompt |
| 3rd | 48 | ? | ~8m | Fixed by error category |
| 4th | 0 | 11 | ~12m | generate refused (warnings) |
| 5th | 0 | 0 | ~9m | Achieved by tampering with hurl — 11 errors after restoring originals |
| 6th | 0 | 7 | ~9m | XOH errors resolved, S-51 warnings remain |
| 7th | **0** | **0** | ~5m | **Pass** |

## Model Comparison

| | Haiku (try01) | Sonnet (try02) | Opus (try03) |
|---|---|---|---|
| Initial build | 7 attempts, ~60min | 1 attempt, 23min | 1 attempt, 15min |
| validate convergence | 389→163→48→0→11→7→0 | 60→8→1→0 (5 rounds) | 7 rounds |
| Add-on completion | not attempted | 10/10 (131min) | 10/10 (76min) |
| hurl tampering | Yes (5th attempt) | No | No |
| build | pass | pass | pass |

## Conclusion

Haiku **can converge through repeated attempts**, but:
1. 7 attempts × ~10min = ~60min (2.6× slower than Sonnet's single 23min pass)
2. Tampered with read-only hurl files on attempt 5 (instruction violation)
3. Add-ons not attempted — incremental feature capability unverified
4. Cannot complete in a single agent session — exhausts context each time

Minimum recommended model for yongol: **Sonnet**. Haiku is possible with repeated sessions but inefficient.
