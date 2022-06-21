package account

import (
	"fireblocks-sdk"
	"fmt"
	"net/url"
)

// GetPaged generates requests for fireblocks.VaultAccounts
func ListPaged(params *fireblocks.VaultAccountsParams) ([]byte, error) {
	var err error
	v := make(url.Values)
	//var result *fireblocks.VaultAccounts

	if len(params.NamePrefix) > 0 {
		v.Add("namePrefix", params.NamePrefix)
	}
	if len(params.NameSuffix) > 0 {
		v.Add("nameSuffix", params.NameSuffix)
	}
	if len(params.MinAmountThreshold) > 0 {
		v.Add("minAmountThreshold", params.MinAmountThreshold)
	}
	if len(params.AssetId) > 0 {
		v.Add("assetId", params.AssetId)
	}
	if len(params.OrderBy) > 0 {
		fmt.Println("params.OrderBy", params.OrderBy)
		v.Add("orderBy", params.OrderBy)
	}
	if len(params.Before) > 0 {
		v.Add("before", params.Before)
	}
	if len(params.After) > 0 {
		v.Add("after", params.After)
	}
	if params.Limit > 0 {
		v.Add("limit", fmt.Sprintf(`%v`, params.Limit))
	}
	response, err := fireblocks.ClientGet("/vault/accounts_paged", v)
	if err != nil {
		return nil, err
	}
	//err = json.Unmarshal(response, &result)
	/*if err != nil {
		return nil, err
	}*/
	return response, nil
}

func GetById(id string) ([]byte, error) {
	var err error
	//v := make(url.Values)
	if len(id) == 0 {
		return nil, fmt.Errorf("error you have to fil acount id")
	}
	path := fmt.Sprintf(`/vault/accounts/%s`, id)
	response, err := fireblocks.ClientGet(path, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func Create(params *fireblocks.VaultAccountsJSONBody) ([]byte, error) {
	var err error
	//var result *fireblocks.VaultAccount
	response, err := fireblocks.ClientPost("/vault/accounts", params)
	if err != nil {
		return nil, err
	}
	//err = json.Unmarshal(response, &result)
	return response, nil
}
