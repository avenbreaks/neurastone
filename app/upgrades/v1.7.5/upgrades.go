package v175

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	erc20keeper "github.com/avenbreaks/neurastone/x/erc20/keeper"
	evmkeeper "github.com/avenbreaks/neurastone/x/evm/keeper"
	liquidvestingkeeper "github.com/avenbreaks/neurastone/x/liquidvesting/keeper"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v1.7.5
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bk bankkeeper.Keeper,
	lk liquidvestingkeeper.Keeper,
	erc20 erc20keeper.Keeper,
	ek evmkeeper.Keeper,
	ak authkeeper.AccountKeeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger()

		logger.Info("##############################################")
		logger.Info("############  RUN UPGRADE v1.7.5  ############")
		logger.Info("##############################################")

		if err := TurnOffLiquidVesting(ctx, bk, lk, erc20, ek, ak); err != nil {
			panic(err)
		}

		logger.Info("##############################################")
		logger.Info("#############  UPGRADE COMPLETE  #############")
		logger.Info("##############################################")

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
