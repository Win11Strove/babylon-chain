package keeper

import (
	"context"
	"errors"

	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/babylonchain/babylon/x/checkpointing/types"
)

type msgServer struct {
	k Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

var _ types.MsgServer = msgServer{}

// AddBlsSig adds BLS sig messages and changes a raw checkpoint status to SEALED if sufficient voting power is accumulated
func (m msgServer) AddBlsSig(goCtx context.Context, msg *types.MsgAddBlsSig) (*types.MsgAddBlsSigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.k.addBlsSig(ctx, msg.BlsSig)
	if err != nil {
		return &types.MsgAddBlsSigResponse{}, err
	}

	return &types.MsgAddBlsSigResponse{}, nil
}

// WrappedCreateValidator registers validator's BLS public key
// and forwards corresponding MsgCreateValidator message to
// the epoching module
func (m msgServer) WrappedCreateValidator(goCtx context.Context, msg *types.MsgWrappedCreateValidator) (*types.MsgWrappedCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// verify Proof-of-Possession
	valAddr, err := sdk.ValAddressFromBech32(msg.MsgCreateValidator.ValidatorAddress)
	if err != nil {
		return nil, err
	}
	valPubkey, found := m.k.GetValidatorPubkey(ctx, valAddr)
	if !found {
		return nil, errors.New("can not find validator by address")
	}
	ok := msg.VerifyPoP(valPubkey)
	if !ok {
		return nil, errors.New("the proof-of-possession is not valid")
	}

	// store BLS public key with validator address
	err = m.k.CreateRegistration(ctx, *msg.Key.Pubkey, valAddr)
	if err != nil {
		return nil, err
	}

	// enqueue the msg into the epoching module
	queueMsg := epochingtypes.QueuedMessage{
		Msg: &epochingtypes.QueuedMessage_MsgCreateValidator{MsgCreateValidator: msg.MsgCreateValidator},
	}

	m.k.epochingKeeper.EnqueueMsg(ctx, queueMsg)

	return &types.MsgWrappedCreateValidatorResponse{}, err
}
