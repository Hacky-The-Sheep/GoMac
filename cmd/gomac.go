package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "mac",
	Short: "Convert to proper mac address formats",
	Long: `Convert to proper mac address format with options.
		For Example:
		
		gomac convert -2d (dash every 2 chars) xxxxxxxxxxxxxxxx`,

	Run: generateMac,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("dash", "d", false, "dash every 2 chars")
	generateCmd.Flags().BoolP("period", "p", false, "dot every 2 chars")
	generateCmd.Flags().BoolP("colon", "c", false, "colon every 2 chars")
	generateCmd.Flags().BoolP("none", "n", false, "none every 2 chars")
	generateCmd.Flags().BoolP("all", "a", false, "all the formats")
}

func generateMac(cmd *cobra.Command, args []string) {
	dash, _ := cmd.Flags().GetBool("dash")
	period, _ := cmd.Flags().GetBool("period")
	colon, _ := cmd.Flags().GetBool("colon")
	none, _ := cmd.Flags().GetBool("none")
	all, _ := cmd.Flags().GetBool("all")

	original := args[0]
	var dashMac = ""
	var periodMac = ""
	var colonMac = ""
	var noneMac = ""
	count := 0

	for i := 0; i < len(original); i++ {
		if string(original[i]) == ":" || string(original[i]) == "-" {
			continue
		}
		if count == 2 {
			periodMac += "."
			dashMac += "-"
			colonMac += ":"
			count = 0
		}
		count += 1
		dashMac += string(original[i])
		periodMac += string(original[i])
		colonMac += string(original[i])
		noneMac += string(original[i])
	}
	if dash {
		fmt.Println(dashMac)
	}
	if period {
		fmt.Println(periodMac)
	}
	if colon {
		fmt.Println(colonMac)
	}
	if none {
		fmt.Println(noneMac)
	}
	if all {
		fmt.Println(dashMac)
		fmt.Println(periodMac)
		fmt.Println(colonMac)
		fmt.Println(noneMac)
	}
}
