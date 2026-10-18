package app

import (
	"context"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// WasmUpgradeName is the name of the upgrade that adds the CosmWasm store
// to an existing chain. It is referenced by the software-upgrade proposal.
const WasmUpgradeName = "wasm"

// StoreUpgrades lists the stores added in this upgrade. The store loader
// uses this to create the wasm store at the current version instead of
// expecting it to already exist. Without it, loading a pre-wasm DB fails
// with "version of store wasm mismatch root store's version".
var StoreUpgrades = storetypes.StoreUpgrades{
	Added: []string{wasmtypes.StoreKey},
}

// setupUpgradeHandlers registers the upgrade handler for the wasm upgrade.
// The store itself is created by the store loader (see New() in app.go);
// RunMigrations applies the module version-map migration.
func (app *App) setupUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(
		WasmUpgradeName,
		func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)
}
