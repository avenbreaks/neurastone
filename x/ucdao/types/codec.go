package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"

	cryptocodec "github.com/avenbreaks/neurastone/crypto/codec"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the dao module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgFund{}, "neura/ucdao/MsgFund")
	legacy.RegisterAminoMsg(cdc, &MsgFundLegacy{}, "neura/dao/MsgFund")
	legacy.RegisterAminoMsg(cdc, &MsgTransferOwnership{}, "neura/ucdao/MsgTransferOwnership")
	legacy.RegisterAminoMsg(cdc, &MsgTransferOwnershipWithRatio{}, "neura/ucdao/MsgTransferWithRatio")
	legacy.RegisterAminoMsg(cdc, &MsgTransferOwnershipWithAmount{}, "neura/ucdao/MsgTransferWithAmount")

	cdc.RegisterConcrete(Params{}, "neura/x/ucdao/Params", nil)
}

// RegisterInterfaces registers the interfaces types with the Interface Registry.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgFund{},
		&MsgTransferOwnership{},
		&MsgTransferOwnershipWithRatio{},
		&MsgTransferOwnershipWithAmount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz and dao Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
