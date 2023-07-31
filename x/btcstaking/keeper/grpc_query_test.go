package keeper_test

import (
	"math/rand"
	"testing"

	"github.com/babylonchain/babylon/testutil/datagen"
	testkeeper "github.com/babylonchain/babylon/testutil/keeper"
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/babylonchain/babylon/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func FuzzActivatedHeight(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		// Setup keeper and context
		keeper, ctx := testkeeper.BTCStakingKeeper(t, nil, nil)
		ctx = sdk.UnwrapSDKContext(ctx)

		// not activated yet
		_, err := keeper.GetBTCStakingActivatedHeight(ctx)
		require.Error(t, err)

		randomActivatedHeight := datagen.RandomInt(r, 100) + 1
		btcVal, err := datagen.GenRandomBTCValidator(r)
		require.NoError(t, err)
		keeper.SetVotingPower(ctx, btcVal.BtcPk.MustMarshal(), randomActivatedHeight, uint64(10))

		// now it's activated
		resp, err := keeper.ActivatedHeight(ctx, &types.QueryActivatedHeightRequest{})
		require.NoError(t, err)
		require.Equal(t, randomActivatedHeight, resp.Height)
	})
}

func FuzzBTCValidators(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		// Setup keeper and context
		keeper, ctx := testkeeper.BTCStakingKeeper(t, nil, nil)
		ctx = sdk.UnwrapSDKContext(ctx)

		// Generate random btc validators and add them to kv store
		btcValsMap := make(map[string]*types.BTCValidator)
		for i := 0; i < int(datagen.RandomInt(r, 10)+1); i++ {
			btcVal, err := datagen.GenRandomBTCValidator(r)
			require.NoError(t, err)

			keeper.SetBTCValidator(ctx, btcVal)
			btcValsMap[btcVal.BtcPk.MarshalHex()] = btcVal
		}
		numOfBTCValsInStore := len(btcValsMap)

		// Test nil request
		resp, err := keeper.BTCValidators(ctx, nil)
		if resp != nil {
			t.Errorf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Generate a page request with a limit and a nil key
		limit := datagen.RandomInt(r, numOfBTCValsInStore) + 1
		pagination := constructRequestWithLimit(r, limit)
		// Generate the initial query
		req := types.QueryBTCValidatorsRequest{Pagination: pagination}
		// Construct a mapping from the btc vals found to a boolean value
		// Will be used later to evaluate whether all the btc vals were returned
		btcValsFound := make(map[string]bool, 0)

		for i := uint64(0); i < uint64(numOfBTCValsInStore); i += limit {
			resp, err = keeper.BTCValidators(ctx, &req)
			if err != nil {
				t.Errorf("Valid request led to an error %s", err)
			}
			if resp == nil {
				t.Fatalf("Valid request led to a nil response")
			}

			for _, val := range resp.BtcValidators {
				// Check if the pk exists in the map
				if _, ok := btcValsMap[val.BtcPk.MarshalHex()]; !ok {
					t.Fatalf("rpc returned a val that was not created")
				}
				btcValsFound[val.BtcPk.MarshalHex()] = true
			}

			// Construct the next page request
			pagination = constructRequestWithKeyAndLimit(r, resp.Pagination.NextKey, limit)
			req = types.QueryBTCValidatorsRequest{Pagination: pagination}
		}

		if len(btcValsFound) != len(btcValsMap) {
			t.Errorf("Some vals were missed. Got %d while %d were expected", len(btcValsFound), len(btcValsMap))
		}
	})
}

