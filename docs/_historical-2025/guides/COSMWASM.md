# Smart Contract Development

Norynth supports CosmWasm smart contracts for decentralized logic. This guide covers:
- Wasm module installation
- Contract scaffolding
- Deployment workflows

# Smart Contract Development

## 1. Set Up CosmWasm
```bash
ignite chain install wasm
```

## 2. Create Your First Contract
```bash
ignite scaffold wasm [CONTRACT_NAME] --signer [KEY_NAME]
```

## 3. Interact with Contract
```bash
# Store contract
noynd tx wasm store [WASM_FILE].wasm --from [KEY_NAME]

# Instantiate
noynd tx wasm instantiate [CODE_ID] '{"init":{}}' --from [KEY_NAME] --label "[LABEL]"
```
