// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GrePopResultItem struct {
	Acceptingtunnels *bool   `json:"acceptingtunnels,omitempty"`
	Distance         *string `json:"distance,omitempty"`
	Gateway          *string `json:"gateway,omitempty"`
	ID               *string `json:"id,omitempty"`
	Location         *string `json:"location,omitempty"`
	Name             *string `json:"name,omitempty"`
	Probeip          *string `json:"probeip,omitempty"`
	Region           *string `json:"region,omitempty"`
}

func (o *GrePopResultItem) GetAcceptingtunnels() *bool {
	if o == nil {
		return nil
	}
	return o.Acceptingtunnels
}

func (o *GrePopResultItem) GetDistance() *string {
	if o == nil {
		return nil
	}
	return o.Distance
}

func (o *GrePopResultItem) GetGateway() *string {
	if o == nil {
		return nil
	}
	return o.Gateway
}

func (o *GrePopResultItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GrePopResultItem) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *GrePopResultItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GrePopResultItem) GetProbeip() *string {
	if o == nil {
		return nil
	}
	return o.Probeip
}

func (o *GrePopResultItem) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}