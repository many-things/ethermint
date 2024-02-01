package evm

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	modulev1 "github.com/evmos/ethermint/api/ethermint/evm/module/v1"
	"github.com/evmos/ethermint/server/config"
	srvflags "github.com/evmos/ethermint/server/flags"
	"github.com/evmos/ethermint/x/evm/keeper"
	"github.com/evmos/ethermint/x/evm/types"
	"github.com/spf13/cast"
)

func init() {
	appmodule.Register(&modulev1.Module{}, appmodule.Provide(ProvideModule))
}

// DepInjectInput is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	ModuleKey depinject.OwnModuleKey
	Config    *modulev1.Module
	Codec     codec.Codec
	AppOpts   servertypes.AppOptions `optional:"true"`

	KVStoreKey        *storetypes.KVStoreKey
	TransientStoreKey *storetypes.TransientStoreKey
	Subspace          paramstypes.Subspace

	AuthKeeper      types.AccountKeeper
	BankKeeper      types.BankKeeper
	StakingKeeper   types.StakingKeeper
	FeeMarketKeeper types.FeeMarketKeeper

	GetStoreKeys func() map[string]storetypes.StoreKey // TODO(thai): think better way
}

// DepInjectOutput is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out

	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

// ProvideModule is a function that provides the module to the application.
func ProvideModule(in DepInjectInput) DepInjectOutput {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	tracer := config.DefaultEVMTracer
	if in.AppOpts != nil {
		tracer = cast.ToString(in.AppOpts.Get(srvflags.EVMTracer))
	}

	k := keeper.NewKeeper(
		in.Codec, in.KVStoreKey, in.TransientStoreKey,
		authority,
		in.AuthKeeper, in.BankKeeper, in.StakingKeeper, in.FeeMarketKeeper,
		tracer,
		in.Subspace,
		in.GetStoreKeys,
	)
	m := NewAppModule(k, in.AuthKeeper, in.Subspace)

	return DepInjectOutput{Keeper: k, Module: m}
}
