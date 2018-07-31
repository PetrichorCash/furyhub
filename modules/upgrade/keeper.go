package upgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/stake"
	"math"
)

const (
	defaultSwichPeriod     int64 = 57600	// 2 days
)

type Keeper struct {
	storeKey   		sdk.StoreKey
	cdc        		*wire.Codec
	coinKeeper 		bank.Keeper
	// The ValidatorSet to get information about validators
	sk              stake.Keeper
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, ck bank.Keeper, sk stake.Keeper) Keeper {
	keeper := Keeper {
		storeKey:   key,
		cdc:        cdc,
		coinKeeper: ck,
		sk:        sk,
	}
	return keeper
}

func (k Keeper) GetCurrentVersion(ctx sdk.Context) *Version {
	kvStore := ctx.KVStore(k.storeKey)
	versionIDBytes := kvStore.Get(GetCurrentVersionKey())
	if versionIDBytes != nil {
		var versionID int64
		err := k.cdc.UnmarshalBinary(versionIDBytes,&versionID)
		if err != nil {
			panic(err)
		}
		curVersionBytes := kvStore.Get(GetVersionIDKey(versionID))
		if curVersionBytes == nil {
			return nil
		}
		var version Version
		err = k.cdc.UnmarshalBinary(curVersionBytes,&version)
		if err != nil {
			panic(err)
		}
		return &version
	}
	return nil
}

func (k Keeper) AddNewVersion(ctx sdk.Context, version Version) {
	kvStore := ctx.KVStore(k.storeKey)
	curVersion := k.GetCurrentVersion(ctx)

	if curVersion == nil {
		version.Id =0
	} else {
		version.Id = curVersion.Id + 1
		if version.ProposalID == curVersion.ProposalID {
			return
		}
	}
	versionBytes,err := k.cdc.MarshalBinary(version)
	if err != nil {
		panic(err)
	}

	kvStore.Set(GetVersionIDKey(version.Id),versionBytes)

	versionIDBytes, err := k.cdc.MarshalBinary(version.Id)
	if err != nil {
		panic(err)
	}

	kvStore.Set(GetCurrentVersionKey(),versionIDBytes)
	kvStore.Set(GetProposalIDKey(version.ProposalID),versionIDBytes)
	kvStore.Set(GetStartHeightKey(version.Start),versionIDBytes)
}

func (k Keeper) GetVersionByHeight(ctx sdk.Context, blockHeight int64) *Version {
	kvStore := ctx.KVStore(k.storeKey)
	iterator := kvStore.ReverseIterator(GetStartHeightKey(0),GetStartHeightKey(blockHeight+1))
	defer iterator.Close()

	if iterator.Valid() {
		versionIDBytes := iterator.Value()
		if versionIDBytes == nil {
			return nil
		}
		var versionID int64
		err := k.cdc.UnmarshalBinary(versionIDBytes,&versionID)
		if err != nil {
			panic(err)
		}
		versionBytes := kvStore.Get(GetVersionIDKey(versionID))
		if versionBytes == nil  {
			return nil
		}
		var version Version
		err = k.cdc.UnmarshalBinary(versionBytes,&version)
		if err != nil {
			panic(err)
		}
		return &version
	}
	return nil
}

func (k Keeper) GetVersionByProposalId(ctx sdk.Context, proposalId int64) *Version {
	kvStore := ctx.KVStore(k.storeKey)
	versionIDBytes := kvStore.Get(GetProposalIDKey(proposalId))
	if versionIDBytes == nil  {
		return nil
	}
	var versionID int64
	err := k.cdc.UnmarshalBinary(versionIDBytes,&versionID)
	if err != nil {
		panic(err)
	}
	versionBytes := kvStore.Get(GetVersionIDKey(versionID))
	if versionBytes != nil {
		var version Version
		err := k.cdc.UnmarshalBinary(versionBytes,&version)
		if err != nil {
			panic(err)
		}
		return &version
	}
	return nil
}

func (k Keeper) GetVersionByVersionId(ctx sdk.Context, versionId int64) *Version {
	kvStore := ctx.KVStore(k.storeKey)
	curVersionBytes := kvStore.Get(GetVersionIDKey(versionId))
	if curVersionBytes != nil {
		var version Version
		err := k.cdc.UnmarshalBinary(curVersionBytes,&version)
		if err != nil {
			panic(err)
		}
		return &version
	}
	return nil
}

