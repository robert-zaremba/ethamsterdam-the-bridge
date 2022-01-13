package module

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type InitGenesisHandler func(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error)
type ExportGenesisHandler func(ctx types.Context, cdc codec.Codec) (json.RawMessage, error)
