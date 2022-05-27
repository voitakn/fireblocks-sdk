package fireblocks

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
	Accounts *[]VaultAccount `json:"accounts,omitempty"`
	NextUrl  string          `json:"nextUrl,omitempty"`
	Paging   struct {
		After  string `json:"after,omitempty"`
		Before string `json:"before,omitempty"`
	} `json:"paging,omitempty"`
	PreviousUrl string `json:"previousUrl,omitempty"`
}

// VaultAccountsParams defines parameters for GetVaultAccountsPaged.
type VaultAccountsParams struct {
	NamePrefix         *string  `json:"namePrefix,omitempty"`
	NameSuffix         *string  `json:"nameSuffix,omitempty"`
	MinAmountThreshold *string  `json:"minAmountThreshold,omitempty"`
	AssetId            *string  `json:"assetId,omitempty"`
	OrderBy            *string  `json:"orderBy,omitempty"`
	Before             *string  `json:"before,omitempty"`
	After              *string  `json:"after,omitempty"`
	Limit              *float64 `json:"limit,omitempty"`
}
