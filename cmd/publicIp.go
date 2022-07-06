/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publicIp called")

		publicIP, err := utils.GetPublicIP()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("IP: %s", publicIP)
	},
}

func init() {
	rootCmd.AddCommand(publicIpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publicIpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publicIpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