func FuzzBTCDelegations(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup keeper and context
		btclcKeeper := types.NewMockBTCLightClientKeeper(ctrl)
		btccKeeper := types.NewMockBtcCheckpointKeeper(ctrl)
		btccKeeper.EXPECT().GetParams(gomock.Any()).Return(btcctypes.DefaultParams()).AnyTimes()
		keeper, ctx := testkeeper.BTCStakingKeeper(t, btclcKeeper, btccKeeper)

		// Generate a random number of BTC validators
		numBTCVals := datagen.RandomInt(r, 5) + 1
		btcVals := []*types.BTCValidator{}
		for i := uint64(0); i < numBTCVals; i++ {
			btcVal, err := datagen.GenRandomBTCValidator(r)
			require.NoError(t, err)
			keeper.SetBTCValidator(ctx, btcVal)
			btcVals = append(btcVals, btcVal)
		}

		// Generate a random number of BTC delegations under each validator
		startHeight := datagen.RandomInt(r, 100) + 1
		endHeight := datagen.RandomInt(r, 1000) + startHeight + btcctypes.DefaultParams().CheckpointFinalizationTimeout + 1
		numBTCDels := datagen.RandomInt(r, 10) + 1
		pendingBtcDelsMap := make(map[string]*types.BTCDelegation)
		for _, btcVal := range btcVals {
			for j := uint64(0); j < numBTCDels; j++ {
				btcDel, err := datagen.GenRandomBTCDelegation(r, btcVal.BtcPk, startHeight, endHeight, 1)
				require.NoError(t, err)
				if datagen.RandomInt(r, 2) == 1 {
					// remove jury sig in random BTC delegations to make them inactive
					btcDel.JurySig = nil
					pendingBtcDelsMap[btcDel.BtcPk.MarshalHex()] = btcDel
				}
				keeper.SetBTCDelegation(ctx, btcDel)
			}
		}

		babylonHeight := datagen.RandomInt(r, 10) + 1
		ctx = ctx.WithBlockHeight(int64(babylonHeight))
		btclcKeeper.EXPECT().GetTipInfo(gomock.Any()).Return(&btclctypes.BTCHeaderInfo{Height: startHeight}).Times(1)
		keeper.IndexBTCHeight(ctx)

		// assert all pending BTC delegations
		resp, err := keeper.PendingBTCDelegations(ctx, &types.QueryPendingBTCDelegationsRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, len(pendingBtcDelsMap), len(resp.BtcDelegations))
		for _, btcDel := range resp.BtcDelegations {
			_, ok := pendingBtcDelsMap[btcDel.BtcPk.MarshalHex()]
			require.True(t, ok)
		}
	})
}

func FuzzBTCValidatorVotingPowerAtHeight(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		// Setup keeper and context
		keeper, ctx := testkeeper.BTCStakingKeeper(t, nil, nil)

		// random BTC validator
		btcVal, err := datagen.GenRandomBTCValidator(r)
		require.NoError(t, err)
		// add this BTC validator
		keeper.SetBTCValidator(ctx, btcVal)
		// set random voting power at random height
		randomHeight := datagen.RandomInt(r, 100) + 1
		randomPower := datagen.RandomInt(r, 100) + 1
		keeper.SetVotingPower(ctx, btcVal.BtcPk.MustMarshal(), randomHeight, randomPower)

		req := &types.QueryBTCValidatorPowerAtHeightRequest{
			ValBtcPkHex: btcVal.BtcPk.MarshalHex(),
			Height:      randomHeight,
		}
		resp, err := keeper.BTCValidatorPowerAtHeight(ctx, req)
		require.NoError(t, err)
		require.Equal(t, randomPower, resp.VotingPower)
	})
}

