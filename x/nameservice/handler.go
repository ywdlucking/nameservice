package nameservice

/**
收到Msg进行处理的方法
 */

import (
	"fmt"
	"github.com/ywdlucking/nameservice/x/nameservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//通过消息类型返回不同的处理方法
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

//setName消息的处理方法
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	//检查Msg的发送者是否就是域名的所有者
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	keeper.SetName(ctx, msg.Name, msg.Value)
	return sdk.Result{}
}

//buyName消息的处理方法
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	//判断出价
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result()
	}
	//判断域名是否有拥有者
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		//如果没有所有者，你的nameservice模块会把Buyer的资金“燃烧”
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}

//deleteName方法
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDeleteName) sdk.Result {
	//判断域名是否存在
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}
	//判断是否是域名的拥有者
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	keeper.DeleteWhois(ctx, msg.Name)
	return sdk.Result{}
}