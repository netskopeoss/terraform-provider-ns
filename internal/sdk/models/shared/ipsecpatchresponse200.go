// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IpsecPatchResponse200 struct {
	Status *int64                  `json:"status,omitempty"`
	Result *string                 `json:"result,omitempty"`
	Data   []IpsecTunnelResultItem `json:"data,omitempty"`
}

func (o *IpsecPatchResponse200) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *IpsecPatchResponse200) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *IpsecPatchResponse200) GetData() []IpsecTunnelResultItem {
	if o == nil {
		return nil
	}
	return o.Data
}