package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

/**
Value - 域名解析出为的值。
Owner - 该域名当前所有者的地址。
Price - 你需要为购买域名支付的费用。
 */
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

//最初从未拥有过的名称的起始价格
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

//名称尚未有所有者，用 MinPrice 对其进行初始化。
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s Value: %s Price: %s`, w.Owner, w.Value, w.Price))
}