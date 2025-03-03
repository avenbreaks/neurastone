package keeper_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcgotesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/avenbreaks/neurastone/app"
	ibctesting "github.com/avenbreaks/neurastone/ibc/testing"
	"github.com/avenbreaks/neurastone/x/erc20/types"
	evm "github.com/avenbreaks/neurastone/x/evm/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx              sdk.Context
	app              *app.neura
	queryClientEvm   evm.QueryClient
	queryClient      types.QueryClient
	address          common.Address
	consAddress      sdk.ConsAddress
	clientCtx        client.Context //nolint:unused
	ethSigner        ethtypes.Signer
	priv             cryptotypes.PrivKey
	validator        stakingtypes.Validator
	signer           keyring.Signer
	mintFeeCollector bool

	coordinator *ibcgotesting.Coordinator

	// testing chains used for convenience and readability
	neuraChain       *ibcgotesting.TestChain
	IBCOsmosisChain *ibcgotesting.TestChain
	IBCCosmosChain  *ibcgotesting.TestChain

	pathOsmosisneura   *ibctesting.Path
	pathCosmosneura    *ibctesting.Path
	pathOsmosisCosmos *ibctesting.Path

	suiteIBCTesting bool
}

var (
	s *KeeperTestSuite
	// sendAndReceiveMsgFee corresponds to the fees paid on neura Network chain when calling the SendAndReceive function
	// This function makes 3 cosmos txs under the hood
	sendAndReceiveMsgFee = math.NewInt(ibctesting.DefaultFeeAmt * 3)
	// sendBackCoinsFee corresponds to the fees paid on neura Network chain when calling the SendBackCoins function
	// or calling the SendAndReceive from another chain to neura Network
	// This function makes 2 cosmos txs under the hood
	sendBackCoinsFee = math.NewInt(ibctesting.DefaultFeeAmt * 2)
)

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.DoSetupTest(suite.T())
}
