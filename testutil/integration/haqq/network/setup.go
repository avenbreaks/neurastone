package network

import (
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	simutils "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/avenbreaks/neurastone/app"
	"github.com/avenbreaks/neurastone/encoding"
	neuratypes "github.com/avenbreaks/neurastone/types"
	coinomicstypes "github.com/avenbreaks/neurastone/x/coinomics/types"
)

// createValidatorSetAndSigners creates validator set with the amount of validators specified
// with the default power of 1.
func createValidatorSetAndSigners(numberOfValidators int) (*tmtypes.ValidatorSet, map[string]tmtypes.PrivValidator) {
	// Create validator set
	tmValidators := make([]*tmtypes.Validator, 0, numberOfValidators)
	signers := make(map[string]tmtypes.PrivValidator, numberOfValidators)

	for i := 0; i < numberOfValidators; i++ {
		privVal := mock.NewPV()
		pubKey, _ := privVal.GetPubKey()
		validator := tmtypes.NewValidator(pubKey, 1)
		tmValidators = append(tmValidators, validator)
		signers[pubKey.Address().String()] = privVal
	}

	return tmtypes.NewValidatorSet(tmValidators), signers
}

// createGenesisAccounts returns a slice of genesis accounts from the given
// account addresses.
func createGenesisAccounts(accounts []sdktypes.AccAddress) []authtypes.GenesisAccount {
	numberOfAccounts := len(accounts)
	genAccounts := make([]authtypes.GenesisAccount, 0, numberOfAccounts)
	emptyCodeHash := crypto.Keccak256Hash(nil).String()
	for _, acc := range accounts {
		baseAcc := authtypes.NewBaseAccount(acc, nil, 0, 0)
		ethAcc := &neuratypes.EthAccount{
			BaseAccount: baseAcc,
			CodeHash:    emptyCodeHash,
		}
		genAccounts = append(genAccounts, ethAcc)
	}
	return genAccounts
}

// getAccAddrsFromBalances returns a slice of genesis accounts from the
// given balances.
func getAccAddrsFromBalances(balances []banktypes.Balance) []sdktypes.AccAddress {
	numberOfBalances := len(balances)
	genAccounts := make([]sdktypes.AccAddress, 0, numberOfBalances)
	for _, balance := range balances {
		genAccounts = append(genAccounts, balance.GetAddress())
	}
	return genAccounts
}

// createBalances creates balances for the given accounts and coin
func createBalances(accounts []sdktypes.AccAddress, coin sdktypes.Coin) []banktypes.Balance {
	numberOfAccounts := len(accounts)
	fundedAccountBalances := make([]banktypes.Balance, 0, numberOfAccounts)
	for _, acc := range accounts {
		balance := banktypes.Balance{
			Address: acc.String(),
			Coins:   sdktypes.NewCoins(coin),
		}

		fundedAccountBalances = append(fundedAccountBalances, balance)
	}
	return fundedAccountBalances
}

// createneuraApp creates a neura app
func createneuraApp(chainID string) *app.neura {
	// Create neura app
	db := dbm.NewMemDB()
	logger := log.NewNopLogger()
	loadLatest := true
	skipUpgradeHeights := map[int64]bool{}
	homePath := app.DefaultNodeHome
	invCheckPeriod := uint(5)
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	appOptions := simutils.NewAppOptionsWithFlagHome(app.DefaultNodeHome)
	baseAppOptions := []func(*baseapp.BaseApp){baseapp.SetChainID(chainID)}

	return app.Newneura(
		logger,
		db,
		nil,
		loadLatest,
		skipUpgradeHeights,
		homePath,
		invCheckPeriod,
		encodingConfig,
		appOptions,
		baseAppOptions...,
	)
}

// createStakingValidator creates a staking validator from the given tm validator and bonded
func createStakingValidator(val *tmtypes.Validator, bondedAmt sdkmath.Int) (stakingtypes.Validator, error) {
	pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
	if err != nil {
		return stakingtypes.Validator{}, err
	}

	pkAny, err := codectypes.NewAnyWithValue(pk)
	if err != nil {
		return stakingtypes.Validator{}, err
	}

	commission := stakingtypes.NewCommission(sdktypes.ZeroDec(), sdktypes.ZeroDec(), sdktypes.ZeroDec())
	validator := stakingtypes.Validator{
		OperatorAddress:   sdktypes.ValAddress(val.Address).String(),
		ConsensusPubkey:   pkAny,
		Jailed:            false,
		Status:            stakingtypes.Bonded,
		Tokens:            bondedAmt,
		DelegatorShares:   sdktypes.OneDec(),
		Description:       stakingtypes.Description{},
		UnbondingHeight:   int64(0),
		UnbondingTime:     time.Unix(0, 0).UTC(),
		Commission:        commission,
		MinSelfDelegation: sdktypes.ZeroInt(),
	}
	return validator, nil
}

