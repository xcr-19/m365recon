package pkg

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/xcr-19/m365recon/utils"
)

var verbose bool = false

func ReconByDomain(domain string, config utils.Config) error {
	verbose = config.Verbose
	output := config.Output
	proxy := config.Proxy

	outputInfo := utils.OutputInfo{
		Domain: domain,
	}
	if verbose {
		fmt.Println("Starting recon for domain: ", domain)
	}

	oidInfo, err := GetOIDInfo(domain, proxy)
	if err != nil {
		if config.Output != "" {
			utils.WriteToFile(utils.OutputInfo{
				Domain: domain,
				Error:  err.Error(),
			}, config.Output)
		}
		return err
	}
	userRelmInfo, err := GetUserRelmInfo(domain, proxy)
	if err != nil {
		return err
	}

	extendedUserRelmInfo, err := GetExtendedUserRelmInfo(domain, proxy)
	if err != nil {
		return err
	}

	additionalDomains, err := GetAdditionalDomains(domain, proxy)
	if err != nil {
		return err
	}
	tenantId, err := url.Parse(oidInfo.SuccessOIDInfo.Issuer)
	if err != nil {
		return err
	}

	cleanTenantId := strings.Split(tenantId.Path, "/")[1]
	outputInfo.TenantID = cleanTenantId
	outputInfo.TenantRegion = oidInfo.SuccessOIDInfo.TenantRegionScope
	outputInfo.UserinfoEndpoint = oidInfo.SuccessOIDInfo.UserinfoEndpoint
	outputInfo.KerberosEndpoint = oidInfo.SuccessOIDInfo.KerberosEndpoint
	outputInfo.TokenEndpoint = oidInfo.SuccessOIDInfo.TokenEndpoint
	outputInfo.FederationBrandName = userRelmInfo.SuccessUserRelmInfo.FederationBrandName
	outputInfo.AuthURL = userRelmInfo.SuccessUserRelmInfo.AuthURL
	outputInfo.IsDssoEnabled = extendedUserRelmInfo.SuccessExtendedUserRelmInfo.IsDssoEnabled
	outputInfo.ForceLoginHint = extendedUserRelmInfo.SuccessExtendedUserRelmInfo.ForceLoginHint
	outputInfo.AdditionalDomains = additionalDomains.Body.GetFederationInformationResponseMessage.Response.Domains.Domain
	outputInfo.Error = "None"

	utils.TablePrintOutputInfo(outputInfo)

	if output != "" {
		fmt.Println("Writing output to file: ", output)
		utils.WriteToFile(outputInfo, output)
	}

	return nil
}

func GetUserRelmInfo(domain string, proxy string) (utils.UserRelmInfo, error) {
	if verbose {
		fmt.Println("Getting user realm info")
	}
	userRelmInfo := utils.UserRelmInfo{}
	client := SetupHTTPClient(proxy)
	request, err := GetRequest(fmt.Sprintf(utils.AzureUserRelm, domain), "GET")
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Request Info:")
		fmt.Println("--------------------------------")
		fmt.Println("URL: ", request.URL)
		fmt.Println("Method: ", request.Method)
		fmt.Println("Headers: ", request.Header)
		fmt.Println("Body: ", request.Body)
		fmt.Println("--------------------------------")
	}
	response, err := client.Do(request)
	if err != nil {
		return utils.UserRelmInfo{}, err
	}
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Response Info:")
		fmt.Println("--------------------------------")
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Headers: ")
		for key, value := range response.Header {
			fmt.Println(key, ":", value)
		}
		fmt.Println("--------------------------------")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return utils.UserRelmInfo{}, fmt.Errorf("Got status code %d while getting user realm info", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return utils.UserRelmInfo{}, err
	}
	if verbose {
		fmt.Println("Response Body:", string(body))
	}
	if err := json.Unmarshal(body, &userRelmInfo.SuccessUserRelmInfo); err != nil {
		err := json.Unmarshal(body, &userRelmInfo.ErrorUserRelmInfo)
		if err != nil {
			return utils.UserRelmInfo{}, err
		}
		return utils.UserRelmInfo{}, fmt.Errorf("Error: %s", userRelmInfo.ErrorUserRelmInfo.NameSpaceType)
	}

	return userRelmInfo, nil
}

