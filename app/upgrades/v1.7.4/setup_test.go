package v174_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/avenbreaks/neurastone/app"
	"github.com/avenbreaks/neurastone/crypto/ethsecp256k1"
	"github.com/avenbreaks/neurastone/encoding"
	"github.com/avenbreaks/neurastone/testutil"
	utiltx "github.com/avenbreaks/neurastone/testutil/tx"
	neuratypes "github.com/avenbreaks/neurastone/types"
	"github.com/avenbreaks/neurastone/utils"
	epochstypes "github.com/avenbreaks/neurastone/x/epochs/types"
	evmtypes "github.com/avenbreaks/neurastone/x/evm/types"
)

var s *UpgradeTestSuite

type UpgradeTestSuite struct {
	suite.Suite

	ctx     sdk.Context
	app     *app.neura
	t       require.TestingT
	chainID string
	// account
	accPrivateKey *ethsecp256k1.PrivKey
	ethAddress    common.Address
	accAddress    sdk.AccAddress
	signer        keyring.Signer

	// validator
	valAddr   sdk.ValAddress
	validator stakingtypes.Validator
	// consensus
	consAddress sdk.ConsAddress
	// query client
	clientCtx client.Context
	// signer
	ethSigner ethtypes.Signer
}

func TestUpgradeTestSuite(t *testing.T) {
	s = new(UpgradeTestSuite)
	suite.Run(t, s)
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.DoSetupTest(suite.T())
}

func (suite *UpgradeTestSuite) DoSetupTest(t require.TestingT) {
	const fixedTimestamp = 1712842608
	var err error
	checkTx := false
	suite.t = t
	suite.chainID = utils.MainNetChainID + "-1"

	// account key
	suite.accPrivateKey, err = ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	suite.ethAddress = common.BytesToAddress(suite.accPrivateKey.PubKey().Address().Bytes())
	suite.accAddress = sdk.AccAddress(suite.ethAddress.Bytes())
	suite.signer = utiltx.NewSigner(suite.accPrivateKey)

	// consensus key
	consPrivateKey, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	suite.consAddress = sdk.ConsAddress(consPrivateKey.PubKey().Address())

	// Init app
	suite.app, _ = app.Setup(checkTx, nil, suite.chainID)
	// Set Context
	header := testutil.NewHeader(1, time.Unix(fixedTimestamp, 0).UTC(), suite.chainID, suite.consAddress, nil, nil)
	suite.ctx = suite.app.BaseApp.NewContext(checkTx, header)
	// Setup query helpers
	queryHelperEvm := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evmtypes.RegisterQueryServer(queryHelperEvm, suite.app.EvmKeeper)

	// Set epoch start time and height for all epoch identifiers from the epoch
	// module
	identifiers := []string{epochstypes.WeekEpochID, epochstypes.DayEpochID}
	for _, identifier := range identifiers {
		epoch, found := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, identifier)
		require.True(t, found)
		epoch.StartTime = suite.ctx.BlockTime()
		epoch.CurrentEpochStartHeight = suite.ctx.BlockHeight()
		suite.app.EpochsKeeper.SetEpochInfo(suite.ctx, epoch)
	}

	acc := &neuratypes.EthAccount{
		BaseAccount: authtypes.NewBaseAccount(suite.accAddress, nil, 0, 0),
		CodeHash:    common.BytesToHash(crypto.Keccak256(nil)).String(),
	}

	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	// fund signer acc to pay for tx fees
	amt := sdk.NewInt(int64(math.Pow10(18) * 2))
	err = testutil.FundAccount(
		suite.ctx,
		suite.app.BankKeeper,
		suite.accAddress,
		sdk.NewCoins(sdk.NewCoin(suite.app.StakingKeeper.BondDenom(suite.ctx), amt)),
	)
	require.NoError(t, err)

	// Set Validator
	suite.valAddr = sdk.ValAddress(suite.accAddress)
	validator, err := stakingtypes.NewValidator(suite.valAddr, consPrivateKey.PubKey(), stakingtypes.Description{})
	require.NoError(t, err)
	validator = stakingkeeper.TestingUpdateValidator(suite.app.StakingKeeper.Keeper, suite.ctx, validator, true)
	err = suite.app.StakingKeeper.Hooks().AfterValidatorCreated(suite.ctx, validator.GetOperator())
	require.NoError(t, err)
	err = suite.app.StakingKeeper.SetValidatorByConsAddr(suite.ctx, validator)
	require.NoError(t, err)
	validators := suite.app.StakingKeeper.GetValidators(suite.ctx, 1)
	suite.validator = validators[0]

	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
	suite.ethSigner = ethtypes.LatestSignerForChainID(suite.app.EvmKeeper.ChainID())
}
