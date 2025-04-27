package utils

const (
	AzureUserRelm              = "https://login.microsoftonline.com/getuserrealm.srf?login=%s"
	AzureExtendedUserRelm      = "https://login.microsoftonline.com/common/userrealm/%s?api-version=2.0"
	AzureOidInfo               = "https://login.microsoftonline.com/%s/v2.0/.well-known/openid-configuration"
	AzureDomainEmum            = "https://autodiscover-s.outlook.com/autodiscover/autodiscover.svc"
	AzureLoginUrl              = "https://login.microsoftonline.com/%s/oauth2/v2.0/authorize"
	AzureFederationInfoRequest = `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
				   xmlns:a="http://www.w3.org/2005/08/addressing"
				   xmlns:autodiscover="http://schemas.microsoft.com/exchange/2010/Autodiscover">
	  <soap:Header>
		<a:Action>http://schemas.microsoft.com/exchange/2010/Autodiscover/Autodiscover/GetFederationInformation</a:Action>
		<a:To>https://autodiscover-s.outlook.com/autodiscover/autodiscover.svc</a:To>
	  </soap:Header>
	  <soap:Body>
		<autodiscover:GetFederationInformationRequestMessage>
		  <autodiscover:Request>
			<autodiscover:Domain>%s</autodiscover:Domain>
		  </autodiscover:Request>
		</autodiscover:GetFederationInformationRequestMessage>
	  </soap:Body>
	</soap:Envelope>`
)

var UserAgent = []string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
}
