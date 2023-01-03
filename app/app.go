package app

import (
	"fmt"
	"io"
<<<<<<< HEAD
	"os"
	"path/filepath"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

=======
	"net/http"
	"os"
	"path/filepath"

>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
<<<<<<< HEAD
	"github.com/cosmos/cosmos-sdk/store/streaming"
=======
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
<<<<<<< HEAD
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
=======
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
<<<<<<< HEAD
=======
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
<<<<<<< HEAD
	sdkupgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

=======
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"
	icahost "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/ibc-go/v5/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v5/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v5/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v5/modules/core/02-client"
<<<<<<< HEAD
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"

	"github.com/irisnet/irismod/modules/coinswap"
	coinswapkeeper "github.com/irisnet/irismod/modules/coinswap/keeper"
	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"
	"github.com/irisnet/irismod/modules/htlc"
	htlckeeper "github.com/irisnet/irismod/modules/htlc/keeper"
	htlctypes "github.com/irisnet/irismod/modules/htlc/types"
	"github.com/irisnet/irismod/modules/mt"
	mtkeeper "github.com/irisnet/irismod/modules/mt/keeper"
	mttypes "github.com/irisnet/irismod/modules/mt/types"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"
	nftmodule "github.com/irisnet/irismod/modules/nft/module"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
	"github.com/irisnet/irismod/modules/oracle"
	oraclekeeper "github.com/irisnet/irismod/modules/oracle/keeper"
	oracletypes "github.com/irisnet/irismod/modules/oracle/types"
	"github.com/irisnet/irismod/modules/random"
	randomkeeper "github.com/irisnet/irismod/modules/random/keeper"
	randomtypes "github.com/irisnet/irismod/modules/random/types"
	"github.com/irisnet/irismod/modules/record"
	recordkeeper "github.com/irisnet/irismod/modules/record/keeper"
	recordtypes "github.com/irisnet/irismod/modules/record/types"
	"github.com/irisnet/irismod/modules/service"
	servicekeeper "github.com/irisnet/irismod/modules/service/keeper"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
	"github.com/irisnet/irismod/modules/token"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
	tokentypes "github.com/irisnet/irismod/modules/token/types"

	"github.com/irisnet/irishub/address"
	irishubante "github.com/irisnet/irishub/ante"
	irisappparams "github.com/irisnet/irishub/app/params"
	"github.com/irisnet/irishub/lite"
	"github.com/irisnet/irishub/modules/guardian"
	guardiankeeper "github.com/irisnet/irishub/modules/guardian/keeper"
	guardiantypes "github.com/irisnet/irishub/modules/guardian/types"
	"github.com/irisnet/irishub/modules/mint"
	mintkeeper "github.com/irisnet/irishub/modules/mint/keeper"
	minttypes "github.com/irisnet/irishub/modules/mint/types"

	"github.com/irisnet/irismod/modules/farm"
	farmkeeper "github.com/irisnet/irismod/modules/farm/keeper"
	farmtypes "github.com/irisnet/irismod/modules/farm/types"

	tibcmttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer"
	tibcmttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/keeper"
	tibcmttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/types"
	tibcnfttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibcnfttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/keeper"
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client"
	tibcclienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	tibcrouting "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing"
	tibcroutingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
	tibckeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
)

const appName = "PetriApp"
=======
	ibcclientclient "github.com/cosmos/ibc-go/v5/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"
	custombankmodule "github.com/petrinetwork/petrichor/custom/bank"
	custombankkeeper "github.com/petrinetwork/petrichor/custom/bank/keeper"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/ignite/cli/ignite/pkg/openapiconsole"

	"github.com/petrinetwork/petrichor/docs"

	petrichormodule "github.com/petrinetwork/petrichor/x/petrichor"
	petrichormoduleclient "github.com/petrinetwork/petrichor/x/petrichor/client"
	petrichormodulekeeper "github.com/petrinetwork/petrichor/x/petrichor/keeper"
	petrichormoduletypes "github.com/petrinetwork/petrichor/x/petrichor/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

