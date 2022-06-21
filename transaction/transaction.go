package transaction

import fireblocks "fireblocks-sdk"

func Create(params *fireblocks.TransactionRequest) ([]byte, error) {
	var err error
	response, err := fireblocks.ClientPost(`/transactions`, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}
