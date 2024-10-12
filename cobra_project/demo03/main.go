package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var profile string

func init() {
	root.Flags().StringVarP(&profile, "profile", "p", "default", "配置文件")
}

var root = &cobra.Command{
	Use:   "aliyun",
	Short: "aliyun 配置中心",
	Run: func(cmd *cobra.Command, args []string) {
		interactive(profile)
	},
}

func main() {
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
