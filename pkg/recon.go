package pkg

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/xcr-19/m365recon/utils"
)

func ReconByDomain(domain string) error {
	outputInfo := utils.OutputInfo{
		Domain: domain,
	}

	oidInfo, err := GetOIDInfo(domain)
	if err != nil {
		return err
	}

	userRelmInfo, err := GetUserRelmInfo(domain)
	if err != nil {
		return err
	}

	extendedUserRelmInfo, err := GetExtendedUserRelmInfo(domain)
	if err != nil {
		return err
	}

	additionalDomains, err := GetAdditionalDomains(domain)
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

	utils.TablePrintOutputInfo(outputInfo)

	return nil
}

func GetUserRelmInfo(domain string) (utils.UserRelmInfo, error) {
	userRelmInfo := utils.UserRelmInfo{}
	request, err := http.Get(fmt.Sprintf(utils.AzureUserRelm, domain))
	if err != nil {
		return utils.UserRelmInfo{}, err
	}
	defer request.Body.Close()
	if request.StatusCode != 200 {
		return utils.UserRelmInfo{}, fmt.Errorf("Got status code %d while getting user realm info", request.StatusCode)
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return utils.UserRelmInfo{}, err
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

func GetExtendedUserRelmInfo(domain string) (utils.ExtendedUserRelmInfo, error) {
	extendedUserRelmInfo := utils.ExtendedUserRelmInfo{}
	request, err := http.Get(fmt.Sprintf(utils.AzureExtendedUserRelm, domain))
	if err != nil {
		return utils.ExtendedUserRelmInfo{}, err
	}
	defer request.Body.Close()
	if request.StatusCode != 200 {
		return utils.ExtendedUserRelmInfo{}, fmt.Errorf("Got status code %d while getting extended user realm info", request.StatusCode)
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return utils.ExtendedUserRelmInfo{}, err
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

func GetOIDInfo(domain string) (utils.OIDInfo, error) {
	oidInfo := utils.OIDInfo{}
	request, err := http.Get(fmt.Sprintf(utils.AzureOidInfo, domain))
	if err != nil {
		return utils.OIDInfo{}, err
	}
	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)
	if request.StatusCode == 400 {
		if err := json.Unmarshal(body, &oidInfo.ErrorOIDInfo); err != nil {
			return utils.OIDInfo{}, err
		}
		return utils.OIDInfo{}, fmt.Errorf("Error: %s", oidInfo.ErrorOIDInfo.Error)
	}
	if request.StatusCode != 200 {
		return utils.OIDInfo{}, fmt.Errorf("Request failed with status code %d", request.StatusCode)
	}
	if err := json.Unmarshal(body, &oidInfo.SuccessOIDInfo); err != nil {
		return utils.OIDInfo{}, err
	}
	return oidInfo, nil
}

func GetAdditionalDomains(domain string) (utils.FederationInfoResponse, error) {
	soapBody := fmt.Sprintf(utils.AzureFederationInfoRequest, domain)
	req, err := http.NewRequest("POST", utils.AzureDomainEmum, bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		return utils.FederationInfoResponse{}, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://schemas.microsoft.com/exchange/2010/Autodiscover/Autodiscover/GetFederationInformation")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return utils.FederationInfoResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.FederationInfoResponse{}, err
	}

	var federationInfo utils.FederationInfoResponse
	if err := xml.Unmarshal(body, &federationInfo); err != nil {
		return utils.FederationInfoResponse{}, fmt.Errorf("failed to unmarshal federation info response: %v", err)
	}

	return federationInfo, nil
}
