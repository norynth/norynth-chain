# Protocol Buffer Definitions

## Directory Structure

```
proto/
├── norynth/norynth/   # Chain-specific types (genesis, params, query, tx, module)
├── buf.gen.*.yaml     # Codegen configs (gogo, pulsar, sta, swagger, ts)
├── buf.lock           # Dependency lock
└── buf.yaml           # Buf workspace config
```

There is no `third_party/` directory — the Cosmos SDK and IBC protos are
resolved by `buf` from `buf.lock`, not vendored.

## Compilation

```bash
make proto-gen            # rebuild all protobufs
make proto-swagger-gen    # regenerate Swagger/OpenAPI
```

## Versions

- Cosmos SDK **v0.50.13** (matches `go.mod` and the binary's
  `cosmos_sdk_version`)
- `buf` resolves dependencies from `buf.lock` — do not hand-edit it

## Packages

| Package           | Description             |
|-------------------|-------------------------|
| `norynth.norynth` | Core chain module types |
