# Review Journal

This journal records the domain cases that matter before widening the public API.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 201, lane `ship`
- `stress`: `sync drift`, score 228, lane `ship`
- `edge`: `local state`, score 178, lane `ship`
- `recovery`: `conflict cost`, score 156, lane `ship`
- `stale`: `form pressure`, score 126, lane `watch`

## Note

A future change should add new cases before it changes the scoring rule.
