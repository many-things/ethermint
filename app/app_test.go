package app

import (
	"github.com/cometbft/cometbft/abci/types"
	cmtlog "github.com/cometbft/cometbft/libs/log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/evmos/ethermint/encoding"
)

func TestEthermintAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := SetupWithDB(false, nil, db)

	_, err := app.FinalizeBlock(&types.RequestFinalizeBlock{
		Height: app.LastBlockHeight() + 1,
	})
	require.NoError(t, err, "ExportAppStateAndValidators FinalizeBlock failed")
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewEthermintApp(
		CometLoggerAdapter{Logger: cmtlog.NewTMLogger(cmtlog.NewSyncWriter(os.Stdout))},
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		0,
		encoding.MakeConfig(ModuleBasics),
		simtestutil.NewAppOptionsWithFlagHome(DefaultNodeHome),
		baseapp.SetChainID(ChainID),
	)
	_, err = app2.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
