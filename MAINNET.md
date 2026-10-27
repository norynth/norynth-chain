# Testnet vs Mainnet — separation

**Last updated:** 2026-10-27

Norynth has two chains: a **testnet** (running) and a **mainnet** (target,
not yet running). They are separate chains with separate accounts, keys, and
state. This document records the separation so it is never ambiguous.

---

## Testnet — `norynth` (running)

| | |
|---|---|
| **Chain ID** | `norynth` |
| **Host** | `norynth-vm` (`158.180.40.102`), Oracle `VM.Standard.E2.1.Micro` |
| **Home** | the node's home on the VM |
| **Denom** | `ukrn` (1 KRN = 1,000,000 ukrn) |
| **Genesis accounts** | the testnet genesis accounts (see the testnet genesis) |
| **Supply** | 1,051,000,000 KRN (1B + 1M faucet + 50M validator) |
| **Keyring** | the testnet keyring (on the VM) |
| **Purpose** | Prove the settlement path: wasm, contracts, upgrade |
| **Status** | Live, wasm-capable since 2026-10-26. Contract store/instantiate/execute proven. |

**The testnet genesis is not a template for mainnet.** It includes testnet-only
accounts (a faucet and a validator allocation), and its `contributors` bucket is
a single account (mainnet splits it into four).

---

## Mainnet — `norynth-1` (target, not yet running)

| | |
|---|---|
| **Chain ID** | `norynth-1` |
| **Host** | Not yet (needs x86_64 VPS, ≥16 GB RAM — see `MAINNET-READINESS.md` §1) |
| **Home** | a dedicated mainnet home (on the VPS, later) |
| **Denom** | `ukrn` (same unit, different supply/accounts) |
| **Supply** | 1,000,000,000 KRN (fixed, mint disabled) |
| **Genesis accounts** | 10: `ecosystem`, `builders`, `treasury`, `liquidity`, `community`, `strategic`, `founder`, `devs`, `advisors`, `future-hires` — **no faucet, no validator allocation** |
| **Vesting** | `ContinuousVestingAccount` for founder (4y/1y cliff), devs (3y/6m), advisors (2y/6m) |
| **Keyring** | a dedicated mainnet keyring on a separate machine (not the testnet host); mnemonics stored offline, encrypted |
| **Genesis construction** | `scripts/genesis-constructor/` (Go program + `accounts.json`) |
| **Purpose** | Settlement with real value |

**The mainnet genesis is built, not running.** The constructor
(`scripts/genesis-constructor/main.go`) reads `accounts.json` and a fresh
`genesis.json` (from `norynthd init --chain-id norynth-1`) and writes the
10 accounts. It is tested: `norynthd genesis validate` passes.

---

## The separation, explicitly

| Aspect | Testnet | Mainnet | Shared? |
|---|---|---|---|
| **Chain ID** | `norynth` | `norynth-1` | ❌ Different |
| **Accounts** | testnet genesis accounts | 10 mainnet accounts | ❌ Different |
| **Addresses** | testnet keys | fresh mainnet keys | ❌ Different |
| **Keyring** | on the testnet host | on a separate machine | ❌ Different |
| **Home** | on the testnet host | on the mainnet host (later) | ❌ Different |
| **Denom** | `ukrn` | `ukrn` | ✅ Same (same unit) |
| **Binary** | `norynthd` (wasm) | `norynthd` (wasm) | ✅ Same code |
| **Genesis** | the testnet genesis | `scripts/genesis-constructor/` output | ❌ Different |
| **Purpose** | Prove the path | Settlement with real value | ❌ Different |

**Same binary, different chains.** The binary (`norynthd`) is identical — it
runs both. What separates them is the **chain ID**, the **genesis**, the
**accounts**, and the **home directory**.

---

## What this means in practice

- **Do not reuse testnet keys for mainnet.** The mainnet keys are generated
  fresh, separate from the testnet keys.
- **Do not reuse the testnet genesis for mainnet.** The mainnet genesis is
  built by the constructor, with 10 accounts and `ContinuousVestingAccount`.
- **Do not run mainnet on the testnet VM.** The testnet VM is a 1 GB micro,
  which cannot run wasm mainnet. Mainnet needs x86_64, ≥16 GB (see
  `MAINNET-READINESS.md` §1).
- **The same `norynthd` binary runs both.** The difference is `--home` and
  `--chain-id`.

---

## Cross-references

**Public docs (live):**
- Chain docs: https://chain.norynth.com/docs/
- Tokenomics: https://chain.norynth.com/tokenomics
- Cloud / infrastructure: https://cloud.norynth.com/

**In-repo:**
- Genesis constructor: [`scripts/genesis-constructor/`](./scripts/genesis-constructor/)
- Developer portal: [`docs/DEVELOPERS.md`](./docs/DEVELOPERS.md)
- Historical (2025 Ignite): [`docs/_historical-2025/`](./docs/_historical-2025/)

**Internal (not in the repo):**
- Tokenomics design, architecture, readiness, and the repo audit are
  maintained internally and are not mirrored to this repository.
