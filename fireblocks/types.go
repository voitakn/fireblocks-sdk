package fireblocks

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Slice returns a slice of int64 pointers given a slice of int64s.
func Int64Slice(v []int64) []*int64 {
	out := make([]*int64, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlice returns a slice of string pointers given a slice of strings.
func StringSlice(v []string) []*string {
	out := make([]*string, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

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
