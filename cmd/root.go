package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(mcCmd)
	rootCmd.AddCommand(mcServer)
	rootCmd.AddCommand(redisCmd)
	rootCmd.AddCommand(kafkaToKafkaCmd)
}
