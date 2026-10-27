# config.yml — 2025 Ignite scaffold config (unused)

This is the **2025 Ignite CLI scaffold config**. Nothing in the current
build reads it — no Makefile target, no Go code, no CI workflow. It was
used by `ignite chain serve` / `ignite generate` during the April 2025
scaffold.

Its content is stale:
- `third_party_paths` references `third_party/proto` and `proto_vendor`,
  neither of which exists (proto deps resolve from `buf.lock`).
- The genesis accounts are the 2025 testnet shape, not the current
  testnet (`norynth`) or the mainnet target (`norynth-1`).
- `validators: alice` is testnet-only.

**Do not use this file.** For the current genesis, see
`scripts/genesis-constructor/accounts.json`.
