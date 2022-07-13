/*
Copyright © 2022 Val Gridnev

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/DrGolem/homeDDNS/utils"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates DNS records to set current IP address",
	Long:  "Updates DNS records to set current IP address",
	Run: func(cmd *cobra.Command, args []string) {

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Fatalln(err)
		}

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
		publicIP, err := utils.GetPublicIP()
		if err != nil {
			log.Fatalln(err)
		}
		if dryRun {
			fmt.Printf("IP: %s\n", publicIP)
			for _, rec := range lst {
				fmt.Printf("id: %s, name: %s, IP: %s\n", rec.Id, rec.Name, rec.IP)
			}
		}

		for _, rec := range lst {
			if rec.IP != publicIP {
				if dryRun {
					log.Printf("Updating [%s] IP %s -> %s", rec.Name, rec.IP, publicIP)
				} else {
					err := UpdateDnsIP(zoneId, rec.Id, publicIP, apiToken)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		}
	},
}

func UpdateDnsIP(zoneId string, recordId string, IPAddress string, apiToken string) error {
	updUrl := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		zoneId, recordId)

	client := http.Client{}

	updateCmd := fmt.Sprintf(`{"type":"A","content":"%s"}`, IPAddress)
	data := strings.NewReader(updateCmd)

	req, err := http.NewRequest("PATCH", updUrl, data)
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", apiToken)},
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(body))

	return nil
}

func init() {
	dnsCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolP("dry-run", "", false, "A dry run, no pdate")
}
