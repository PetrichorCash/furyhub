package app

import (
	"encoding/json"
<<<<<<< HEAD
=======

	"github.com/cosmos/cosmos-sdk/codec"
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
)

// The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
<<<<<<< HEAD
func NewDefaultGenesisState() GenesisState {
	encCfg := MakeEncodingConfig()
	return ModuleBasics.DefaultGenesis(encCfg.Marshaler)
=======
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	return ModuleBasics.DefaultGenesis(cdc)
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
}
