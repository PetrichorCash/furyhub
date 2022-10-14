package keeper

import (
	"alliance/x/alliance/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	storeKey   storetypes.StoreKey
	paramstore paramtypes.Subspace
	cdc        codec.BinaryCodec
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ps paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		kt := paramtypes.NewKeyTable().RegisterParamSet(&types.Params{})
		ps = ps.WithKeyTable(kt)
	}

	return Keeper{
		storeKey:   storeKey,
		paramstore: ps,
		cdc:        cdc,
	}
}
