package nameservice

import (
	"github.com/anhdungle93/namespace/x/nameservice/keeper"
	"github.com/anhdungle93/namespace/x/nameservice/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	RegisterCodec    = types.RegisterCodec
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper          = keeper.Keeper
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois

	// TODO: Fill out module types
)
