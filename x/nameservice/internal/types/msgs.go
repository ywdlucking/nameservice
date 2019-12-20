package types

/**
Msg触发状态转变。Msgs被包裹在客户端提交至网络的Txs中
Cosmos SDK从Txs中打包和解包来自Msgs
Handler定义了在接收到一个特定Msg时，需要采取的操作

本模块有两种Msg：SetName和BuyName
用户可以发送这些Msg来和应用程序状态进行交互，它们各自同其Handler关联

注意自定义的msg需要继承Msg的五个方法
Type() //返回消息类型
Route() //返回路由消息，路由至合适的模块进行处理
ValidateBasic() //进行一些基本的无状态检查
GetSignBytes() //定义了如何编码Msg以进行签名
GetSigners() //定义一个Tx上需要哪些人的签名才能使其有效
 */
import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

//设置域名值的消息数据结构
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

//构造函数
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

//返回消息处理的模块
func (msg MsgSetName) Route() string { return RouterKey }

//返回消息的处理类型
func (msg MsgSetName) Type() string { return "set_name" }

//基本状态检查
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

//对消息进行编码以进行签名
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

//定义消息需要哪些地址签名
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

//购买域名的消息数据结构
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

func (msg MsgBuyName) Route() string { return RouterKey }

func (msg MsgBuyName) Type() string { return "buy_name" }

func (msg MsgBuyName) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}

func (msg MsgDeleteName) Route() string { return RouterKey }

func (msg MsgDeleteName) Type() string { return "delete_name" }

func (msg MsgDeleteName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}

func (msg MsgDeleteName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}