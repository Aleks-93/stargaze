package funding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/funding/types"
)

// NewHandler creates an sdk.Handler for all the funding type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgBuy:
			return handleMsgBuy(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgBuy(ctx sdk.Context, k Keeper, msg types.MsgBuy) (*sdk.Result, error) {
	// transfer/ibczeroxfer/stake
	denom := fmt.Sprintf("transfer/%s/%s", types.Counterparty, types.CounterpartyDenom)
	lockCoin := sdk.NewCoin(denom, msg.Amount.Amount)

	err := k.DistributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(lockCoin), msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "can't transfer %s coins from sender to dao", denom)
	}

	tokenSupply := k.BankKeeper.GetSupply(ctx)
	poolBalance := k.DistributionKeeper.GetFeePool(ctx).CommunityPool.AmountOf(lockCoin.Denom)
	curve := GetBondingCurve()
	toMint, err := curve.ToMint(lockCoin, tokenSupply.GetTotal().AmountOf(types.Denom), poolBalance.TruncateInt())
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "can't determine the number of tokens to mint from bonding curve")
	}
	err = mintNewToken(ctx, k, toMint, msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "can't mint new token from bonding curve")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func mintNewToken(ctx sdk.Context, k Keeper, amount sdk.Int, sender sdk.AccAddress) error {
	newCoin := sdk.NewCoin(types.Denom, amount)
	err := k.BankKeeper.MintCoins(ctx, ModuleName, sdk.NewCoins(newCoin))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "can't mint %s", newCoin.Denom)
	}
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, sender, sdk.NewCoins(newCoin))
	if err != nil {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "can't transfer %s coins from module account to sender", newCoin.Denom)
	}

	return nil
}
