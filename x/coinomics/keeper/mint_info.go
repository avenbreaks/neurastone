package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/avenbreaks/neurastone/x/coinomics/types"
)

func (k Keeper) GetPrevBlockTS(ctx sdk.Context) sdkmath.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixPrevBlockTS)
	if len(bz) == 0 {
		return sdk.ZeroInt()
	}

	var prevBlockTSValue sdkmath.Int
	err := prevBlockTSValue.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal prevBlockTSValue value: %w", err))
	}

	return prevBlockTSValue
}

func (k Keeper) SetPrevBlockTS(ctx sdk.Context, prevBlockTS sdkmath.Int) {
	binaryInfValue, err := prevBlockTS.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value: %w", err))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixPrevBlockTS, binaryInfValue)
}

func (k Keeper) GetMaxSupply(ctx sdk.Context) sdk.Coin {
	params := k.GetParams(ctx)

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixMaxSupply)
	if len(bz) == 0 {
		return sdk.NewCoin(params.MintDenom, sdk.ZeroInt())
	}

	var maxSupply sdk.Coin
	err := maxSupply.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal maxSupply value: %w", err))
	}

	return maxSupply
}

func (k Keeper) SetMaxSupply(ctx sdk.Context, maxSupply sdk.Coin) {
	binaryMaxSupply, err := maxSupply.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value: %w", err))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixMaxSupply, binaryMaxSupply)
}
