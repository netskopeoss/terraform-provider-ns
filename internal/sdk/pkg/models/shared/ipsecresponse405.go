// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IpsecResponse405 struct {
	Result *string `json:"result,omitempty"`
	Status *int64  `json:"status,omitempty"`
}

func (o *IpsecResponse405) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *IpsecResponse405) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}