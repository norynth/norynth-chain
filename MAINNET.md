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
| **Home** | `~/.norynth/` on the VM |
| **Denom** | `ukrn` (1 KRN = 1,000,000 ukrn) |
| **Genesis accounts** | `faucet`, `alice` (validator, 50M staked), `ecosystem`, `contributors` (single 200M), `builders`, `treasury`, `liquidity`, `community`, `strategic` |
| **Supply** | 1,051,000,000 KRN (1B + 1M faucet + 50M alice) |
| **Keyring** | `~/.norynth/keyring-test/` on the VM |
| **Purpose** | Prove the settlement path: wasm, contracts, upgrade |
| **Status** | Live, wasm-capable since 2026-10-26. Contract store/instantiate/execute proven. |

**The testnet genesis is not a template for mainnet.** It has `faucet` and
`alice` (testnet-only), and `contributors` is a single account (mainnet
splits it into four).

---

## Mainnet — `norynth-1` (target, not yet running)

| | |
|---|---|
| **Chain ID** | `norynth-1` |
| **Host** | Not yet (needs x86_64 VPS, ≥16 GB RAM — see `MAINNET-READINESS.md` §1) |
| **Home** | `~/.norynth-mainnet/` (to be created on the VPS) |
| **Denom** | `ukrn` (same unit, different supply/accounts) |
| **Supply** | 1,000,000,000 KRN (fixed, mint disabled) |
| **Genesis accounts** | 10: `ecosystem`, `builders`, `treasury`, `liquidity`, `community`, `strategic`, `founder`, `devs`, `advisors`, `future-hires` — **no `faucet`, no `alice`** |
| **Vesting** | `ContinuousVestingAccount` for founder (4y/1y cliff), devs (3y/6m), advisors (2y/6m) |
| **Keyring** | `~/.norynth-mainnet-keys/` on WSL (mnemonics in VeraCrypt `D:\secure.vc`) |
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
| **Accounts** | 9 (incl. faucet, alice) | 10 (no faucet, alice) | ❌ Different |
| **Addresses** | Testnet keys (VM) | Fresh mainnet keys (WSL) | ❌ Different |
| **Keyring** | `~/.norynth/keyring-test/` (VM) | `~/.norynth-mainnet-keys/` (WSL) | ❌ Different |
| **Home** | `~/.norynth/` (VM) | `~/.norynth-mainnet/` (VPS, later) | ❌ Different |
| **Denom** | `ukrn` | `ukrn` | ✅ Same (same unit) |
| **Binary** | `norynthd` (wasm) | `norynthd` (wasm) | ✅ Same code |
| **Genesis** | VM's `config/genesis.json` | `scripts/genesis-constructor/` output | ❌ Different |
| **Purpose** | Prove the path | Settlement with real value | ❌ Different |

**Same binary, different chains.** The binary (`norynthd`) is identical — it
runs both. What separates them is the **chain ID**, the **genesis**, the
**accounts**, and the **home directory**.

---

## What this means in practice

- **Do not reuse testnet keys for mainnet.** The mainnet keys are in
  `~/.norynth-mainnet-keys/`, generated fresh. The testnet keys are on the VM.
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

**Internal (not in the repo — WSL home, `~/docs/projects/`):**
- Tokenomics design: `MAINNET-TOKENOMICS-DESIGN-2026-10-25.md`
- Architecture: `MAINNET-ARCHITECTURE-2026-10-26.md`
- Readiness: `MAINNET-READINESS.md`
- Repo audit: `REPO-AUDIT-2026-10-27.md`
