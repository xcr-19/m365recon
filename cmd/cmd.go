package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xcr-19/m365recon/pkg"
	"github.com/xcr-19/m365recon/utils"
)

const banner = `                                
       ___ ___ ___                     
 _____|_  |  _|  _|___ ___ ___ ___ ___ 
|     |_  | . |_  |  _| -_|  _| . |   |
|_|_|_|___|___|___|_| |___|___|___|_|_|
                                       
`

var (
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "m365recon [recon] | [enum] | [brute]",
	Short:   banner + "\n" + "Microsoft Recon Tool by xcr-19",
	Long:    banner + "\n" + "Microsoft Recon Tool by xcr-19",
	Version: version,
	Example: "m365recon recon ... [commands]",
}

var reconCmd = &cobra.Command{
	Use: "recon [-d domain] [-o output] [-v verbose]",
	Example: `m365recon recon -d google.com -o output.json
m365recon recon -d google.com`,
	Short: "Recon module",
	Long:  banner + "\n" + "Recon module for Microsoft 365",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.Config{}
		domain, _ := cmd.Flags().GetString("domain")
		verbose, _ := cmd.Flags().GetBool("verbose")
		output, _ := cmd.Flags().GetString("output")
		if domain == "" {
			fmt.Println("Error: domain is not provided")
			cmd.Usage()
			os.Exit(1)
		}
		config.Verbose = verbose
		config.Output = output
		err := pkg.ReconByDomain(domain, config)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(reconCmd)
	reconCmd.Flags().StringP("domain", "d", "", "Domain name")
	reconCmd.MarkFlagRequired("domain")
	reconCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	reconCmd.Flags().StringP("output", "o", "", "Output file")
}
