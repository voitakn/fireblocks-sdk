package deposit

import (
	"fmt"
	"github.com/voitakn/fireblocks-sdk"
)

func Create(symbol string, params *fireblocks.CreateDeposit) ([]byte, error) {
	var err error
	path := fmt.Sprintf(`/vault/accounts/%s/%s/addresses`, fireblocks.VaultId, symbol)
	response, err := fireblocks.ClientPost(path, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func List(symbol string) ([]byte, error) {
	var err error
	if len(symbol) == 0 {
		return nil, fmt.Errorf("error you have to fil acount id")
	}
	path := fmt.Sprintf(`/vault/accounts/%s/%s/addresses`, fireblocks.VaultId, symbol)
	response, err := fireblocks.ClientGet(path, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func Update(symbol string, address string, params *fireblocks.UpdateDeposit) ([]byte, error) {
	var err error
	path := fmt.Sprintf(`/vault/accounts/%s/%s/addresses/%s`, fireblocks.VaultId, symbol, address)
	response, err := fireblocks.ClientPut(path, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}
