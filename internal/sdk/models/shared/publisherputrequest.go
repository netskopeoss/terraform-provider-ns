// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PublisherPutRequest struct {
	Name                       *string `json:"name,omitempty"`
	ID                         int     `json:"id"`
	Lbrokerconnect             *bool   `json:"lbrokerconnect,omitempty"`
	PublisherUpgradeProfilesID *int    `json:"publisher_upgrade_profiles_id,omitempty"`
}

func (o *PublisherPutRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PublisherPutRequest) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *PublisherPutRequest) GetLbrokerconnect() *bool {
	if o == nil {
		return nil
	}
	return o.Lbrokerconnect
}

func (o *PublisherPutRequest) GetPublisherUpgradeProfilesID() *int {
	if o == nil {
		return nil
	}
	return o.PublisherUpgradeProfilesID
}
