// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package encoding

import (
	"cosmossdk.io/simapp/params"
	"cosmossdk.io/x/tx/signing"
	amino "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"
	evmv1 "github.com/evmos/ethermint/api/ethermint/evm/v1"
	feemarketv1 "github.com/evmos/ethermint/api/ethermint/feemarket/v1"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	protov2 "google.golang.org/protobuf/proto"

	enccodec "github.com/evmos/ethermint/encoding/codec"
)

// MakeConfig creates an EncodingConfig for testing
func MakeConfig(mb module.BasicManager) params.EncodingConfig {
	signingOptions := signing.Options{
		AddressCodec: address.Bech32Codec{
			Bech32Prefix: sdk.GetConfig().GetBech32AccountAddrPrefix(),
		},
		ValidatorAddressCodec: address.Bech32Codec{
			Bech32Prefix: sdk.GetConfig().GetBech32ValidatorAddrPrefix(),
		},
	}

	// evm/MsgEthereumTx, evm/MsgUpdateParams, feemarket/MsgUpdateParams
	signingOptions.DefineCustomGetSigners(protov2.MessageName(&evmv1.MsgEthereumTx{}), evmtypes.GetSignersFromMsgEthereumTxV2)
	signingOptions.DefineCustomGetSigners(protov2.MessageName(&evmv1.MsgUpdateParams{}), evmtypes.GetSignersFromMsgUpdateParamsV2)
	signingOptions.DefineCustomGetSigners(protov2.MessageName(&feemarketv1.MsgUpdateParams{}), feemarkettypes.GetSignersFromMsgUpdateParamsV2)

	interfaceRegistry, _ := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles:     proto.HybridResolver,
		SigningOptions: signingOptions,
	})

	codec := amino.NewProtoCodec(interfaceRegistry)
	cdc := amino.NewLegacyAmino()

	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
