/*
Copyright © 2022 Val Gridnev

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DrGolem/homeDDNS/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List DNS records for Cloudflare registered zone",
	Long:  "List DNS records for Cloudflare registered zone",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := Cfg.ApiToken
		if Cfg.ApiToken == "" {
			log.Fatalln("API Token not set")
		}
		zoneId, err := utils.GetZoneId(Cfg.ZoneName, apiToken)
		if err != nil {
			log.Fatalln(err)
		}
		lst, err := utils.GetDnsRecords(zoneId, apiToken)
		if err != nil {
			log.Fatalln(err)
		}
		for _, rec := range lst {
			fmt.Printf("id: %s, name: %s, IP: %s\n", rec.Id, rec.Name, rec.IP)
		}
	},
}

func init() {
	dnsCmd.AddCommand(listCmd)
}
