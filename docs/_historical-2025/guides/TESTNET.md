# Testnet Deployment

## 1. Initialize Node
```bash
noynd init mynode --chain-id noyn-testnet-1
```

## 2. Configure Peers
Edit `config/config.toml`:
```toml
persistent_peers = "[testnet-peers]"
```

## 3. Start Node
```bash
noynd start
```
