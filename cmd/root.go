/*
Copyright © 2022 Val Gridnev

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ZoneName string `mapstructure:"ZONE_NAME"`
	ApiToken string `mapstructure:"API_TOKEN"`
}

var cfgFile string

var Cfg Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "homeDDNS",
	Short: "A tool to change IP addresses of cloudflare DNS records",
	Long:  `A tool to change IP addresses of cloudflare DNS records`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.homeDDNS)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".homeDDNS" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("env")
		viper.SetConfigName(".homeDDNS")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.BindEnv("ZONE_NAME")
	viper.BindEnv("API_TOKEN")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		fmt.Fprintln(os.Stderr, "Read config file:", err)
	}
}
