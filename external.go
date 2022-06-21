package fireblocks

type WalletRequest struct {
	Name          string `json:"name,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

type WalletAddressRequest struct {
	Address string `json:"address,omitempty"`
	Tag     string `json:"tag,omitempty"`
}
