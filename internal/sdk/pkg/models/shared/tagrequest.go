// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TagRequest struct {
	ID            *string   `json:"id,omitempty"`
	Ids           []string  `json:"ids,omitempty"`
	PublisherTags []TagItem `json:"publisher_tags,omitempty"`
	Tags          []TagItem `json:"tags,omitempty"`
}

func (o *TagRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TagRequest) GetIds() []string {
	if o == nil {
		return nil
	}
	return o.Ids
}

func (o *TagRequest) GetPublisherTags() []TagItem {
	if o == nil {
		return nil
	}
	return o.PublisherTags
}

func (o *TagRequest) GetTags() []TagItem {
	if o == nil {
		return nil
	}
	return o.Tags
}