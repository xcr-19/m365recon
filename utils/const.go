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
	banner = `
	_ __________                        _________                      __  .__    .__                 
	\__    ___/__.__.______   ____    /   _____/ ____   _____   _____/  |_|  |__ |__| ____    ____   
	  |    | <   |  |\____ \_/ __ \   \_____  \ /  _ \ /     \_/ __ \   __\  |  \|  |/    \  / ___\  
	  |    |  \___  ||  |_> >  ___/   /        (  <_> )  Y Y  \  ___/|  | |   Y  \  |   |  \/ /_/  > 
	  |____|  / ____||   __/ \___  > /_______  /\____/|__|_|  /\___  >__| |___|  /__|___|  /\___  /  
	           \/     |__|        \/          \/             \/     \/          \/        \//_____/ `
)
