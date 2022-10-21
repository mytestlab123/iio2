package iio2_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iio2/testutil/keeper"
	"github.com/mytestlab123/iio2/testutil/nullify"
	"github.com/mytestlab123/iio2/x/iio2"
	"github.com/mytestlab123/iio2/x/iio2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Iio2Keeper(t)
	iio2.InitGenesis(ctx, *k, genesisState)
	got := iio2.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
