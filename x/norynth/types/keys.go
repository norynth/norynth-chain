package types

const (
	// ModuleName defines the module name
	ModuleName = "norynth"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_norynth"
)

var (
	ParamsKey = []byte("p_norynth")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
