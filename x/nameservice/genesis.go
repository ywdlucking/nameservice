package nameservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	WhoisRecords []Whois `json:"whois_records"`
}

func NewGenesisState(whoIsRecords []Whois) GenesisState {
	return GenesisState{WhoisRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: []Whois{},
	}
}

//在链启动时被调用，此功能将创始状态导入
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

//在链停止后调用，此函数将应用程序状态加载到GenesisState结构中，随后与其他模块的数据一起导出到genesis.json中。
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetWhois(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{WhoisRecords: records}
}