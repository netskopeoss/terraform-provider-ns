// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IpsecResponse400 struct {
	Status *int64  `json:"status,omitempty"`
	Result *string `json:"result,omitempty"`
}

func (o *IpsecResponse400) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *IpsecResponse400) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}