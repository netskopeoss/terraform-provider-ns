// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GreResponse405 struct {
	Status *int64  `json:"status,omitempty"`
	Result *string `json:"result,omitempty"`
}

func (o *GreResponse405) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GreResponse405) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}