func FuzzActiveBTCValidatorsAtHeight(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		// Setup keeper and context
		keeper, ctx := testkeeper.BTCStakingKeeper(t, nil, nil)

		// Generate a random batch of validators
		var btcVals []*types.BTCValidator
		numBTCValsWithVotingPower := datagen.RandomInt(r, 10) + 1
		numBTCVals := numBTCValsWithVotingPower + datagen.RandomInt(r, 10)
		for i := uint64(0); i < numBTCVals; i++ {
			btcVal, err := datagen.GenRandomBTCValidator(r)
			require.NoError(t, err)
			keeper.SetBTCValidator(ctx, btcVal)
			btcVals = append(btcVals, btcVal)
		}

		// For numBTCValsWithVotingPower validators, generate a random number of BTC delegations
		numBTCDels := datagen.RandomInt(r, 10) + 1
		babylonHeight := datagen.RandomInt(r, 10) + 1
		btcValsWithVotingPowerMap := make(map[string]*types.BTCValidator)
		for i := uint64(0); i < numBTCValsWithVotingPower; i++ {
			valBTCPK := btcVals[i].BtcPk
			btcValsWithVotingPowerMap[valBTCPK.MarshalHex()] = btcVals[i]

			var totalVotingPower uint64
			for j := uint64(0); j < numBTCDels; j++ {
				btcDel, err := datagen.GenRandomBTCDelegation(r, valBTCPK, 1, 1000, 1) // timelock period: 1-1000
				require.NoError(t, err)
				keeper.SetBTCDelegation(ctx, btcDel)
				totalVotingPower += btcDel.TotalSat
			}

			keeper.SetVotingPower(ctx, valBTCPK.MustMarshal(), babylonHeight, totalVotingPower)
		}

		// Test nil request
		resp, err := keeper.ActiveBTCValidatorsAtHeight(ctx, nil)
		if resp != nil {
			t.Errorf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Generate a page request with a limit and a nil key
		limit := datagen.RandomInt(r, int(numBTCValsWithVotingPower)) + 1
		pagination := constructRequestWithLimit(r, limit)
		// Generate the initial query
		req := types.QueryActiveBTCValidatorsAtHeightRequest{Height: babylonHeight, Pagination: pagination}
		// Construct a mapping from the btc vals found to a boolean value
		// Will be used later to evaluate whether all the btc vals were returned
		btcValsFound := make(map[string]bool, 0)

		for i := uint64(0); i < numBTCValsWithVotingPower; i += limit {
			resp, err = keeper.ActiveBTCValidatorsAtHeight(ctx, &req)
			if err != nil {
				t.Errorf("Valid request led to an error %s", err)
			}
			if resp == nil {
				t.Fatalf("Valid request led to a nil response")
			}

			for _, val := range resp.BtcValidators {
				// Check if the pk exists in the map
				if _, ok := btcValsWithVotingPowerMap[val.BtcPk.MarshalHex()]; !ok {
					t.Fatalf("rpc returned a val that was not created")
				}
				btcValsFound[val.BtcPk.MarshalHex()] = true
			}

			// Construct the next page request
			pagination = constructRequestWithKeyAndLimit(r, resp.Pagination.NextKey, limit)
			req = types.QueryActiveBTCValidatorsAtHeightRequest{Height: babylonHeight, Pagination: pagination}
		}

		if len(btcValsFound) != len(btcValsWithVotingPowerMap) {
			t.Errorf("Some vals were missed. Got %d while %d were expected", len(btcValsFound), len(btcValsWithVotingPowerMap))
		}
	})
}

