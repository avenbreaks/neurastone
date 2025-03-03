package werc20_test

import (
	"testing"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/suite"

	"github.com/avenbreaks/neurastone/precompiles/werc20"
	"github.com/avenbreaks/neurastone/testutil/integration/neura/factory"
	"github.com/avenbreaks/neurastone/testutil/integration/neura/grpc"
	testkeyring "github.com/avenbreaks/neurastone/testutil/integration/neura/keyring"
	"github.com/avenbreaks/neurastone/testutil/integration/neura/network"
)

var s *PrecompileTestSuite

// PrecompileTestSuite is the implementation of the TestSuite interface for ERC20 precompile
// unit tests.
type PrecompileTestSuite struct {
	suite.Suite

	bondDenom   string
	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring
	tokenDenom  string

	precompile *werc20.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	s = new(PrecompileTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "WISLM Extension Suite")
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	integrationNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(integrationNetwork)
	txFactory := factory.New(integrationNetwork, grpcHandler)

	s.bondDenom = integrationNetwork.GetDenom()
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring
	s.network = integrationNetwork
}