// createStakingValidators creates staking validators from the given tm validators and bonded
// amounts
func createStakingValidators(tmValidators []*tmtypes.Validator, bondedAmt sdkmath.Int) ([]stakingtypes.Validator, error) {
	amountOfValidators := len(tmValidators)
	stakingValidators := make([]stakingtypes.Validator, 0, amountOfValidators)
	for _, val := range tmValidators {
		validator, err := createStakingValidator(val, bondedAmt)
		if err != nil {
			return nil, err
		}
		stakingValidators = append(stakingValidators, validator)
	}
	return stakingValidators, nil
}

// createDelegations creates delegations for the given validators and account
func createDelegations(tmValidators []*tmtypes.Validator, fromAccount sdktypes.AccAddress) []stakingtypes.Delegation {
	amountOfValidators := len(tmValidators)
	delegations := make([]stakingtypes.Delegation, 0, amountOfValidators)
	for _, val := range tmValidators {
		delegation := stakingtypes.NewDelegation(fromAccount, val.Address.Bytes(), sdktypes.OneDec())
		delegations = append(delegations, delegation)
	}
	return delegations
}

// StakingCustomGenesisState defines the staking genesis state
type StakingCustomGenesisState struct {
	denom string

	validators  []stakingtypes.Validator
	delegations []stakingtypes.Delegation
}

// setStakingGenesisState sets the staking genesis state
func setStakingGenesisState(neuraApp *app.neura, genesisState neuratypes.GenesisState, overwriteParams StakingCustomGenesisState) neuratypes.GenesisState {
	// Set staking params
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = overwriteParams.denom

	stakingGenesis := stakingtypes.NewGenesisState(stakingParams, overwriteParams.validators, overwriteParams.delegations)
	genesisState[stakingtypes.ModuleName] = neuraApp.AppCodec().MustMarshalJSON(stakingGenesis)
	return genesisState
}

// setAuthGenesisState sets the auth genesis state
func setAuthGenesisState(neuraApp *app.neura, genesisState neuratypes.GenesisState, genAccounts []authtypes.GenesisAccount) neuratypes.GenesisState {
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccounts)
	genesisState[authtypes.ModuleName] = neuraApp.AppCodec().MustMarshalJSON(authGenesis)
	return genesisState
}

// setCoinomicsGenesisState sets the coinomics genesis state
func setCoinomicsGenesisState(neuraApp *app.neura, genesisState neuratypes.GenesisState) neuratypes.GenesisState {
	coinomicsParams := coinomicstypes.DefaultParams()
	coinomicsParams.EnableCoinomics = false

	coinomicsGenesis := coinomicstypes.NewGenesisState(coinomicsParams, sdktypes.NewCoin("aISLM", sdkmath.NewInt(10_000_000_000)))
	genesisState[coinomicstypes.ModuleName] = neuraApp.AppCodec().MustMarshalJSON(&coinomicsGenesis)
	return genesisState
}

type BankCustomGenesisState struct {
	totalSupply sdktypes.Coins
	balances    []banktypes.Balance
}

// setBankGenesisState sets the bank genesis state
func setBankGenesisState(neuraApp *app.neura, genesisState neuratypes.GenesisState, overwriteParams BankCustomGenesisState) neuratypes.GenesisState {
	bankGenesis := banktypes.NewGenesisState(
		banktypes.DefaultGenesisState().Params,
		overwriteParams.balances,
		overwriteParams.totalSupply,
		[]banktypes.Metadata{},
		[]banktypes.SendEnabled{},
	)
	genesisState[banktypes.ModuleName] = neuraApp.AppCodec().MustMarshalJSON(bankGenesis)
	return genesisState
}

// calculateTotalSupply calculates the total supply from the given balances
func calculateTotalSupply(fundedAccountsBalances []banktypes.Balance) sdktypes.Coins {
	totalSupply := sdktypes.NewCoins()
	for _, balance := range fundedAccountsBalances {
		totalSupply = totalSupply.Add(balance.Coins...)
	}
	return totalSupply
}

// addBondedModuleAccountToFundedBalances adds bonded amount to bonded pool module account and include it on funded accounts
func addBondedModuleAccountToFundedBalances(fundedAccountsBalances []banktypes.Balance, totalBonded sdktypes.Coin) []banktypes.Balance {
	return append(fundedAccountsBalances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdktypes.Coins{totalBonded},
	})
}
