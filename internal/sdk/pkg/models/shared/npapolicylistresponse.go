// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NpaPolicyListResponse struct {
	Data []NpaPolicyResponseItem `json:"data,omitempty"`
}

func (o *NpaPolicyListResponse) GetData() []NpaPolicyResponseItem {
	if o == nil {
		return nil
	}
	return o.Data
}