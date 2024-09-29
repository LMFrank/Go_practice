package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var root = &cobra.Command{
	Use:   "greeting",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		greeting(name, age)
	},
}

var (
	name = ""
	age  = 0
)

func init() {
	root.Flags().StringVarP(&name, "name", "n", "", "姓名")
	root.Flags().IntVarP(&age, "age", "a", 20, "年龄")
}

func greeting(name string, age int) {
	fmt.Printf("%s 你好，今年 %d 岁\n", name, age)
}

func main() {
	err := root.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
