package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var config string

var root = &cobra.Command{
	Use:   "greeting",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		person := readConfig(config)
		greeting(person.Name, person.Age)
		dumpConfig(person)
	},
}

func init() {
	root.Flags().StringVarP(&config, "config", "c", "config.yml", "配置文件")
}

type Person struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Age  int    `yaml:"age,omitempty" json:"age,omitempty"`
}

func readConfig(config string) *Person {
	person := &Person{}
	b, err := os.ReadFile(config)
	if err != nil {
		fmt.Println(err)
	}

	err2 := yaml.Unmarshal(b, person)
	if err2 != nil {
		fmt.Println(err2)
	}

	return person
}

func dumpConfig(person *Person) {
	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	err2 := os.WriteFile("config.json", b, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
}
func greeting(name string, age int) {
	fmt.Printf("%s 你好，今年 %d 岁\n", name, age)
}

func main() {
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
