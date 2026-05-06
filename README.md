# thread-mob-vault-index

`thread-mob-vault-index` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for vault workflows, centered on format conversion, round-trip fixtures, and lossless normalization checks.

## Project Rationale

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Thread Mob Vault Index Review Notes

`stress` and `stale` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## Feature Set

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/thread-mob-vault-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `sync drift` and `form pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Architecture

The implementation keeps the scoring rule plain: reward signal and confidence, preserve slack, penalize drag, then classify the result into a review lane.

The Go implementation avoids hidden state so fixture changes are easy to reason about.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Test Command

The same command runs the local verification path. The highest-scoring domain case is `stress` at 228, which lands in `ship`. The most cautious case is `stale` at 126, which lands in `watch`.

## Next Improvements

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
