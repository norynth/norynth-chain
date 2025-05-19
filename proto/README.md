# Protocol Buffer Definitions

## Directory Structure
```diff
- proto/
+ proto/                # Root of all protocol buffers
   ├── norynth/         # Chain-specific types
   │   ├── genesis.proto
   │   ├── params.proto
   │   └── tx.proto
   └── third_party/     # Cosmos SDK dependencies
       ├── cosmos/      # Cosmos SDK types
       └── ibc/        # IBC-related definitions
```

## Compilation
```bash
# Rebuild all protobufs
make proto-gen

# Generate Swagger/OpenAPI documentation
make proto-swagger-gen
```

## Versioning
- Requires `libprotoc` 3.21.12
- Compatible with Cosmos SDK `v0.47.0+`

## Key Packages

| Package               | Description                |
|-----------------------|----------------------------|
| `norynth.norynth`     | Core blockchain types      |
| `cosmos.bank.v1beta1` | Token operations           |
| `ibc.core.channel.v1` | IBC channel logic          |

