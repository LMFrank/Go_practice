package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

var qs = []*survey.Question{
	{
		Name:     "id",
		Prompt:   &survey.Input{Message: "Access Secret ID: "},
		Validate: survey.Required,
		// Transform: survey.Title,
	},
	{
		Name:     "key",
		Prompt:   &survey.Password{Message: "Access Secret Key:"},
		Validate: survey.Required,
	},
	{
		Name: "region",
		Prompt: &survey.Select{
			Message: "Choose a region",
			Options: []string{"cn-hangzhou", "cn-shanghai"},
			Default: "cn-hangzhou",
		},
	},
	{
		Name: "Language",
		Prompt: &survey.MultiSelect{
			Message: "Supported Configure Language: ",
			Options: []string{"zh", "en", "jp"},
		},
	},
}

func interactive(profile string) {
	answers := struct {
		ID          string
		Key         string
		ChinaRegion string `survey:"region"`
		Language    []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if !confirm() {
		fmt.Println("用户取消文件保存.")
	}

	dumpConfig(profile, answers)
}

func dumpConfig(profile string, answers any) {
	b, err := json.MarshalIndent(answers, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	name := fmt.Sprintf("%s.config.json", profile)
	err2 := os.WriteFile(name, b, os.ModePerm)
	if err != nil {
		fmt.Println(err2)
	}
}
func confirm() bool {
	ok := false
	prompt := &survey.Confirm{
		Message: "是否保存文件？",
	}
	_ = survey.AskOne(prompt, &ok)
	return ok
}