func FuzzBTCValidatorDelegations(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Setup keeper and context
		btclcKeeper := types.NewMockBTCLightClientKeeper(ctrl)
		btccKeeper := types.NewMockBtcCheckpointKeeper(ctrl)
		btccKeeper.EXPECT().GetParams(gomock.Any()).Return(btcctypes.DefaultParams()).AnyTimes()
		keeper, ctx := testkeeper.BTCStakingKeeper(t, btclcKeeper, btccKeeper)

		// Generate a btc validator
		btcVal, err := datagen.GenRandomBTCValidator(r)
		require.NoError(t, err)
		keeper.SetBTCValidator(ctx, btcVal)

		startHeight := datagen.RandomInt(r, 100) + 1
		endHeight := datagen.RandomInt(r, 1000) + startHeight + btcctypes.DefaultParams().CheckpointFinalizationTimeout + 1
		// Generate a random number of BTC delegations under this validator
		numBTCDels := datagen.RandomInt(r, 10) + 1
		activeBtcDelsMap := make(map[string]*types.BTCDelegation)
		pendingBtcDelsMap := make(map[string]*types.BTCDelegation)
		for j := uint64(0); j < numBTCDels; j++ {
			btcDel, err := datagen.GenRandomBTCDelegation(r, btcVal.BtcPk, startHeight, endHeight, 1)
			require.NoError(t, err)
			if datagen.RandomInt(r, 2) == 1 {
				// remove jury sig in random BTC delegations to make them inactive
				btcDel.JurySig = nil
				pendingBtcDelsMap[btcDel.BtcPk.MarshalHex()] = btcDel
			} else {
				activeBtcDelsMap[btcDel.BtcPk.MarshalHex()] = btcDel
			}
			keeper.SetBTCDelegation(ctx, btcDel)
		}

		// Test nil request
		resp, err := keeper.BTCValidatorDelegations(ctx, nil)
		require.Nil(t, resp)
		require.Error(t, err)

		babylonHeight := datagen.RandomInt(r, 10) + 1
		ctx = ctx.WithBlockHeight(int64(babylonHeight))
		btclcKeeper.EXPECT().GetTipInfo(gomock.Any()).Return(&btclctypes.BTCHeaderInfo{Height: startHeight}).Times(1)
		keeper.IndexBTCHeight(ctx)

		// Generate a page request with a limit and a nil key
		// query a page of active BTC delegations and assert consistency
		if len(activeBtcDelsMap) > 0 {
			limit := datagen.RandomInt(r, len(activeBtcDelsMap)) + 1
			pagination := constructRequestWithLimit(r, limit)
			// Generate the initial query
			req := types.QueryBTCValidatorDelegationsRequest{
				ValBtcPkHex: btcVal.BtcPk.MarshalHex(),
				DelStatus:   types.BTCDelegationStatus_ACTIVE,
				Pagination:  pagination,
			}
			// Construct a mapping from the btc vals found to a boolean value
			// Will be used later to evaluate whether all the btc vals were returned
			btcDelsFound := make(map[string]bool, 0)

			for i := uint64(0); i < numBTCDels; i += limit {
				resp, err = keeper.BTCValidatorDelegations(ctx, &req)
				require.NoError(t, err)
				require.NotNil(t, resp)
				for _, btcDel := range resp.BtcDelegations {
					require.Equal(t, btcVal.BtcPk, btcDel.ValBtcPk)
					// Check if the pk exists in the map
					_, ok := activeBtcDelsMap[btcDel.BtcPk.MarshalHex()]
					require.True(t, ok)
					btcDelsFound[btcDel.BtcPk.MarshalHex()] = true
				}
				// Construct the next page request
				pagination = constructRequestWithKeyAndLimit(r, resp.Pagination.NextKey, limit)
				req = types.QueryBTCValidatorDelegationsRequest{
					ValBtcPkHex: btcVal.BtcPk.MarshalHex(),
					DelStatus:   types.BTCDelegationStatus_ACTIVE,
					Pagination:  pagination,
				}
			}
			require.Equal(t, len(btcDelsFound), len(activeBtcDelsMap))
		}

		// query a page of pending BTC delegations and assert consistency
		if len(pendingBtcDelsMap) > 0 {
			limit := datagen.RandomInt(r, len(pendingBtcDelsMap)) + 1
			pagination := constructRequestWithLimit(r, limit)
			req := types.QueryBTCValidatorDelegationsRequest{
				ValBtcPkHex: btcVal.BtcPk.MarshalHex(),
				DelStatus:   types.BTCDelegationStatus_PENDING, // only request BTC delegations without jury sig
				Pagination:  pagination,
			}
			pendingBtcDelsFound := make(map[string]bool, 0)

			for i := uint64(0); i < uint64(len(pendingBtcDelsMap)); i += limit {
				resp, err := keeper.BTCValidatorDelegations(ctx, &req)
				require.NoError(t, err)
				require.NotNil(t, resp)
				for _, btcDel := range resp.BtcDelegations {
					require.Equal(t, btcVal.BtcPk, btcDel.ValBtcPk)
					// Check if the pk exists in the map
					_, ok := pendingBtcDelsMap[btcDel.BtcPk.MarshalHex()]
					require.True(t, ok)
					pendingBtcDelsFound[btcDel.BtcPk.MarshalHex()] = true
				}
				// Construct the next page request
				pagination = constructRequestWithKeyAndLimit(r, resp.Pagination.NextKey, limit)
				req = types.QueryBTCValidatorDelegationsRequest{
					ValBtcPkHex: btcVal.BtcPk.MarshalHex(),
					DelStatus:   types.BTCDelegationStatus_PENDING, // only request BTC delegations without jury sig
					Pagination:  pagination,
				}
			}
			require.Equal(t, len(pendingBtcDelsFound), len(pendingBtcDelsMap))
		}
	})
}

// Constructors for PageRequest objects
func constructRequestWithKeyAndLimit(r *rand.Rand, key []byte, limit uint64) *query.PageRequest {
	// If limit is 0, set one randomly
	if limit == 0 {
		limit = uint64(r.Int63() + 1) // Use Int63 instead of Uint64 to avoid overflows
	}
	return &query.PageRequest{
		Key:        key,
		Offset:     0, // only offset or key is set
		Limit:      limit,
		CountTotal: false, // only used when offset is used
		Reverse:    false,
	}
}

func constructRequestWithLimit(r *rand.Rand, limit uint64) *query.PageRequest {
	return constructRequestWithKeyAndLimit(r, nil, limit)
}