package ante_test

import (
	"time"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/avenbreaks/neurastone/crypto/ethsecp256k1"
	"github.com/avenbreaks/neurastone/testutil"
	testutiltx "github.com/avenbreaks/neurastone/testutil/tx"
	"github.com/avenbreaks/neurastone/utils"
)

var _ = DescribeTableSubtree("when sending a Cosmos transaction", func(signMode signing.SignMode) {
	var (
		addr sdk.AccAddress
		priv *ethsecp256k1.PrivKey
		msg  sdk.Msg
	)

	BeforeEach(func() {
		s.SetupTest()
	})

	Context("and the sender account has enough balance to pay for the transaction cost", Ordered, func() {
		var (
			rewardsAmt = sdkmath.NewInt(1e5)
			balance    = sdkmath.NewInt(1e18)
		)

		BeforeEach(func() {
			addr, priv = testutiltx.NewAccAddressAndKey()

			msg = &banktypes.MsgSend{
				FromAddress: addr.String(),
				ToAddress:   "neura1hdr0lhv75vesvtndlh78ck4cez6esz8u2lk0hq",
				Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(1e14), Denom: utils.BaseDenom}},
			}

			s.ctx, _ = testutil.PrepareAccountsForDelegationRewards(
				s.T(), s.ctx, s.app, addr, balance, rewardsAmt,
			)

			var err error
			s.ctx, err = testutil.CommitAndCreateNewCtx(s.ctx, s.app, time.Second*0, nil)
			Expect(err).To(BeNil())
		})

		It("should succeed & not withdraw any staking rewards", func() {
			res, err := testutil.DeliverTx(s.ctx, s.app, priv, nil, signMode, msg)
			Expect(err).To(BeNil())
			Expect(res.IsOK()).To(BeTrue())

			rewards, err := testutil.GetTotalDelegationRewards(s.ctx, s.app.DistrKeeper, addr)
			Expect(err).To(BeNil())
			Expect(rewards).To(Equal(sdk.NewDecCoins(sdk.NewDecCoin(utils.BaseDenom, rewardsAmt))))
		})
	})

	Context("and the sender account neither has enough balance nor sufficient staking rewards to pay for the transaction cost", func() {
		var (
			rewardsAmt = sdkmath.NewInt(0)
			balance    = sdkmath.NewInt(0)
		)

		BeforeEach(func() {
			addr, priv = testutiltx.NewAccAddressAndKey()

			msg = &banktypes.MsgSend{
				FromAddress: addr.String(),
				ToAddress:   "neura1hdr0lhv75vesvtndlh78ck4cez6esz8u2lk0hq",
				Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(1e14), Denom: utils.BaseDenom}},
			}

			s.ctx, _ = testutil.PrepareAccountsForDelegationRewards(
				s.T(), s.ctx, s.app, addr, balance, rewardsAmt,
			)

			var err error
			s.ctx, err = testutil.CommitAndCreateNewCtx(s.ctx, s.app, time.Second*0, nil)
			Expect(err).To(BeNil())
		})

		It("should fail", func() {
			res, err := testutil.DeliverTx(s.ctx, s.app, priv, nil, signMode, msg)
			Expect(res.IsOK()).To(BeTrue())
			Expect(err).To(HaveOccurred())
		})

		It("should not withdraw any staking rewards", func() {
			rewards, err := testutil.GetTotalDelegationRewards(s.ctx, s.app.DistrKeeper, addr)
			Expect(err).To(BeNil())
			Expect(rewards.Empty()).To(BeTrue())
		})
	})

	Context("and the sender account has not enough balance but sufficient staking rewards to pay for the transaction cost", func() {
		var (
			rewardsAmt = sdkmath.NewInt(1e18)
			balance    = sdkmath.NewInt(0)
		)

		BeforeEach(func() {
			addr, priv = testutiltx.NewAccAddressAndKey()

			msg = &banktypes.MsgSend{
				FromAddress: addr.String(),
				ToAddress:   "neura1hdr0lhv75vesvtndlh78ck4cez6esz8u2lk0hq",
				Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(1), Denom: utils.BaseDenom}},
			}

			s.ctx, _ = testutil.PrepareAccountsForDelegationRewards(
				s.T(), s.ctx, s.app, addr, balance, rewardsAmt,
			)
			var err error
			s.ctx, err = testutil.CommitAndCreateNewCtx(s.ctx, s.app, time.Second*0, nil)
			Expect(err).To(BeNil())
		})

		It("should withdraw enough staking rewards to cover the transaction cost", func() {
			rewards, err := testutil.GetTotalDelegationRewards(s.ctx, s.app.DistrKeeper, addr)
			Expect(err).To(BeNil())
			Expect(rewards).To(Equal(sdk.NewDecCoins(sdk.NewDecCoin(utils.BaseDenom, rewardsAmt))))

			balance := s.app.BankKeeper.GetBalance(s.ctx, addr, utils.BaseDenom)
			Expect(balance.Amount).To(Equal(sdkmath.NewInt(0)))

			res, err := testutil.DeliverTx(s.ctx, s.app, priv, nil, signMode, msg)
			Expect(res.IsOK()).To(BeTrue())
			Expect(err).To(BeNil())
		})
	})
},
	Entry("Direct sign mode", signing.SignMode_SIGN_MODE_DIRECT),
	Entry("Legacy Amino JSON sign mode", signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON),
)
