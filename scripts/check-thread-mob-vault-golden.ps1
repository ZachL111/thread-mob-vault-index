$ErrorActionPreference = "Stop"
$root = Resolve-Path (Join-Path $PSScriptRoot "..")
Set-Location $root

$profile = Get-Content -Raw metadata/property-map.json | ConvertFrom-Json
$w = $profile.weights
$thresholds = $profile.thresholds

function Get-Score($row) {
  return ([int]$row.signal * [int]$w.signal) + ([int]$row.slack * [int]$w.slack) + ([int]$row.confidence * [int]$w.confidence) - ([int]$row.drag * [int]$w.drag)
}

function Get-Lane($score) {
  if ($score -ge [int]$thresholds.ship) { return "ship" }
  if ($score -ge [int]$thresholds.watch) { return "watch" }
  return "hold"
}

$matrix = Import-Csv fixtures/golden/lane-matrix.csv
if ($matrix.Count -ne 27) { throw "thread-mob-vault-index lane matrix count changed" }
foreach ($row in $matrix) {
  $score = Get-Score $row
  if ($score -ne [int]$row.score) { throw "thread-mob-vault-index matrix score mismatch for $($row.id)" }
  if ((Get-Lane $score) -ne $row.lane) { throw "thread-mob-vault-index matrix lane mismatch for $($row.id)" }
}

$boundary = Get-Content fixtures/fuzz/boundary-cases.jsonl | ForEach-Object { $_ | ConvertFrom-Json }
if (@($boundary).Count -lt 20) { throw "thread-mob-vault-index boundary fixture too small" }
foreach ($row in $boundary) {
  $score = Get-Score $row
  if ($score -ne [int]$row.score) { throw "thread-mob-vault-index boundary score mismatch for $($row.id)" }
  if ((Get-Lane $score) -ne $row.lane) { throw "thread-mob-vault-index boundary lane mismatch for $($row.id)" }
}

Write-Host "thread-mob-vault-index golden checks passed"
