// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PrivateAppsResponseData struct {
	ClientlessAccess            *bool                            `json:"clientless_access,omitempty"`
	Host                        *string                          `json:"host,omitempty"`
	ID                          *int                             `json:"id,omitempty"`
	Name                        *string                          `json:"name,omitempty"`
	RealHost                    *string                          `json:"real_host,omitempty"`
	ResolvedProtocols           []ProtocolResponseItem           `json:"protocols,omitempty"`
	ServicePublisherAssignments []ServicePublisherAssignmentItem `json:"service_publisher_assignments,omitempty"`
	Tags                        []TagItem                        `json:"tags,omitempty"`
	TrustSelfSignedCerts        *bool                            `json:"trust_self_signed_certs,omitempty"`
	UsePublisherDNS             *bool                            `json:"use_publisher_dns,omitempty"`
}

func (o *PrivateAppsResponseData) GetClientlessAccess() *bool {
	if o == nil {
		return nil
	}
	return o.ClientlessAccess
}

func (o *PrivateAppsResponseData) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *PrivateAppsResponseData) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PrivateAppsResponseData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PrivateAppsResponseData) GetRealHost() *string {
	if o == nil {
		return nil
	}
	return o.RealHost
}

func (o *PrivateAppsResponseData) GetResolvedProtocols() []ProtocolResponseItem {
	if o == nil {
		return nil
	}
	return o.ResolvedProtocols
}

func (o *PrivateAppsResponseData) GetServicePublisherAssignments() []ServicePublisherAssignmentItem {
	if o == nil {
		return nil
	}
	return o.ServicePublisherAssignments
}

func (o *PrivateAppsResponseData) GetTags() []TagItem {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PrivateAppsResponseData) GetTrustSelfSignedCerts() *bool {
	if o == nil {
		return nil
	}
	return o.TrustSelfSignedCerts
}

func (o *PrivateAppsResponseData) GetUsePublisherDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UsePublisherDNS
}

type PrivateAppsResponse struct {
	Data *PrivateAppsResponseData `json:"data,omitempty"`
}

func (o *PrivateAppsResponse) GetData() *PrivateAppsResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}