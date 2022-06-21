package fireblocks

import "encoding/json"

// OneTimeAddress defines model for OneTimeAddress.
type OneTimeAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag,omitempty"`
}

// TransferPeerPath defines model for TransferPeerPath.
type TransferPeerPath struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type"`
}

// DestinationTransferPeerPath defines model for DestinationTransferPeerPath.
type DestinationTransferPeerPath struct {
	// Embedded struct due to allOf(#/components/schemas/TransferPeerPath)
	TransferPeerPath `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	OneTimeAddress *OneTimeAddress `json:"oneTimeAddress,omitempty"`
}

// TransactionRequestDestination defines model for TransactionRequestDestination.
type TransactionRequestDestination struct {
	Amount      string                       `json:"amount,omitempty"`
	Destination *DestinationTransferPeerPath `json:"destination,omitempty"`
}

// TransactionRequest defines model for TransactionRequest.
type TransactionRequest struct {
	Amount          interface{}                      `json:"amount,omitempty"`
	AssetId         string                           `json:"assetId,omitempty"`
	AutoStaking     bool                             `json:"autoStaking,omitempty"`
	CpuStaking      interface{}                      `json:"cpuStaking,omitempty"`
	CustomerRefId   string                           `json:"customerRefId,omitempty"`
	Destination     *DestinationTransferPeerPath     `json:"destination,omitempty"`
	Destinations    []*TransactionRequestDestination `json:"destinations,omitempty"`
	ExternalTxId    string                           `json:"externalTxId,omitempty"`
	ExtraParameters map[string]interface{}           `json:"extraParameters,omitempty"`
	FailOnLowFee    bool                             `json:"failOnLowFee,omitempty"`
	// - For BTC based assets, the fee per bytes in the asset's smallest unit (Satoshi/Latoshi etc.)
	// - For XRP, the fee for the transaction
	Fee          interface{}                 `json:"fee,omitempty"`
	FeeLevel     *TransactionRequestFeeLevel `json:"feeLevel,omitempty"`
	FeePayerInfo struct {
		FeePayerAccountId string `json:"feePayerAccountId,omitempty"`
	} `json:"feePayerInfo,omitempty"`
	ForceSweep bool        `json:"forceSweep,omitempty"`
	GasLimit   interface{} `json:"gasLimit,omitempty"`

	// For ETH based asset only. this will be used instead of the fee property. Value is in gwei
	GasPrice interface{} `json:"gasPrice,omitempty"`
	// - For ETH based assets only, must be provided when using the priorityFee parameter. This will be used to limit the fee max possible fee (according to eip-1559). Value is in gwei.
	// - For other assets, it fails a transaction when the automatically selected fee is higher than the provided value.
	MaxFee         string                `json:"maxFee,omitempty"`
	NetworkFee     interface{}           `json:"networkFee,omitempty"`
	NetworkStaking interface{}           `json:"networkStaking,omitempty"`
	Note           string                `json:"note,omitempty"`
	Operation      *TransactionOperation `json:"operation,omitempty"`
	// - For ETH based assets only, the fee for eip-1559 transaction pricing mechanism. Value is in gwei.
	PriorityFee        interface{}       `json:"priorityFee,omitempty"`
	ReplaceTxByHash    string            `json:"replaceTxByHash,omitempty"`
	Source             *TransferPeerPath `json:"source,omitempty"`
	TreatAsGrossAmount bool              `json:"treatAsGrossAmount,omitempty"`
}

func (c *TransactionRequest) ToStruct(data []byte) error {
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	return nil
}
func (c *TransactionRequest) ToJson() ([]byte, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return []byte(``), err
	}
	return b, nil
}
