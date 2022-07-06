/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/DrGolem/homeDDNS/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		zoneId := Cfg.ZoneId
		apiToken := Cfg.ApiToken
		if Cfg.ApiToken == "" {
			log.Fatalln("API Token not set")
		}
		lst, err := utils.GetDnsRecords(zoneId, apiToken)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("DNS: %v", lst)
	},
}

func init() {
	dnsCmd.AddCommand(listCmd)
}