const (
	AccountAddressPrefix = "petri"
	Name                 = "petrichor"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		petrichormoduleclient.CreatePetrichorProposalHandler,
		petrichormoduleclient.UpdatePetrichorProposalHandler,
		petrichormoduleclient.DeletePetrichorProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

<<<<<<< HEAD
	// Denominations can be 3 ~ 128 characters long and support letters, followed by either
	// a letter, a number, ('-'), or a separator ('/').
	// overwite sdk reDnmString
	reDnmString = `[a-zA-Z][a-zA-Z0-9/-]{2,127}`

=======
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
<<<<<<< HEAD
		//authzmodule.AppModuleBasic{},
=======
		authzmodule.AppModuleBasic{},
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
<<<<<<< HEAD
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
				distrclient.ProposalHandler,
				upgradeclient.LegacyProposalHandler,
				upgradeclient.LegacyCancelProposalHandler,
				tibcclient.CreateClientProposalHandler,
				tibcclient.UpgradeClientProposalHandler,
				tibcclient.RegisterRelayerProposalHandler,
				tibcrouting.SetRoutingRulesProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
=======
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
<<<<<<< HEAD
		vesting.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		ica.AppModuleBasic{},

		guardian.AppModuleBasic{},
		token.AppModuleBasic{},
		record.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		htlc.AppModuleBasic{},
		coinswap.AppModuleBasic{},
		service.AppModuleBasic{},
		oracle.AppModuleBasic{},
		random.AppModuleBasic{},
		farm.AppModuleBasic{},
		tibc.AppModuleBasic{},
		tibcnfttransfer.AppModuleBasic{},
		tibcmttransfer.AppModuleBasic{},
		mt.AppModuleBasic{},
=======
		ica.AppModuleBasic{},
		vesting.AppModuleBasic{},
		petrichormodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	)

	// module account permissions
	maccPerms = map[string][]string{
<<<<<<< HEAD
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		tokentypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
		htlctypes.ModuleName:           {authtypes.Minter, authtypes.Burner},
		coinswaptypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		servicetypes.DepositAccName:    {authtypes.Burner},
		servicetypes.RequestAccName:    nil,
		servicetypes.FeeCollectorName:  {authtypes.Burner},
		farmtypes.ModuleName:           {authtypes.Burner},
		farmtypes.RewardCollector:      nil,
		farmtypes.EscrowCollector:      nil,
		tibcnfttypes.ModuleName:        nil,
		tibcmttypes.ModuleName:         nil,
		nfttypes.ModuleName:            nil,
		icatypes.ModuleName:            nil,
	}

	nativeToken tokentypes.Token
)

var (
	_ simapp.App              = (*PetriApp)(nil)
	_ servertypes.Application = (*PetriApp)(nil)
)

// PetriApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type PetriApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
=======
		authtypes.FeeCollectorName:          nil,
		distrtypes.ModuleName:               nil,
		icatypes.ModuleName:                 nil,
		minttypes.ModuleName:                {authtypes.Minter},
		stakingtypes.BondedPoolName:         {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:      {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:                 {authtypes.Burner},
		ibctransfertypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
		petrichormoduletypes.ModuleName:      {authtypes.Minter, authtypes.Burner},
		petrichormoduletypes.RewardsPoolName: nil,
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}
)

var (
	_ cosmoscmd.App           = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

<<<<<<< HEAD
	// cosmos
	FeeGrantKeeper   feegrantkeeper.Keeper
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
=======
	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       custombankkeeper.Keeper
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
<<<<<<< HEAD
	EvidenceKeeper   evidencekeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper

	//ibc
	IBCKeeper         *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCTransferKeeper ibctransferkeeper.Keeper
	ICAHostKeeper     icahostkeeper.Keeper

	// make scoped keepers public for test purposes
	scopedIBCKeeper         capabilitykeeper.ScopedKeeper
	scopedTransferKeeper    capabilitykeeper.ScopedKeeper
	scopedIBCMockKeeper     capabilitykeeper.ScopedKeeper
	scopedNFTTransferKeeper capabilitykeeper.ScopedKeeper
	// tibc
	scopedTIBCKeeper     capabilitykeeper.ScopedKeeper
	scopedTIBCMockKeeper capabilitykeeper.ScopedKeeper

	GuardianKeeper        guardiankeeper.Keeper
	TokenKeeper           tokenkeeper.Keeper
	RecordKeeper          recordkeeper.Keeper
	NFTKeeper             nftkeeper.Keeper
	MTKeeper              mtkeeper.Keeper
	HTLCKeeper            htlckeeper.Keeper
	CoinswapKeeper        coinswapkeeper.Keeper
	ServiceKeeper         servicekeeper.Keeper
	OracleKeeper          oraclekeeper.Keeper
	RandomKeeper          randomkeeper.Keeper
	FarmKeeper            farmkeeper.Keeper
	TIBCKeeper            *tibckeeper.Keeper
	TIBCNFTTransferKeeper tibcnfttransferkeeper.Keeper
	TIBCMTTransferKeeper  tibcmttransferkeeper.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager
}

