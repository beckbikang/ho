package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var mcCmd = &cobra.Command{
	Use:   "mc",
	Short: "运行服务",
	Long:  "运行服务",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mcKafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mc to kafka")
	},
}

func init() {
	mcCmd.AddCommand(mcKafkaCmd)
}
