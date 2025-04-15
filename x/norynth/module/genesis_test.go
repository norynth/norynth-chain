package norynth_test

import (
	"testing"

	keepertest "norynth/testutil/keeper"
	"norynth/testutil/nullify"
	norynth "norynth/x/norynth/module"
	"norynth/x/norynth/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NorynthKeeper(t)
	norynth.InitGenesis(ctx, k, genesisState)
	got := norynth.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
