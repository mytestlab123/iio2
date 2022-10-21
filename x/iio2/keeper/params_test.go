package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iio2/testutil/keeper"
	"github.com/mytestlab123/iio2/x/iio2/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Iio2Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