func (k Keeper) GetVersionList(ctx sdk.Context) VersionList {
	kvStore := ctx.KVStore(k.storeKey)
	iterator := kvStore.Iterator(GetVersionIDKey(0),GetVersionIDKey(math.MaxInt64))
	defer iterator.Close()

	var versionList VersionList
	for iterator.Valid() {
		versionBytes := iterator.Value()
		iterator.Next()
		if versionBytes == nil {
			continue
		}
		var version Version
		err := k.cdc.UnmarshalBinary(versionBytes,&version)
		if err != nil {
			panic(err)
		}
		versionList = append(versionList, version)
	}

	return versionList
}

func (k Keeper) GetCurrentProposalID(ctx sdk.Context) int64 {
	kvStore := ctx.KVStore(k.storeKey)
	proposalIdBytes := kvStore.Get(GetCurrentProposalIdKey())
	if proposalIdBytes != nil {
		var proposalId int64
		err := k.cdc.UnmarshalBinary(proposalIdBytes,&proposalId)
		if err != nil {
			panic(err)
		}
		return proposalId
	}
	return -1
}

func (k Keeper) SetCurrentProposalID(ctx sdk.Context, proposalID int64) {
	kvStore := ctx.KVStore(k.storeKey)
	proposalIDBytes, err := k.cdc.MarshalBinary(proposalID)
	if err != nil {
		panic(err)
	}
	kvStore.Set(GetCurrentProposalIdKey(),proposalIDBytes)
}

func (k Keeper) GetMsgTypeInCurrentVersion(ctx sdk.Context, msg sdk.Msg) (string, sdk.Error) {
	currentVersion := k.GetCurrentVersion(ctx)
	return currentVersion.getMsgType(msg)
}

func (k Keeper) SetSwitch(ctx sdk.Context ,propsalID int64, address sdk.AccAddress,cmsg MsgSwitch) {
	kvStore := ctx.KVStore(k.storeKey)
	cmsgBytes,err := k.cdc.MarshalBinary(cmsg)
	if err != nil {
		panic(err)
	}
	kvStore.Set(GetSwitchKey(propsalID,address),cmsgBytes)
}

func (k Keeper) GetSwitch(ctx sdk.Context ,propsalID int64, address sdk.AccAddress) (MsgSwitch, bool) {
	kvStore := ctx.KVStore(k.storeKey)
	cmsgBytes := kvStore.Get(GetSwitchKey(propsalID,address))
	if cmsgBytes != nil {
		var cmsg MsgSwitch
		err := k.cdc.UnmarshalBinary(cmsgBytes,&cmsg)
		if err != nil {
			panic(err)
		}
		return cmsg, true
	}
	return MsgSwitch{}, false
}

func (k Keeper) SetCurrentProposalAcceptHeight(ctx sdk.Context, height int64) {
	kvStore := ctx.KVStore(k.storeKey)
	heightBytes, err := k.cdc.MarshalBinary(height)
	if err != nil {
		panic(err)
	}
	kvStore.Set(GetCurrentProposalAcceptHeightKey(),heightBytes)
}

func (k Keeper) GetCurrentProposalAcceptHeight(ctx sdk.Context) int64 {
	kvStore := ctx.KVStore(k.storeKey)
	proposalAcceptHeightBytes := kvStore.Get(GetCurrentProposalAcceptHeightKey())
	if proposalAcceptHeightBytes != nil {
		var proposalAcceptHeight int64
		err := k.cdc.UnmarshalBinary(proposalAcceptHeightBytes, &proposalAcceptHeight)
		if err != nil {
			panic(err)
		}
		return proposalAcceptHeight
	}
	return -1
}

func (k Keeper) SetDoingSwitch(ctx sdk.Context, doing bool) {
	kvStore := ctx.KVStore(k.storeKey)

	var bytes []byte
	if doing {
		bytes = []byte{byte(1)}
	} else {
		bytes = []byte{byte(0)}
	}
	kvStore.Set(GetDoingSwitchKey(), bytes)
}

func (k Keeper) GetDoingSwitch(ctx sdk.Context) (bool) {
	kvStore := ctx.KVStore(k.storeKey)

	bytes := kvStore.Get(GetDoingSwitchKey())
	if len(bytes) == 1 {
		return bytes[0] == byte(1)
	}

	return false
}

func (k Keeper) DoSwitchBegin(ctx sdk.Context) {
	k.SetDoingSwitch(ctx, true)

}
