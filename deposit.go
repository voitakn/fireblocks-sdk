package fireblocks

// CreateDeposit defines parameters for fireblocks.deposit.Create
type CreateDeposit struct {
	// Optional - Sets a customer reference ID
	CustomerRefId string `json:"customerRefId,omitempty"`
	// (Optional) Attach a description to the new address
	Description string `json:"description,omitempty"`
}

// UpdateDeposit defines parameters for fireblocks.deposit.Update
type UpdateDeposit struct {
	// (Optional) Attach a description to the new address
	Description string `json:"description,omitempty"`
}

// CreateDepositResponse defines model for CreateAddressResponse.
type CreateDepositResponse struct {
	Address           string `json:"address,omitempty"`
	Bip44AddressIndex int64  `json:"bip44AddressIndex,omitempty"`
	EnterpriseAddress string `json:"enterpriseAddress,omitempty"`
	LegacyAddress     string `json:"legacyAddress,omitempty"`
	Tag               string `json:"tag,omitempty"`
}
