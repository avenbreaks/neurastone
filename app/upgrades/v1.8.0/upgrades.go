package v180

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/avenbreaks/neurastone/utils"
	evmkeeper "github.com/avenbreaks/neurastone/x/evm/keeper"
	evmtypes "github.com/avenbreaks/neurastone/x/evm/types"
	ucdaokeeper "github.com/avenbreaks/neurastone/x/ucdao/keeper"
	ucdaotypes "github.com/avenbreaks/neurastone/x/ucdao/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v1.8.0
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	ek evmkeeper.Keeper,
	bk bankkeeper.Keeper,
	dk ucdaokeeper.Keeper,
	ucDaoStoreKey storetypes.StoreKey,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		logger.Info("migrate UC DAO balance")
		if err := migrateUCDAObalance(ctx, bk); err != nil {
			logger.Error("error while migrating ucdao module balance", "error", err)
			return nil, err
		}

		logger.Info("fix UC DAO total ISLM balance")
		if err := fixUCDAOTotalBalance(ctx, dk, ucDaoStoreKey); err != nil {
			logger.Error("error while reducing ucdao total ISLM balance", "error", err)
			return nil, err
		}

		logger.Info("setting precompiles parameters")
		if err := setPrecompilesParams(ctx, ek); err != nil {
			logger.Error("error while setting precompiles parameters", "error", err)
			return nil, err
		}

		// Leave modules are as-is to avoid running InitGenesis.
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func setPrecompilesParams(ctx sdk.Context, ek evmkeeper.Keeper) error {
	params := ek.GetParams(ctx)
	params.ActivePrecompiles = evmtypes.AvailableEVMExtensions
	return ek.SetParams(ctx, params)
}

func migrateUCDAObalance(ctx sdk.Context, bk bankkeeper.Keeper) error {
	oldDaoAddress := "neura1vwr8z00ty7mqnk4dtchr9mn9j96nuh6wme0t2z"
	oldDaoAccAddr := sdk.MustAccAddressFromBech32(oldDaoAddress)
	oldDaoBalances := bk.GetAllBalances(ctx, oldDaoAccAddr)

	return bk.SendCoinsFromAccountToModule(ctx, oldDaoAccAddr, ucdaotypes.ModuleName, oldDaoBalances)
}

func fixUCDAOTotalBalance(ctx sdk.Context, dk ucdaokeeper.Keeper, ucDaoStoreKey storetypes.StoreKey) error {
	// Reduce total balance for 20000000000000000000aISLM
	logger := ctx.Logger().With("upgrade", UpgradeName)
	balISLM := dk.GetTotalBalanceOf(ctx, utils.BaseDenom)
	logger.Info(fmt.Sprintf("Old ISLM balance in UC DAO: %s", balISLM.String()))

	amt, err := sdk.ParseCoinNormalized("20000000000000000000aISLM")
	if err != nil {
		return err
	}
	balISLM = balISLM.Sub(amt)

	intBytes, err := balISLM.Amount.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value %v", err))
	}

	store := ctx.KVStore(ucDaoStoreKey)
	supplyStore := prefix.NewStore(store, ucdaotypes.TotalBalanceKey)

	// Bank invariants and IBC requires to remove zero coins.
	if balISLM.IsZero() {
		supplyStore.Delete(utils.UnsafeStrToBytes(balISLM.GetDenom()))
	} else {
		supplyStore.Set([]byte(balISLM.GetDenom()), intBytes)
	}

	newBalISLM := dk.GetTotalBalanceOf(ctx, utils.BaseDenom)
	logger.Info(fmt.Sprintf("New ISLM balance in UC DAO: %s", newBalISLM.String()))

	return nil
}
