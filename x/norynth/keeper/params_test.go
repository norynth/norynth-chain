package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "norynth/testutil/keeper"
	"norynth/x/norynth/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.NorynthKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
