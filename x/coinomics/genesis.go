package coinomics

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/avenbreaks/neurastone/x/coinomics/keeper"
	"github.com/avenbreaks/neurastone/x/coinomics/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	ak types.AccountKeeper,
	_ types.StakingKeeper,
	data types.GenesisState,
) {
	// Ensure inflation module account is set on genesis
	if acc := ak.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		panic("the inflation module account has not been set")
	}

	// Set genesis state
	params := data.Params

	k.SetParams(ctx, params)

	// Set genesis state
	maxSupply := data.MaxSupply
	k.SetMaxSupply(ctx, maxSupply)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:      k.GetParams(ctx),
		PrevBlockTs: k.GetPrevBlockTS(ctx),
		MaxSupply:   k.GetMaxSupply(ctx),
	}
}
