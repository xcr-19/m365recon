package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xcr-19/m365recon/pkg"
)

const banner = `                                
       ___ ___ ___                     
 _____|_  |  _|  _|___ ___ ___ ___ ___ 
|     |_  | . |_  |  _| -_|  _| . |   |
|_|_|_|___|___|___|_| |___|___|___|_|_|
                                       
`

var rootCmd = &cobra.Command{
	Use:   "m365recon",
	Short: banner + "\n" + "Microsoft Recon Tool by xcr-19",
	Long:  banner + "\n" + "Microsoft Recon Tool by xcr-19",
	Run: func(cmd *cobra.Command, args []string) {
		domain, _ := cmd.Flags().GetString("domain")
		if domain == "" {
			fmt.Println("Error: domain is not provided")
			cmd.Usage()
			os.Exit(1)
		}

		err := pkg.ReconByDomain(domain)
		if err != nil {
			fmt.Println("Error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("domain", "d", "", "Domain name")
	rootCmd.MarkFlagRequired("domain")
}
