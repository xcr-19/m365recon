package cmd

import (
	"github.com/spf13/cobra"
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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(reconCmd)
	reconCmd.Flags().StringP("domain", "d", "", "Domain name")
	reconCmd.MarkFlagRequired("domain")
	reconCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	reconCmd.Flags().StringP("output", "o", "", "Output file")
	reconCmd.Flags().StringP("proxy", "p", "", "Proxy")
}