func init() {
	// set bech32 prefix
	address.ConfigureBech32Prefix()

	// set coin denom regexs
	sdk.SetCoinDenomRegex(DefaultCoinDenomRegex)

	nativeToken = tokentypes.Token{
		Symbol:        "petri",
		Name:          "Petrihub staking token",
		Scale:         6,
		MinUnit:       "upetri",
		InitialSupply: 48000000000000,
		MaxSupply:     420000000000000,
		Mintable:      true,
		Owner:         sdk.AccAddress(crypto.AddressHash([]byte(tokentypes.ModuleName))).String(),
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".petri")
	owner, err := sdk.AccAddressFromBech32(nativeToken.Owner)
	if err != nil {
		panic(err)
	}

	tokentypes.SetNativeToken(
		nativeToken.Symbol,
		nativeToken.Name,
		nativeToken.MinUnit,
		nativeToken.Scale,
		nativeToken.InitialSupply,
		nativeToken.MaxSupply,
		nativeToken.Mintable,
		owner,
	)
}

// DefaultCoinDenomRegex returns the default regex string
func DefaultCoinDenomRegex() string {
	return reDnmString
}

// NewIrisApp returns a reference to an initialized PetriApp.
func NewIrisApp(
=======
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	ICAHostKeeper    icahostkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	GroupKeeper      groupkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper  capabilitykeeper.ScopedKeeper

	PetrichorKeeper petrichormodulekeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

// New returns a reference to an initialized blockchain app
func New(
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
<<<<<<< HEAD
	encodingConfig irisappparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *PetriApp {

	// TODO: Remove cdc in favor of appCodec once all modules are migrated.
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
=======
	encodingConfig cosmoscmd.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) cosmoscmd.App {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
<<<<<<< HEAD
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		guardiantypes.StoreKey, tokentypes.StoreKey, nfttypes.StoreKey, htlctypes.StoreKey, recordtypes.StoreKey,
		coinswaptypes.StoreKey, servicetypes.StoreKey, oracletypes.StoreKey, randomtypes.StoreKey,
		farmtypes.StoreKey, feegrant.StoreKey, tibchost.StoreKey, tibcnfttypes.StoreKey, tibcmttypes.StoreKey, mttypes.StoreKey,
		authzkeeper.StoreKey, icahosttypes.StoreKey,
=======
		authtypes.StoreKey, authz.ModuleName, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey, govtypes.StoreKey,
		paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey, evidencetypes.StoreKey,
		ibctransfertypes.StoreKey, icahosttypes.StoreKey, capabilitytypes.StoreKey, group.StoreKey,
		petrichormoduletypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

<<<<<<< HEAD
	// configure state listening capabilities using AppOptions
	// we are doing nothing with the returned streamingServices and waitGroup in this case
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, keys); err != nil {
		tmos.Exit(err.Error())
	}

	app := &PetriApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
=======
	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

<<<<<<< HEAD
	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	// set the BaseApp's parameter store
	bApp.SetParamStore(
		app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)

=======
	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
<<<<<<< HEAD
<<<<<<< HEAD
		sdk.Bech32MainPrefix,
=======
		address.Bech32PrefixAccAddr,
>>>>>>> d6b39449b465b07920690ecfd91a0275d2a2af37
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
=======
		sdk.Bech32PrefixAccAddr,
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authz.ModuleName],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	app.BankKeeper = custombankkeeper.NewBaseKeeper(
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedModuleAccountAddrs(),
	)

	stakingKeeper := stakingkeeper.NewKeeper(
<<<<<<< HEAD
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName),
=======
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
<<<<<<< HEAD
=======
		&stakingKeeper,
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

<<<<<<< HEAD
	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// set the governance module account as the authority for conducting upgrades
	// UpgradeKeeper must be created before IBCKeeper
=======
	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	app.GroupKeeper = groupkeeper.NewKeeper(
		keys[group.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		groupConfig,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

<<<<<<< HEAD
	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	scopedTIBCKeeper := app.CapabilityKeeper.ScopeToModule(tibchost.ModuleName)
	// UpgradeKeeper must be created before IBCKeeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
=======
	app.PetrichorKeeper = petrichormodulekeeper.NewKeeper(
		appCodec,
		keys[petrichormoduletypes.StoreKey],
		app.GetSubspace(petrichormoduletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		app.DistrKeeper,
	)

	app.BankKeeper.RegisterKeepers(app.PetrichorKeeper, &stakingKeeper)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks(), app.PetrichorKeeper.StakingHooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey],
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

<<<<<<< HEAD
=======
	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.TransferKeeper)

>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)
	icaModule := ica.NewAppModule(nil, &app.ICAHostKeeper)
	icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

<<<<<<< HEAD
	// register the proposal types
	app.TIBCKeeper = tibckeeper.NewKeeper(
		appCodec,
		keys[tibchost.StoreKey],
		app.GetSubspace(tibchost.ModuleName), app.StakingKeeper,
	)

	app.NFTKeeper = nftkeeper.NewKeeper(
		appCodec,
		keys[nfttypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.TIBCNFTTransferKeeper = tibcnfttransferkeeper.NewKeeper(
		appCodec,
		keys[tibcnfttypes.StoreKey],
		app.GetSubspace(tibcnfttypes.ModuleName),
		app.AccountKeeper,
		nftkeeper.NewLegacyKeeper(app.NFTKeeper),
		app.TIBCKeeper.PacketKeeper,
		app.TIBCKeeper.ClientKeeper,
	)

	app.MTKeeper = mtkeeper.NewKeeper(
		appCodec, keys[mttypes.StoreKey],
	)

	app.TIBCMTTransferKeeper = tibcmttransferkeeper.NewKeeper(
		appCodec,
		keys[tibcnfttypes.StoreKey],
		app.GetSubspace(tibcnfttypes.ModuleName),
		app.AccountKeeper, app.MTKeeper,
		app.TIBCKeeper.PacketKeeper,
		app.TIBCKeeper.ClientKeeper,
	)

	app.IBCTransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.IBCTransferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.IBCTransferKeeper)

	// routerModule := router.NewAppModule(app.RouterKeeper, transferIBCModule)
	// create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferIBCModule).
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule)
	app.IBCKeeper.SetRouter(ibcRouter)

	nfttransferModule := tibcnfttransfer.NewAppModule(app.TIBCNFTTransferKeeper)
	mttransferModule := tibcmttransfer.NewAppModule(app.TIBCMTTransferKeeper)

	tibcRouter := tibcroutingtypes.NewRouter()
	tibcRouter.AddRoute(tibcnfttypes.ModuleName, nfttransferModule).
		AddRoute(tibcmttypes.ModuleName, mttransferModule)
	app.TIBCKeeper.SetRouter(tibcRouter)

	// create evidence keeper with router
=======
	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

<<<<<<< HEAD
	app.GuardianKeeper = guardiankeeper.NewKeeper(
		appCodec,
		keys[guardiantypes.StoreKey],
	)

	app.TokenKeeper = tokenkeeper.NewKeeper(
		appCodec,
		keys[tokentypes.StoreKey],
		app.GetSubspace(tokentypes.ModuleName),
		app.BankKeeper,
		app.ModuleAccountAddrs(),
		authtypes.FeeCollectorName,
	)

	app.RecordKeeper = recordkeeper.NewKeeper(
		appCodec,
		keys[recordtypes.StoreKey],
	)

	app.HTLCKeeper = htlckeeper.NewKeeper(
		appCodec, keys[htlctypes.StoreKey],
		app.GetSubspace(htlctypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.ModuleAccountAddrs(),
	)

	app.CoinswapKeeper = coinswapkeeper.NewKeeper(
		appCodec,
		keys[coinswaptypes.StoreKey],
		app.GetSubspace(coinswaptypes.ModuleName),
		app.BankKeeper,
		app.AccountKeeper,
		app.ModuleAccountAddrs(),
		authtypes.FeeCollectorName,
	)

	app.ServiceKeeper = servicekeeper.NewKeeper(
		appCodec,
		keys[servicetypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(servicetypes.ModuleName),
		app.ModuleAccountAddrs(),
		servicetypes.FeeCollectorName,
	)

	app.OracleKeeper = oraclekeeper.NewKeeper(
		appCodec,
		keys[oracletypes.StoreKey],
		app.GetSubspace(oracletypes.ModuleName),
		app.ServiceKeeper,
	)

	app.RandomKeeper = randomkeeper.NewKeeper(
		appCodec,
		keys[randomtypes.StoreKey],
		app.BankKeeper,
		app.ServiceKeeper,
	)

	app.FarmKeeper = farmkeeper.NewKeeper(appCodec,
		keys[farmtypes.StoreKey],
		app.BankKeeper,
		app.AccountKeeper,
		app.DistrKeeper,
		&app.GovKeeper,
		app.CoinswapKeeper.ValidatePool,
		app.GetSubspace(farmtypes.ModuleName),
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
	)

	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibchost.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(tibcclienttypes.RouterKey, tibcclient.NewClientProposalHandler(app.TIBCKeeper.ClientKeeper)).
		AddRoute(tibcroutingtypes.RouterKey, tibcrouting.NewSetRoutingProposalHandler(app.TIBCKeeper.RoutingKeeper)).
		AddRoute(farmtypes.RouterKey, farm.NewCommunityPoolCreateFarmProposalHandler(app.FarmKeeper))

	govConfig := govtypes.DefaultConfig()

=======
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(petrichormoduletypes.RouterKey, petrichormodule.NewPetrichorProposalHandler(app.PetrichorKeeper))
	govConfig := govtypes.DefaultConfig()
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
<<<<<<< HEAD
		app.StakingKeeper,
=======
		&stakingKeeper,
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		govRouter,
		app.MsgServiceRouter(),
		govConfig,
	)

<<<<<<< HEAD
	govHooks := govtypes.NewMultiGovHooks(farmkeeper.NewGovHook(app.FarmKeeper))
	app.GovKeeper.SetHooks(govHooks)

	/****  Module Options ****/
	var skipGenesisInvariants = false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
=======
	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, transferIBCModule)
	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
<<<<<<< HEAD
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper),
=======
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		custombankmodule.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		groupmodule.NewAppModule(appCodec, app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, minttypes.DefaultInflationCalculationFn),
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
<<<<<<< HEAD
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper), tibc.NewAppModule(app.TIBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		icaModule,
		nfttransferModule,
		mttransferModule,
		guardian.NewAppModule(appCodec, app.GuardianKeeper),
		token.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		record.NewAppModule(appCodec, app.RecordKeeper, app.AccountKeeper, app.BankKeeper),
		nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper),
		mt.NewAppModule(appCodec, app.MTKeeper, app.AccountKeeper, app.BankKeeper),
		htlc.NewAppModule(appCodec, app.HTLCKeeper, app.AccountKeeper, app.BankKeeper),
		coinswap.NewAppModule(appCodec, app.CoinswapKeeper, app.AccountKeeper, app.BankKeeper),
		service.NewAppModule(appCodec, app.ServiceKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		random.NewAppModule(appCodec, app.RandomKeeper, app.AccountKeeper, app.BankKeeper),
		farm.NewAppModule(appCodec, app.FarmKeeper, app.AccountKeeper, app.BankKeeper),
=======
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		icaModule,
		petrichormodule.NewAppModule(appCodec, app.PetrichorKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		// this line is used by starport scaffolding # stargate/app/appModule
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
<<<<<<< HEAD
		//sdk module
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		vestingtypes.ModuleName,

		//self module
		tokentypes.ModuleName,
		tibchost.ModuleName,
		ibctransfertypes.ModuleName,
		nfttypes.ModuleName,
		htlctypes.ModuleName,
		recordtypes.ModuleName,
		coinswaptypes.ModuleName,
		servicetypes.ModuleName,
		oracletypes.ModuleName,
		randomtypes.ModuleName,
		farmtypes.ModuleName,
		mttypes.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		guardiantypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		//sdk module
		upgradetypes.ModuleName,
=======
		// upgrades should be run first
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		petrichormoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
<<<<<<< HEAD
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
=======
		slashingtypes.ModuleName,
		minttypes.ModuleName,
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
<<<<<<< HEAD
		paramstypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		vestingtypes.ModuleName,

		//self module
		tokentypes.ModuleName,
		tibchost.ModuleName,
		ibctransfertypes.ModuleName,
		nfttypes.ModuleName,
		htlctypes.ModuleName,
		recordtypes.ModuleName,
		coinswaptypes.ModuleName,
		servicetypes.ModuleName,
		oracletypes.ModuleName,
		randomtypes.ModuleName,
		farmtypes.ModuleName,
		mttypes.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		guardiantypes.ModuleName,
=======
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		petrichormoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
<<<<<<< HEAD
		//sdk module
		upgradetypes.ModuleName,
=======
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
<<<<<<< HEAD
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		vestingtypes.ModuleName,

		//self module
		tokentypes.ModuleName,
		tibchost.ModuleName,
		ibctransfertypes.ModuleName,
		nfttypes.ModuleName,
		htlctypes.ModuleName,
		recordtypes.ModuleName,
		coinswaptypes.ModuleName,
		servicetypes.ModuleName,
		oracletypes.ModuleName,
		randomtypes.ModuleName,
		farmtypes.ModuleName,
		mttypes.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		guardiantypes.ModuleName,
	)

	cfg := module.NewConfigurator(appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(cfg)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
=======
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		petrichormoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	// Uncomment if you want to set a custom migration order here.
	// app.mm.SetOrderMigrations(custom order)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, minttypes.DefaultInflationCalculationFn),
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
<<<<<<< HEAD
		evidence.NewAppModule(app.EvidenceKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),

		ibc.NewAppModule(app.IBCKeeper),
		transferModule,

		guardian.NewAppModule(appCodec, app.GuardianKeeper),
		token.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		record.NewAppModule(appCodec, app.RecordKeeper, app.AccountKeeper, app.BankKeeper),
		nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper),
		htlc.NewAppModule(appCodec, app.HTLCKeeper, app.AccountKeeper, app.BankKeeper),
		coinswap.NewAppModule(appCodec, app.CoinswapKeeper, app.AccountKeeper, app.BankKeeper),
		service.NewAppModule(appCodec, app.ServiceKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		random.NewAppModule(appCodec, app.RandomKeeper, app.AccountKeeper, app.BankKeeper),
		farm.NewAppModule(appCodec, app.FarmKeeper, app.AccountKeeper, app.BankKeeper),
		tibc.NewAppModule(app.TIBCKeeper),
	)

=======
		groupmodule.NewAppModule(appCodec, app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		petrichormodule.NewAppModule(appCodec, app.PetrichorKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		// this line is used by starport scaffolding # stargate/app/appModule
	)
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

<<<<<<< HEAD
	anteHandler, err := irishubante.NewAnteHandler(
		irishubante.HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			BankKeeper:           app.BankKeeper,
			TokenKeeper:          app.TokenKeeper,
			OracleKeeper:         app.OracleKeeper,
			GuardianKeeper:       app.GuardianKeeper,
			BypassMinFeeMsgTypes: []string{},
=======
	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

<<<<<<< HEAD
	// Set software upgrade execution logic
	app.RegisterUpgradePlan(cfg)

=======
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
<<<<<<< HEAD

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		app.CapabilityKeeper.Seal()
	}
	app.scopedTIBCKeeper = scopedTIBCKeeper
	app.scopedIBCKeeper = scopedIBCKeeper
	app.scopedTransferKeeper = scopedTransferKeeper
=======
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app
}

// Name returns the name of the App
<<<<<<< HEAD
func (app *PetriApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *PetriApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
=======
func (app *App) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app App) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
<<<<<<< HEAD
func (app *PetriApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
=======
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
<<<<<<< HEAD
func (app *PetriApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
=======
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
<<<<<<< HEAD

	// add system service at InitChainer, overwrite if it exists
	var serviceGenState servicetypes.GenesisState
	app.appCodec.MustUnmarshalJSON(genesisState[servicetypes.ModuleName], &serviceGenState)
	serviceGenState.Definitions = append(serviceGenState.Definitions, servicetypes.GenOraclePriceSvcDefinition())
	serviceGenState.Bindings = append(serviceGenState.Bindings, servicetypes.GenOraclePriceSvcBinding(nativeToken.MinUnit))
	serviceGenState.Definitions = append(serviceGenState.Definitions, randomtypes.GetSvcDefinition())
	genesisState[servicetypes.ModuleName] = app.appCodec.MustMarshalJSON(&serviceGenState)

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

=======
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
<<<<<<< HEAD
func (app *PetriApp) LoadHeight(height int64) error {
=======
func (app *App) LoadHeight(height int64) error {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
<<<<<<< HEAD
func (app *PetriApp) ModuleAccountAddrs() map[string]bool {
=======
func (app *App) ModuleAccountAddrs() map[string]bool {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedModuleAccountAddrs returns all the app's blocked module account
// addresses.
<<<<<<< HEAD
func (app *PetriApp) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()

	// remove module accounts that are ALLOWED to received funds
	//
	// TODO: Blocked on updating to v0.46.x
	// delete(modAccAddrs, authtypes.NewModuleAddress(grouptypes.ModuleName).String())
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())
=======
func (app *App) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	delete(modAccAddrs, authtypes.NewModuleAddress(petrichormoduletypes.ModuleName).String())
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
<<<<<<< HEAD
func (app *PetriApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns PetriApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *PetriApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns PetriApp's InterfaceRegistry
func (app *PetriApp) InterfaceRegistry() types.InterfaceRegistry {
=======
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
<<<<<<< HEAD
func (app *PetriApp) GetKey(storeKey string) *storetypes.KVStoreKey {
=======
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
<<<<<<< HEAD
func (app *PetriApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
=======
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
<<<<<<< HEAD
func (app *PetriApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
=======
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
<<<<<<< HEAD
func (app *PetriApp) GetSubspace(moduleName string) paramstypes.Subspace {
=======
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

<<<<<<< HEAD
// SimulationManager implements the SimulationApp interface
func (app *PetriApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided API server.
func (app *PetriApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
=======
// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

<<<<<<< HEAD
	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		lite.RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *PetriApp) RegisterTxService(clientCtx client.Context) {
=======
	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", openapiconsole.Handler(Name, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
<<<<<<< HEAD
func (app *PetriApp) RegisterTendermintService(clientCtx client.Context) {
=======
func (app *App) RegisterTendermintService(clientCtx client.Context) {
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

<<<<<<< HEAD
// RegisterUpgradeHandler implements the upgrade execution logic of the upgrade module
func (app *PetriApp) RegisterUpgradeHandler(
	planName string,
	upgrades *storetypes.StoreUpgrades,
	upgradeHandler sdkupgrade.UpgradeHandler,
) {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		app.Logger().Info("not found upgrade plan", "planName", planName, "err", err.Error())
		return
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		// this configures a no-op upgrade handler for the planName upgrade
		app.UpgradeKeeper.SetUpgradeHandler(planName, upgradeHandler)
		// configure store loader that checks if version+1 == upgradeHeight and applies store upgrades
		app.SetStoreLoader(sdkupgrade.UpgradeStoreLoader(upgradeInfo.Height, upgrades))
	}
=======
// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
<<<<<<< HEAD
	paramsKeeper.Subspace(tokentypes.ModuleName)
	paramsKeeper.Subspace(recordtypes.ModuleName)
	paramsKeeper.Subspace(htlctypes.ModuleName)
	paramsKeeper.Subspace(coinswaptypes.ModuleName)
	paramsKeeper.Subspace(servicetypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(farmtypes.ModuleName)
	paramsKeeper.Subspace(tibchost.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)

	return paramsKeeper
}
=======
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(petrichormoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}
>>>>>>> 78fd37fb52f5e5350fdce166ceb395895b35d201
