// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GrePatchResponse200 struct {
	Data   []GreTunnelResultItem `json:"data,omitempty"`
	Result *string               `json:"result,omitempty"`
	Status *int64                `json:"status,omitempty"`
}

func (o *GrePatchResponse200) GetData() []GreTunnelResultItem {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GrePatchResponse200) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *GrePatchResponse200) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}