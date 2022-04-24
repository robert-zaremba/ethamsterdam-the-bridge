package axelarbridge

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const DataCodespace = "axelar.bridge"

var (
	ErrResolverURLExists           = sdkerrors.Register(DataCodespace, 4, "resolver URL already exists")
	ErrResolverUndefined           = sdkerrors.Register(DataCodespace, 5, "resolver undefined")
	ErrUnauthorizedResolverManager = sdkerrors.Register(DataCodespace, 6, "unauthorized resolver manager")
)
