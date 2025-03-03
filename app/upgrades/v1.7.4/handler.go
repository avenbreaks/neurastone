package v174

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdkvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	liquidvestingkeeper "github.com/avenbreaks/neurastone/x/liquidvesting/keeper"
	liquidvestingtypes "github.com/avenbreaks/neurastone/x/liquidvesting/types"
	vestingtypes "github.com/avenbreaks/neurastone/x/vesting/types"
)

func StretchLockupScheduleForAccounts(ctx sdk.Context, ak authkeeper.AccountKeeper, stretchLength int64, lockupLengthThreshold time.Time) error {
	logger := ctx.Logger()
	logger.Info("Stretching lockup schedules for accounts: ")

	// Iterate all accounts
	ak.IterateAccounts(ctx, func(acc authtypes.AccountI) (stop bool) {
		// Check if acc is vesting account
		vacc, ok := acc.(*vestingtypes.ClawbackVestingAccount)
		if !ok {
			return false
		}

		// if end time for unlock account is after 2026-01-01 modify schedule
		if time.Unix(vacc.GetEndTime(), 0).After(lockupLengthThreshold) {
			upcomingPeriods := liquidvestingtypes.ExtractUpcomingPeriods(vacc.GetStartTime(), vacc.GetEndTime(), vacc.LockupPeriods, ctx.BlockTime().Unix())
			pastPeriods := liquidvestingtypes.ExtractPastPeriods(vacc.GetStartTime(), vacc.GetEndTime(), vacc.LockupPeriods, ctx.BlockTime().Unix())

			// stretched upcoming periods
			stretchedUpcomingPeriods := stretchPeriods(ctx, upcomingPeriods, stretchLength, "account", vacc.GetAddress().String())

			// add 1095 days (three years to the end time)
			oldEndTime := vacc.GetEndTime()
			newEndTime := oldEndTime + OneDayInSeconds*stretchLength
			vacc.EndTime = newEndTime

			// set stretched lockup periods
			newLookupPeriods := make(sdkvesting.Periods, 0, len(pastPeriods)+len(stretchedUpcomingPeriods))
			newLookupPeriods = append(newLookupPeriods, pastPeriods...)
			newLookupPeriods = append(newLookupPeriods, stretchedUpcomingPeriods...)
			vacc.LockupPeriods = newLookupPeriods
			ak.SetAccount(ctx, vacc)

			logger.Info(fmt.Sprintf(" > account: %s — from %d to %d", vacc.GetAddress().String(), oldEndTime, newEndTime))
		}

		return false
	})

	return nil
}

func StretchLockupScheduleForLiquidVestingTokens(ctx sdk.Context, lk liquidvestingkeeper.Keeper, stretchLength int64, lockupLengthThreshold time.Time) error {
	logger := ctx.Logger()
	logger.Info("Stretching lockup schedules for liquid vesting tokens: ")

	// Iterate all denoms
	lk.IterateDenoms(ctx, func(denom liquidvestingtypes.Denom) (stop bool) {
		// if end time for liquid denom is after 2026-01-01 modify schedule
		if denom.EndTime.After(lockupLengthThreshold) {
			upcomingPeriods := liquidvestingtypes.ExtractUpcomingPeriods(denom.StartTime.Unix(), denom.EndTime.Unix(), denom.LockupPeriods, ctx.BlockTime().Unix())
			pastPeriods := liquidvestingtypes.ExtractPastPeriods(denom.StartTime.Unix(), denom.EndTime.Unix(), denom.LockupPeriods, ctx.BlockTime().Unix())

			// stretched upcoming periods
			stretchedUpcomingPeriods := stretchPeriods(ctx, upcomingPeriods, stretchLength, "liquid", denom.DisplayDenom)

			// add 1095 days (three years to the end time)
			oldEndTime := denom.EndTime.Unix()
			newEndTime := oldEndTime + OneDayInSeconds*stretchLength
			denom.EndTime = time.Unix(newEndTime, 0)

			// set stretched lockup periods
			newLookupPeriods := make(sdkvesting.Periods, 0, len(pastPeriods)+len(stretchedUpcomingPeriods))
			newLookupPeriods = append(newLookupPeriods, pastPeriods...)
			newLookupPeriods = append(newLookupPeriods, stretchedUpcomingPeriods...)
			denom.LockupPeriods = newLookupPeriods
			lk.SetDenom(ctx, denom)

			logger.Info(fmt.Sprintf(" > denom: %s — from %d to %d", denom.DisplayDenom, oldEndTime, newEndTime))
		}

		return false
	})

	return nil
}

func stretchPeriods(ctx sdk.Context, periods sdkvesting.Periods, stretchDays int64, entityType string, entity string) sdkvesting.Periods {
	const Denom = "aISLM"
	logger := ctx.Logger()

	periodsLengthInDays := periods.TotalLength() / OneDayInSeconds

	totalAmount := periods.TotalAmount().AmountOf(Denom)

	logger.Info(fmt.Sprintf(" > type: %s entity: %s — amount: %d", entityType, entity, totalAmount))

	stretchedPerDayLockupAmount := totalAmount.Quo(sdkmath.NewInt(periodsLengthInDays + stretchDays))
	extraLengthAmount := stretchedPerDayLockupAmount.Mul(sdkmath.NewInt(stretchDays))

	// update amount of existing periods
	updatedPeriods := make(sdkvesting.Periods, len(periods))
	copy(updatedPeriods, periods)

	totalUpdatedAmount := totalAmount.Sub(extraLengthAmount)

	for _, period := range updatedPeriods {
		periodAmount := period.Amount.AmountOf(Denom)
		newAmount := periodAmount.Mul(totalUpdatedAmount).Quo(totalAmount)

		for i, coin := range period.Amount {
			if coin.Denom == Denom {
				period.Amount[i] = sdk.NewCoin(Denom, newAmount)
			}
		}
	}

	// allocate extra periods
	extraPeriods := make(sdkvesting.Periods, stretchDays)
	for i := range extraPeriods {
		extraPeriods[i] = sdkvesting.Period{
			Length: OneDayInSeconds,
			Amount: sdk.NewCoins(sdk.NewCoin(Denom, stretchedPerDayLockupAmount)),
		}
	}

	// calculate total remainder and add it to the last period of the stretched periods
	updatedPeriodsAmount := updatedPeriods.TotalAmount().AmountOf(Denom)
	extraPeriodsAmount := extraPeriods.TotalAmount().AmountOf(Denom)

	calculationDiff := totalAmount.Sub(updatedPeriodsAmount.Add(extraPeriodsAmount))
	extraPeriods[stretchDays-1].Amount.Add(sdk.NewCoin(Denom, calculationDiff))

	return append(updatedPeriods, extraPeriods...)
}
