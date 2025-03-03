package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/avenbreaks/neurastone/x/coinomics/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	params := k.GetParams(ctx)

	// NOTE: ignore end of block if coinomics is disabled
	if !params.EnableCoinomics {
		return
	}

	if err := k.MintAndAllocate(ctx); err != nil {
		ctx.Logger().Error("Failed MintAndAllocateInflation: ", err.Error())
	}
}
