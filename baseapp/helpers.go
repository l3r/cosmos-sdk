package baseapp

import (
	"github.com/tendermint/abci/server"
	abci "github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// nolint - Mostly for testing
func (app *BaseApp) Check(tx sdk.Tx) (result sdk.Result) {
	return app.runTx(runTxModeCheck, nil, tx)
}

// nolint - full tx execution
func (app *BaseApp) Simulate(tx sdk.Tx) (result sdk.Result) {
	return app.runTx(runTxModeSimulate, nil, tx)
}

// nolint
func (app *BaseApp) Deliver(tx sdk.Tx) (result sdk.Result) {
	return app.runTx(runTxModeDeliver, nil, tx)
}

// RunForever - BasecoinApp execution and cleanup
func RunForever(app abci.Application) {

	// Start the ABCI server
	srv, err := server.NewServer("0.0.0.0:26658", "socket", app)
	if err != nil {
		cmn.Exit(err.Error())
	}
	srv.Start()

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
}
