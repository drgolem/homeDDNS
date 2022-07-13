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

// publicIpCmd represents the publicIp command
var publicIpCmd = &cobra.Command{
	Use:   "publicIp",
	Short: "Prints public IP address",
	Long:  "Prints public IP address",
	Run: func(cmd *cobra.Command, args []string) {
		publicIP, err := utils.GetPublicIP()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s\n", publicIP)
	},
}

func init() {
	rootCmd.AddCommand(publicIpCmd)
}
