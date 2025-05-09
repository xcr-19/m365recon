package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xcr-19/m365recon/pkg"
	"github.com/xcr-19/m365recon/utils"
)

var reconCmd = &cobra.Command{
	Use: "recon [-d domain] [-o output] [-v verbose]",
	Example: `m365recon recon -d google.com -o output.json
m365recon recon -d google.com`,
	Short: "Recon module",
	Long:  banner + "\n" + "Recon module for Microsoft 365",
	Run: func(cmd *cobra.Command, args []string) {
		config := setupConfig(cmd)
		err := pkg.ReconByDomain(config.Domain, config)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func setupConfig(cmd *cobra.Command) utils.Config {
	config := utils.Config{}
	config.Domain, _ = cmd.Flags().GetString("domain")
	config.Verbose, _ = cmd.Flags().GetBool("verbose")
	config.Output, _ = cmd.Flags().GetString("output")
	return config
}
