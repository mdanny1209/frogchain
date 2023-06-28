package keeper

import (
	"context"

	"frogchain/x/amm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SwapExactTokensForTokens(goCtx context.Context, msg *types.MsgSwapExactTokensForTokens) (*types.MsgSwapExactTokensForTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapExactTokensForTokensResponse{}, nil
}
