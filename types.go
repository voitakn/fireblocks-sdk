package fireblocks

// Error defines model for Error.
type Error struct {
	Code    float64 `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
}

// FeeInfo defines model for FeeInfo.
type FeeInfo struct {
	GasPrice   string `json:"gasPrice,omitempty"`
	NetworkFee string `json:"networkFee,omitempty"`
	ServiceFee string `json:"serviceFee,omitempty"`
}

type Asset2 struct {
	Id                   string `json:"id"`    //required
	Total                string `json:"total"` //required
	Balance              string `json:"balance"`
	LockedAmount         string `json:"lockedAmount"`
	Available            string `json:"available"` //required
	Pending              string `json:"pending"`   //required
	SelfStakedCPU        string `json:"selfStakedCPU"`
	SelfStakedNetwork    string `json:"selfStakedNetwork"`
	PendingRefundCPU     string `json:"pendingRefundCPU"`
	PendingRefundNetwork string `json:"pendingRefundNetwork"`
	BlockHeight          string `json:"blockHeight"`
	BlockHash            string `json:"blockHash"`
}

type VaultAccount2 struct {
	Id            string    `json:"id"`         //required
	Name          string    `json:"name"`       //required
	HiddenOnUI    bool      `json:"hiddenOnUI"` //required
	Assets        []*Asset2 `json:"assets"`     //required
	CustomerRefId string    `json:"customerRefId"`
	AutoFuel      bool      `json:"autoFuel"`
}

/*
export interface AssetResponse {
    id: string;
    total: string;
	balance?: string;
	lockedAmount?: string;
	available: string;
	pending: string;
	selfStakedCPU?: string;
	selfStakedNetwork?: string;
	pendingRefundCPU?: string;
	pendingRefundNetwork?: string;
	totalStakedCPU?: string;
	totalStakedNetwork?: string;
	rewardInfo?: BalanceRewardInfo;
	blockHeight?: string;
	blockHash?: string;
	allocatedBalances?: {
		allocationId: string;
		thirdPartyAccountId?: string;
		affiliation?: VirtualAffiliation;
		virtualType?: VirtualType;
		total: string;
		available: string;
		pending?: string;
		frozen?: string;
		locked?: string;
	}[];
}
*/

/*
export interface VaultAccountResponse {
    id: string;
    name: string;
    hiddenOnUI: boolean;
    assets: AssetResponse[];
    customerRefId?: string;
    autoFuel: boolean;
}
*/
