# Norynth Developer Portal

## Current documentation (live)

- **Chain docs:** https://chain.norynth.com/docs/
- **API reference:** https://chain.norynth.com/docs/apis
- **Node operator guide:** https://chain.norynth.com/docs/node-operators/validator-guide
- **Tokenomics:** https://chain.norynth.com/tokenomics

## In-repo

- **Protobuf schemas:** [`../proto/README.md`](../proto/README.md)

## Historical

- **2025 Ignite tryout guides:** [`./_historical-2025/guides/`](./_historical-2025/guides/) — retired, do not follow.

## Build

```bash
# API docs (docs.go serves static/openapi.yml + template/index.tpl)
make docs

## Get testnet tokens

The Norynth testnet faucet sends test KRN to your address.

**Endpoint:** `https://faucet.chain.norynth.com/` (once DNS + nginx are configured;
during local development, use an SSH tunnel to `localhost:8000`).

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"address":"norynth1...","coins":["10000000ukrn"]}' \
  https://faucet.chain.norynth.com/
```

**Parameters:**
- `address` — your `norynth1...` address (required)
- `coins` — the amount and denom, e.g. `["10000000ukrn"]` = 10 KRN (required)

**Response:** `{}` on success (the transaction is broadcast).

**Note:** the faucet's default `--credit-amount` + `--denoms` builder has a known
issue and produces `invalid coins`. **Always pass the `coins` field explicitly.**
This is a testnet-only limitation; mainnet has no faucet.

**Rate limiting:** 1 request per minute per IP (once nginx is configured).
