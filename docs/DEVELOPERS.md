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
