// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TagResponse400 struct {
	Status *int64  `json:"status,omitempty"`
	Result *string `json:"result,omitempty"`
}

func (o *TagResponse400) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TagResponse400) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}
