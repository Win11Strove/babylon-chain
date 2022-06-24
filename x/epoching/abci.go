package epoching

import (
	"time"

	"github.com/babylonchain/babylon/x/epoching/keeper"
	"github.com/babylonchain/babylon/x/epoching/types"

	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	defer telemetry.ModuleMeasureSince(evidencetypes.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	defer telemetry.ModuleMeasureSince(slashingtypes.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// TODO: unimplemented:
	// - slashing equivocating/unlive validators without jailing them
}

// Called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	defer telemetry.ModuleMeasureSince(stakingtypes.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// TODO: unimplemented:
	// - if an epoch is newly checkpointed, make unbonding validators/delegations in this epoch unbonded
	// - if reaching an epoch boundary, execute validator-related msgs (bonded -> unbonding)
	return nil
}
