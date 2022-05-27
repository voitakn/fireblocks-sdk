package account

import (
	"encoding/json"
	"fmt"
	"gitlab.com/fast-rabbit/fireblocks-sdk"
	"net/url"
)

// GetPaged generates requests for fireblocks.VaultAccounts
func ListPaged(params *fireblocks.VaultAccountsParams) (*fireblocks.VaultAccounts, error) {
	var err error
	var v *url.Values
	var result *fireblocks.VaultAccounts

	if params.NamePrefix != nil {
		v.Add("namePrefix", *params.NamePrefix)
	}
	if params.NameSuffix != nil {
		v.Add("nameSuffix", *params.NameSuffix)
	}
	if params.MinAmountThreshold != nil {
		v.Add("minAmountThreshold", *params.MinAmountThreshold)
	}
	if params.AssetId != nil {
		v.Add("assetId", *params.AssetId)
	}
	if params.OrderBy != nil {
		v.Add("orderBy", *params.OrderBy)
	}
	if params.Before != nil {
		v.Add("before", *params.Before)
	}
	if params.After != nil {
		v.Add("after", *params.After)
	}
	if params.Limit != nil {
		v.Add("limit", fmt.Sprintf(`%v`, *params.Limit))
	}
	response, err := fireblocks.ClientGet("/vault/accounts_paged", v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetById(id string) (*fireblocks.VaultAccount, error) {
	var err error
	var v *url.Values
	var result *fireblocks.VaultAccount
	if len(id) == 0 {
		return nil, fmt.Errorf("error you have to fil acount id")
	}
	path := fmt.Sprintf(`/vault/accounts/%s`, id)
	response, err := fireblocks.ClientGet(path, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
