package external

import (
	"fmt"
	"github.com/voitakn/fireblocks-sdk"
)

func WalletCreate(params *fireblocks.WalletRequest) ([]byte, error) {
	var err error
	response, err := fireblocks.ClientPost(`/external_wallets`, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func AddressCreate(walletId string, symbol string, params *fireblocks.WalletAddressRequest) ([]byte, error) {
	var err error
	path := fmt.Sprintf(`/external_wallets/%s/%s`, walletId, symbol)
	response, err := fireblocks.ClientPost(path, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}
