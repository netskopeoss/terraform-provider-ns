// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PublisherPutRequest struct {
	ID             *int      `json:"id,omitempty"`
	Lbrokerconnect *bool     `json:"lbrokerconnect,omitempty"`
	Name           *string   `json:"name,omitempty"`
	Tags           []TagItem `json:"tags,omitempty"`
}

func (o *PublisherPutRequest) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PublisherPutRequest) GetLbrokerconnect() *bool {
	if o == nil {
		return nil
	}
	return o.Lbrokerconnect
}

func (o *PublisherPutRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PublisherPutRequest) GetTags() []TagItem {
	if o == nil {
		return nil
	}
	return o.Tags
}