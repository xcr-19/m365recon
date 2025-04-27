package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func TablePrintOutputInfo(info OutputInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Field", "Value"})

	table.Append([]string{"Domain", info.Domain})
	table.Append([]string{"Tenant ID", info.TenantID})
	table.Append([]string{"Tenant Region", info.TenantRegion})
	table.Append([]string{"Federation Brand Name", info.FederationBrandName})
	table.Append([]string{"Auth URL", info.AuthURL})
	table.Append([]string{"Userinfo Endpoint", info.UserinfoEndpoint})
	table.Append([]string{"Kerberos Endpoint", info.KerberosEndpoint})
	table.Append([]string{"Token Endpoint", info.TokenEndpoint})
	table.Append([]string{"Is Dsso Enabled", fmt.Sprintf("%v", info.IsDssoEnabled)})
	table.Append([]string{"Force Login Hint", fmt.Sprintf("%v", info.ForceLoginHint)})
	table.Append([]string{"Additional Domains"})
	for _, domain := range info.AdditionalDomains {
		table.Append([]string{"", domain})
	}
	table.Render()
}
