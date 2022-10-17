package keeper_test

import (
	test_helpers "alliance/app"
	"alliance/x/alliance/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
)

var ALLIANCE_TOKEN_DENOM = "alliance"
var ALLIANCE_2_TOKEN_DENOM = "alliance2"

func TestDelegation(t *testing.T) {
	app, ctx := createTestContext(t)
	app.AllianceKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.Params{
			RewardDelayTime: time.Duration(1000000),
		},
		Assets: []types.AllianceAsset{
			{
				Denom:        ALLIANCE_TOKEN_DENOM,
				RewardWeight: sdk.NewDec(2),
				TakeRate:     sdk.NewDec(0),
				TotalShares:  sdk.NewDec(0),
				TotalTokens:  sdk.ZeroInt(),
			},
			{
				Denom:        ALLIANCE_2_TOKEN_DENOM,
				RewardWeight: sdk.NewDec(10),
				TakeRate:     sdk.NewDec(0),
				TotalShares:  sdk.NewDec(0),
				TotalTokens:  sdk.ZeroInt(),
			},
		},
	})
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	require.Len(t, delegations, 1)

	// All the addresses needed
	delAddr, err := sdk.AccAddressFromBech32(delegations[0].DelegatorAddress)
	require.NoError(t, err)
	valAddr, err := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	require.NoError(t, err)
	val, _ := app.StakingKeeper.GetValidator(ctx, valAddr)
	moduleAddr := app.AccountKeeper.GetModuleAddress(types.ModuleName)

	// Mint alliance tokens
	err = app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(2000_000))))
	require.NoError(t, err)
	err = app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(ALLIANCE_2_TOKEN_DENOM, sdk.NewInt(2000_000))))
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(2000_000))))
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(ALLIANCE_2_TOKEN_DENOM, sdk.NewInt(2000_000))))
	require.NoError(t, err)

	// Delegate
	_, err = app.AllianceKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, err)

	// Check delegation in staking module
	delegations = app.StakingKeeper.GetAllDelegations(ctx)
	require.Len(t, delegations, 2)
	i := slices.IndexFunc(delegations, func(d stakingtypes.Delegation) bool {
		return d.DelegatorAddress == moduleAddr.String()
	})
	newDelegation := delegations[i]
	require.Equal(t, stakingtypes.Delegation{
		DelegatorAddress: moduleAddr.String(),
		ValidatorAddress: valAddr.String(),
		Shares:           sdk.NewDec(2),
	}, newDelegation)

	// Check delegation in alliance module
	allianceDelegation, found := app.AllianceKeeper.GetDelegation(ctx, delAddr, val, ALLIANCE_TOKEN_DENOM)
	require.True(t, found)
	require.Equal(t, types.Delegation{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: val.OperatorAddress,
		Denom:            ALLIANCE_TOKEN_DENOM,
		Shares:           sdk.NewDec(1000_000),
	}, allianceDelegation)

	// Delegate with same denom again
	_, err = app.AllianceKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, err)
	// Check delegation in alliance module
	allianceDelegation, found = app.AllianceKeeper.GetDelegation(ctx, delAddr, val, ALLIANCE_TOKEN_DENOM)
	require.True(t, found)
	require.Equal(t, types.Delegation{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: val.OperatorAddress,
		Denom:            ALLIANCE_TOKEN_DENOM,
		Shares:           sdk.NewDec(2000_000),
	}, allianceDelegation)

	// Delegate with another denom
	_, err = app.AllianceKeeper.Delegate(ctx, delAddr, val, sdk.NewCoin(ALLIANCE_2_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.NoError(t, err)

	// Check delegation in staking module
	delegations = app.StakingKeeper.GetAllDelegations(ctx)
	require.Len(t, delegations, 2)
	i = slices.IndexFunc(delegations, func(d stakingtypes.Delegation) bool {
		return d.DelegatorAddress == moduleAddr.String()
	})
	newDelegation = delegations[i]
	require.Equal(t, stakingtypes.Delegation{
		DelegatorAddress: moduleAddr.String(),
		ValidatorAddress: valAddr.String(),
		Shares:           sdk.NewDec(14),
	}, newDelegation)
}

// TODO: test using unsupported denoms

func TestRedelegation(t *testing.T) {
	app, ctx := createTestContext(t)
	app.AllianceKeeper.InitGenesis(ctx, &types.GenesisState{
		Params: types.Params{
			RewardDelayTime: time.Duration(1000000),
		},
		Assets: []types.AllianceAsset{
			{
				Denom:        ALLIANCE_TOKEN_DENOM,
				RewardWeight: sdk.NewDec(2),
				TakeRate:     sdk.NewDec(0),
				TotalShares:  sdk.NewDec(0),
				TotalTokens:  sdk.ZeroInt(),
			},
			{
				Denom:        ALLIANCE_2_TOKEN_DENOM,
				RewardWeight: sdk.NewDec(10),
				TakeRate:     sdk.NewDec(0),
				TotalShares:  sdk.NewDec(0),
				TotalTokens:  sdk.ZeroInt(),
			},
		},
	})

	// Get all the addresses needed for the test
	moduleAddr := app.AccountKeeper.GetModuleAddress(types.ModuleName)
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	valAddr1, err := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	require.NoError(t, err)
	val1, found := app.StakingKeeper.GetValidator(ctx, valAddr1)
	require.True(t, found)
	addrs := test_helpers.AddTestAddrsIncremental(app, ctx, 3, sdk.NewCoins(
		sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(1000_000)),
		sdk.NewCoin(ALLIANCE_2_TOKEN_DENOM, sdk.NewInt(1000_000)),
	))
	valAddr2 := sdk.ValAddress(addrs[0])
	val2 := teststaking.NewValidator(t, valAddr2, test_helpers.CreateTestPubKeys(1)[0])
	test_helpers.RegisterNewValidator(t, app, ctx, val2)
	delAddr1 := addrs[1]
	delAddr2 := addrs[2]

	// Mint alliance tokens
	err = app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(2000_000))))
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr1, sdk.NewCoins(sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(1000_000))))
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, delAddr2, sdk.NewCoins(sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(1000_000))))
	require.NoError(t, err)

	// First delegate to validator 1
	_, err = app.AllianceKeeper.Delegate(ctx, delAddr1, val1, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	require.NoError(t, err)

	// Then redelegate to validator2
	_, err = app.AllianceKeeper.Redelegate(ctx, delAddr1, val1, val2, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	require.NoError(t, err)

	delegations = app.StakingKeeper.GetAllDelegations(ctx)
	for _, d := range delegations {
		if d.DelegatorAddress == moduleAddr.String() {
			require.Equal(t, stakingtypes.Delegation{
				DelegatorAddress: moduleAddr.String(),
				ValidatorAddress: val2.OperatorAddress,
				Shares:           sdk.NewDec(1_000_000),
			}, d)
		}
	}

	// Check if there is a re-delegation event stored
	iter := app.AllianceKeeper.IterateRedelegationsByDelegator(ctx, delAddr1)
	require.True(t, iter.Valid())
	for ; iter.Valid(); iter.Next() {
		var redelegation types.Redelegation
		app.AppCodec().MustUnmarshal(iter.Value(), &redelegation)
		require.Equal(t, types.Redelegation{
			DelegatorAddress:    delAddr1.String(),
			SrcValidatorAddress: valAddr1.String(),
			DstValidatorAddress: valAddr2.String(),
			Balance:             sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)),
		}, redelegation)
	}

	// Check if the delegation objects are correct
	_, found = app.AllianceKeeper.GetDelegation(ctx, delAddr1, val1, ALLIANCE_TOKEN_DENOM)
	require.False(t, found)
	dstDelegation, found := app.AllianceKeeper.GetDelegation(ctx, delAddr1, val2, ALLIANCE_TOKEN_DENOM)
	require.Equal(t, types.Delegation{
		DelegatorAddress: delAddr1.String(),
		ValidatorAddress: val2.GetOperator().String(),
		Denom:            ALLIANCE_TOKEN_DENOM,
		Shares:           sdk.NewDec(500_000),
	}, dstDelegation)
	require.True(t, found)

	// Should fail when re-delegating back to validator1
	// Same user who re-delegated to from 1 -> 2 cannot re-re-delegate from 2 -> X
	_, err = app.AllianceKeeper.Redelegate(ctx, delAddr1, val2, val1, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	require.Error(t, err)

	// Another user tries to re-delegate without having an initial delegation but fails
	_, err = app.AllianceKeeper.Redelegate(ctx, delAddr2, val2, val1, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	require.Error(t, err)

	// User then delegates to validator2
	_, err = app.AllianceKeeper.Delegate(ctx, delAddr2, val2, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	require.NoError(t, err)

	// Then redelegate to validator1 with more than what was delegated but fails
	_, err = app.AllianceKeeper.Redelegate(ctx, delAddr2, val2, val1, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(1000_000)))
	require.Error(t, err)

	// Then redelegate to validator1 correctly
	_, err = app.AllianceKeeper.Redelegate(ctx, delAddr2, val2, val1, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	// Should pass since we removed the re-delegate attempt on x/staking that prevents this
	require.NoError(t, err)

	// Immediately calling complete re-delegation should do nothing
	deleted := app.AllianceKeeper.CompleteRedelegations(ctx)
	require.Equal(t, 0, deleted)

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(app.StakingKeeper.UnbondingTime(ctx)).Add(time.Minute))
	// Calling after re-delegation has matured will delete it from the store
	deleted = app.AllianceKeeper.CompleteRedelegations(ctx)
	require.Equal(t, 2, deleted)

	// There shouldn't be any more delegations in the store
	iter = app.AllianceKeeper.IterateRedelegationsByDelegator(ctx, delAddr1)
	require.False(t, iter.Valid())
	iter = app.AllianceKeeper.IterateRedelegationsByDelegator(ctx, delAddr2)
	require.False(t, iter.Valid())

	// Another user first delegates to validator2
	_, err = app.AllianceKeeper.Delegate(ctx, delAddr2, val2, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))
	require.NoError(t, err)

	// Then redelegate to validator1
	_, err = app.AllianceKeeper.Redelegate(ctx, delAddr2, val2, val1, sdk.NewCoin(ALLIANCE_TOKEN_DENOM, sdk.NewInt(500_000)))

	// Should pass since we removed the re-delegate attempt on x/staking that prevents this
	require.NoError(t, err)
}