func GetExtendedUserRelmInfo(domain string, proxy string) (utils.ExtendedUserRelmInfo, error) {
	if verbose {
		fmt.Println("Getting extended user realm info")
	}
	extendedUserRelmInfo := utils.ExtendedUserRelmInfo{}
	client := SetupHTTPClient(proxy)
	request, err := GetRequest(fmt.Sprintf(utils.AzureExtendedUserRelm, domain), "GET")
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Request Info:")
		fmt.Println("--------------------------------")
		fmt.Println("URL: ", request.URL)
		fmt.Println("Method: ", request.Method)
		fmt.Println("Headers: ", request.Header)
	}
	response, err := client.Do(request)
	if err != nil {
		return utils.ExtendedUserRelmInfo{}, err
	}
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Response Info:")
		fmt.Println("--------------------------------")
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Headers: ")
		for key, value := range response.Header {
			fmt.Println(key, ":", value)
		}
		fmt.Println("--------------------------------")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return utils.ExtendedUserRelmInfo{}, fmt.Errorf("Got status code %d while getting extended user realm info", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return utils.ExtendedUserRelmInfo{}, err
	}
	if verbose {
		fmt.Println("Response Body:", string(body))
	}
	if err := json.Unmarshal(body, &extendedUserRelmInfo.SuccessExtendedUserRelmInfo); err != nil {
		err := json.Unmarshal(body, &extendedUserRelmInfo.ErrorExtendedUserRelmInfo)
		if err != nil {
			return utils.ExtendedUserRelmInfo{}, err
		}
		return utils.ExtendedUserRelmInfo{}, fmt.Errorf("Error: %s", extendedUserRelmInfo.ErrorExtendedUserRelmInfo.NameSpaceType)
	}
	return extendedUserRelmInfo, nil
}

func GetOIDInfo(domain string, proxy string) (utils.OIDInfo, error) {
	if verbose {
		fmt.Println("Getting OID info")
	}
	oidInfo := utils.OIDInfo{}
	client := SetupHTTPClient(proxy)
	request, err := GetRequest(fmt.Sprintf(utils.AzureOidInfo, domain), "GET")
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Request Info:")
		fmt.Println("--------------------------------")
		fmt.Println("URL: ", request.URL)
		fmt.Println("Method: ", request.Method)
		fmt.Println("Headers: ", request.Header)
		fmt.Println("--------------------------------")
	}
	response, err := client.Do(request)
	if err != nil {
		return utils.OIDInfo{}, err
	}
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Response Info:")
		fmt.Println("--------------------------------")
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Headers: ")
		for key, value := range response.Header {
			fmt.Println(key, ":", value)
		}
		fmt.Println("--------------------------------")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if verbose {
		fmt.Println("Response Body:", string(body))
	}
	if response.StatusCode == 400 {
		if err := json.Unmarshal(body, &oidInfo.ErrorOIDInfo); err != nil {
			return utils.OIDInfo{}, err
		}
		return utils.OIDInfo{}, fmt.Errorf("Error: %s", oidInfo.ErrorOIDInfo.Error)
	}
	if response.StatusCode != 200 {
		return utils.OIDInfo{}, fmt.Errorf("Request failed with status code %d", response.StatusCode)
	}
	if err := json.Unmarshal(body, &oidInfo.SuccessOIDInfo); err != nil {
		return utils.OIDInfo{}, err
	}
	return oidInfo, nil
}

func GetAdditionalDomains(domain string, proxy string) (utils.FederationInfoResponse, error) {
	if verbose {
		fmt.Println("Getting additional domains via autodiscover")
	}
	soapBody := fmt.Sprintf(utils.AzureFederationInfoRequest, domain)
	client := SetupHTTPClient(proxy)
	request, err := GetRequest(utils.AzureDomainEmum, "POST")
	request.Body = io.NopCloser(strings.NewReader(soapBody))
	request.Header.Set("Content-Type", "text/xml; charset=utf-8")
	request.Header.Set("SOAPAction", "http://schemas.microsoft.com/exchange/2010/Autodiscover/Autodiscover/GetFederationInformation")
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Request Info:")
		fmt.Println("--------------------------------")
		fmt.Println("URL: ", request.URL)
		fmt.Println("Method: ", request.Method)
		fmt.Println("Headers: ", request.Header)
		fmt.Println("--------------------------------")
	}
	response, err := client.Do(request)
	if err != nil {
		return utils.FederationInfoResponse{}, err
	}
	defer response.Body.Close()
	if verbose {
		fmt.Println("--------------------------------")
		fmt.Println("Response Info:")
		fmt.Println("--------------------------------")
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Headers: ")
		for key, value := range response.Header {
			fmt.Println(key, ":", value)
		}
		fmt.Println("--------------------------------")
	}
	if response.Header.Get("X-Proxyerrormessage") == "The network is busy." {
		fmt.Println("Microsoft network is busy got from proxymessage, retrying...")
		time.Sleep(1 * time.Second)
		return GetAdditionalDomains(domain, proxy)
	}
	body, err := io.ReadAll(response.Body)
	if verbose {
		fmt.Println("Response Body:", string(body))
	}
	var federationInfo utils.FederationInfoResponse
	if err := xml.Unmarshal(body, &federationInfo); err != nil {
		return utils.FederationInfoResponse{}, fmt.Errorf("failed to unmarshal federation info response: %v", err)
	}

	return federationInfo, nil
}
