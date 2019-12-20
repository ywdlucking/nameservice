package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//默认名称，以及异常码
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeNameDoesNotExist sdk.CodeType = 101
)

//module找不到返回异常
func ErrNameDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNameDoesNotExist, "Name does not exist")
}