package types

import "github.com/cosmos/cosmos-sdk/codec"

var ModuleCdc = codec.New()

//创建的任何接口和实现接口的任何结构都需要在RegisterCodec函数中声明
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "nameservice/DeleteName", nil)
}