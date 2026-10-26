# Norynth Chain

The settlement layer for Norynth commerce. KRN settles service sales,
escrow, and licensing on-chain.

**Status:** Testnet live (`norynth`), wasm-capable. CosmWasm contract
store/instantiate/execute proven 2026-10-26. Mainnet target: `norynth-1`.

## What this is

- Cosmos SDK v0.50.13 chain, CometBFT consensus
- CosmWasm v0.52.0 — live on testnet
- Denom: `ukrn` (1 KRN = 1,000,000 ukrn)
- Testnet chain ID: `norynth`

## Build from source

```bash
git clone https://github.com/norynth/norynth-chain.git
cd norynth-chain
CGO_ENABLED=1 go build -tags sonic.noasm -o build/norynthd ./cmd/norynthd
```

Requires Go 1.23+ and `libwasmvm.x86_64.so` (shipped with wasmvm v2.1.0).

## Run a node

Testnet join instructions: https://chain.norynth.com/docs/node-operators/validator-guide

## Documentation

- Chain docs: https://chain.norynth.com/docs/
- Cloud / infrastructure: https://cloud.norynth.com/

## License

Apache-2.0
