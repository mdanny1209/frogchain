package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePool = "create_pool"

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, poolParam *PoolParam, poolAssets []*PoolAsset) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:    creator,
		PoolParam:  poolParam,
		PoolAssets: poolAssets,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidAddress, "create | invalid creator address (%s)", err)
	}

	swapFeeAmount := msg.PoolParam.SwapFee
	if swapFeeAmount > 10^18 {
		return sdkerrors.Wrapf(ErrFeeOverflow, "create | invalid swap fee (%s)", fmt.Sprint(swapFeeAmount))
	}

	exitFeeAmount := msg.PoolParam.SwapFee
	if exitFeeAmount > 10^18 {
		return sdkerrors.Wrapf(ErrFeeOverflow, "create | invalid exit fee (%s)", fmt.Sprint(exitFeeAmount))
	}

	for _, poolAsset := range msg.PoolAssets {
		weight := poolAsset.Weight
		if weight == 0 {
			return ErrWeightZero
		}

		token := poolAsset.Token
		if !token.IsValid() {
			return sdkerrors.Wrapf(ErrInvalidAddress, "create | invalid pool token address (%s)", fmt.Sprint(exitFeeAmount))
		}
	}
	return nil
}
