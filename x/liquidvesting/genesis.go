package liquidvesting

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/avenbreaks/neurastone/x/liquidvesting/keeper"
	"github.com/avenbreaks/neurastone/x/liquidvesting/types"
)

// InitGenesis import module genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	// Set genesis state
	err := k.SetParams(ctx, data.Params)
	if err != nil {
		panic(fmt.Errorf("error setting params %s", err))
	}

	k.SetDenomCounter(ctx, data.DenomCounter)

	// Set all the request
	for _, denom := range data.Denoms {
		k.SetDenom(ctx, denom)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:       k.GetParams(ctx),
		DenomCounter: k.GetDenomCounter(ctx),
		Denoms:       k.GetAllDenoms(ctx),
	}
}
