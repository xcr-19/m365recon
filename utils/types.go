package utils

import "encoding/xml"

type SuccessOIDInfo struct {
	TokenEndpoint                     string   `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	JwksURI                           string   `json:"jwks_uri"`
	ResponseModesSupported            []string `json:"response_modes_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	Issuer                            string   `json:"issuer"`
	MicrosoftMultiRefreshToken        bool     `json:"microsoft_multi_refresh_token"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	DeviceAuthorizationEndpoint       string   `json:"device_authorization_endpoint"`
	HTTPLogoutSupported               bool     `json:"http_logout_supported"`
	FrontchannelLogoutSupported       bool     `json:"frontchannel_logout_supported"`
	EndSessionEndpoint                string   `json:"end_session_endpoint"`
	ClaimsSupported                   []string `json:"claims_supported"`
	CheckSessionIframe                string   `json:"check_session_iframe"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	KerberosEndpoint                  string   `json:"kerberos_endpoint"`
	TenantRegionScope                 string   `json:"tenant_region_scope"`
	CloudInstanceName                 string   `json:"cloud_instance_name"`
	CloudGraphHostName                string   `json:"cloud_graph_host_name"`
	MsgraphHost                       string   `json:"msgraph_host"`
	RbacURL                           string   `json:"rbac_url"`
}

type ErrorOIDInfo struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	Timestamp        string `json:"timestamp"`
	TraceID          string `json:"trace_id"`
	CorrelationID    string `json:"correlation_id"`
	ErrorURI         string `json:"error_uri"`
}

type OIDInfo struct {
	SuccessOIDInfo SuccessOIDInfo `json:"success"`
	ErrorOIDInfo   ErrorOIDInfo   `json:"error"`
}

type UserRelmInfo struct {
	SuccessUserRelmInfo SuccessUserRelmInfo `json:"success"`
	ErrorUserRelmInfo   ErrorUserRelmInfo   `json:"error"`
}

type SuccessUserRelmInfo struct {
	State                   int    `json:"State"`
	UserState               int    `json:"UserState"`
	Login                   string `json:"Login"`
	NameSpaceType           string `json:"NameSpaceType"`
	DomainName              string `json:"DomainName"`
	FederationGlobalVersion int    `json:"FederationGlobalVersion"`
	AuthURL                 string `json:"AuthURL"`
	FederationBrandName     string `json:"FederationBrandName"`
	CloudInstanceName       string `json:"CloudInstanceName"`
	CloudInstanceIssuerURI  string `json:"CloudInstanceIssuerUri"`
}

type ErrorUserRelmInfo struct {
	State         int    `json:"State"`
	UserState     int    `json:"UserState"`
	Login         string `json:"Login"`
	NameSpaceType string `json:"NameSpaceType"`
}

type SuccessExtendedUserRelmInfo struct {
	NameSpaceType       string `json:"NameSpaceType"`
	Login               string `json:"Login"`
	DomainName          string `json:"DomainName"`
	FederationBrandName string `json:"FederationBrandName"`
	TenantBrandingInfo  []struct {
		Locale                 int    `json:"Locale"`
		BoilerPlateText        string `json:"BoilerPlateText"`
		KeepMeSignedInDisabled bool   `json:"KeepMeSignedInDisabled"`
		UseTransparentLightBox bool   `json:"UseTransparentLightBox"`
	} `json:"TenantBrandingInfo"`
	CloudInstanceName string `json:"cloud_instance_name"`
	IsDssoEnabled     bool   `json:"is_dsso_enabled"`
	ForceLoginHint    bool   `json:"force_login_hint"`
}

type ErrorExtendedUserRelmInfo struct {
	NameSpaceType     string `json:"NameSpaceType"`
	Login             string `json:"Login"`
	CloudInstanceName string `json:"cloud_instance_name"`
}

type ExtendedUserRelmInfo struct {
	SuccessExtendedUserRelmInfo SuccessExtendedUserRelmInfo `json:"success"`
	ErrorExtendedUserRelmInfo   ErrorExtendedUserRelmInfo   `json:"error"`
}

type FederationInfoResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  struct {
		Action        string `xml:"Action"`
		ServerVersion struct {
			MajorVersion     int    `xml:"MajorVersion"`
			MinorVersion     int    `xml:"MinorVersion"`
			MajorBuildNumber int    `xml:"MajorBuildNumber"`
			MinorBuildNumber int    `xml:"MinorBuildNumber"`
			Version          string `xml:"Version"`
		} `xml:"ServerVersionInfo"`
	} `xml:"Header"`
	Body struct {
		GetFederationInformationResponseMessage struct {
			Response struct {
				ErrorCode      string `xml:"ErrorCode"`
				ErrorMessage   string `xml:"ErrorMessage"`
				ApplicationUri string `xml:"ApplicationUri"`
				Domains        struct {
					Domain []string `xml:"Domain"`
				} `xml:"Domains"`
				TokenIssuers struct {
					TokenIssuer []struct {
						Endpoint string `xml:"Endpoint"`
						Uri      string `xml:"Uri"`
					} `xml:"TokenIssuer"`
				} `xml:"TokenIssuers"`
			} `xml:"Response"`
		} `xml:"GetFederationInformationResponseMessage"`
	} `xml:"Body"`
}

type OutputInfo struct {
	Domain              string
	TenantID            string
	TenantRegion        string
	FederationBrandName string
	AdditionalDomains   []string
	IsDssoEnabled       bool
	ForceLoginHint      bool
	AuthURL             string
	UserinfoEndpoint    string
	KerberosEndpoint    string
	TokenEndpoint       string
}
