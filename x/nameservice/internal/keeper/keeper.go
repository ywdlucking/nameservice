package keeper

import (
	//提供负责Cosmos编码格式的工具——Amino
	"github.com/cosmos/cosmos-sdk/codec"
	//bank模块控制账户和转账
	sdk "github.com/cosmos/cosmos-sdk/types"
	//types包含了整个SDK常用的类型
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ywdlucking/nameservice/x/nameservice/internal/types"
)

//使用哥哥模块的状态机，维护数据存储
type Keeper struct {
	//指向bank module，并且可以调用该模块的的方法
	CoinKeeper bank.Keeper
	//访问一个持久化保存你的应用程序状态的sdk.KVStore
	storeKey  sdk.StoreKey
	//二进制机构的编码解码器
	cdc *codec.Codec
}

//添加一个函数来为指定域名设置解析字符串值
//sdk.Context该对象持有访问像blockHeight和chainID这样重要部分状态的函数。
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	//存储只接受[]byte
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}

//添加一个函数来解析域名（即查找域名对应的解析值）
func (k Keeper) GetWhois(ctx sdk.Context, name string) types.Whois {
	store := ctx.KVStore(k.storeKey)
	//如果一个域名尚未在存储中，它返回一个新的 Whois 信息
	if !k.IsNamePresent(ctx, name) {
		return types.NewWhois()
	}
	bz := store.Get([]byte(name))
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}

//添加一个函数来删除通过域名
func (k Keeper) DeleteWhois(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// 通过GetWhois和SetWhois来获取和修改参数的值
//通过name 返回值
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetWhois(ctx, name).Value
}

//通过name修改value
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

// 判断域名是否有拥有者
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx, name).Owner.Empty()
}

// 返回域名的拥有者
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx, name).Owner
}

// 设置域名的拥有者
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois := k.GetWhois(ctx, name)
	whois.Owner = owner
	k.SetWhois(ctx, name, whois)
}

// 返回域名的价格
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

// 设置域名的价格
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

// 判断store是否存有该域名
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

//增加一个函数用于获取遍历 store 中所有已知域名的迭代器。
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

//keeper的构造器
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}