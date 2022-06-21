package fireblocks

import "encoding/json"

// VaultAccount defines model for VaultAccount.
type VaultAccount struct {
	Assets        []*VaultAsset `json:"assets,omitempty"`
	AutoFuel      bool          `json:"autoFuel,omitempty"`
	CustomerRefId string        `json:"customerRefId,omitempty"`
	HiddenOnUI    bool          `json:"hiddenOnUI,omitempty"`
	Id            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
}

// VaultAccounts defines model for VaultAccountsPagedResponse.
type VaultAccounts struct {
	Accounts []*VaultAccount `json:"accounts,omitempty"`
	NextUrl  string          `json:"nextUrl,omitempty"`
	Paging   struct {
		After  string `json:"after,omitempty"`
		Before string `json:"before,omitempty"`
	} `json:"paging,omitempty"`
	PreviousUrl string `json:"previousUrl,omitempty"`
}

// VaultAccountsParams defines parameters for GetVaultAccountsPaged.
type VaultAccountsParams struct {
	NamePrefix         string `json:"namePrefix,omitempty"`
	NameSuffix         string `json:"nameSuffix,omitempty"`
	MinAmountThreshold string `json:"minAmountThreshold,omitempty"`
	AssetId            string `json:"assetId,omitempty"`
	OrderBy            string `json:"orderBy,omitempty"`
	Before             string `json:"before,omitempty"`
	After              string `json:"after,omitempty"`
	Limit              int64  `json:"limit,omitempty"`
}

// VaultAccountsJSONBody defines parameters for PostVaultAccounts.
type VaultAccountsJSONBody struct {
	// Optional - Sets the autoFuel property of the vault account
	AutoFuel bool `json:"autoFuel,omitempty"`
	// Optional - Sets a customer reference ID
	CustomerRefId string `json:"customerRefId,omitempty"`
	// Optional - if true, the created account and all related transactions will not be shown on Fireblocks console
	HiddenOnUI bool `json:"hiddenOnUI,omitempty"`
	// Account Name
	Name string `json:"name,omitempty"`
}

func (c *VaultAccount) UnmarshalJSON(data []byte) error {
	type va VaultAccount
	var v va
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*c = VaultAccount(v)
	return nil
}

func (c *VaultAccounts) UnmarshalJSON(data []byte) error {
	type va VaultAccounts
	var v va
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*c = VaultAccounts(v)
	return nil
}
