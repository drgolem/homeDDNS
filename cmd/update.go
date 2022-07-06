/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Fatalln(err)
		}

		zoneId := Cfg.ZoneId
		apiToken := Cfg.ApiToken
		if Cfg.ApiToken == "" {
			log.Fatalln("API Token not set")
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
			log.Printf("DNS: %v", lst)
			log.Printf("IP: %s", publicIP)
		}

		for _, rec := range lst {
			if rec.IP != publicIP {
				if dryRun {
					log.Printf("Updating [%s] IP %s -> %s", rec.Name, rec.IP, publicIP)
				} else {
					err := UpdateDnsIP(zoneId, rec.ID, publicIP, apiToken)
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
