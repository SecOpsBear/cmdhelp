package helpui

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/secopsbear/cmdhelp/data"
)

func PromptGetSelect(pc data.PromptContent) string {

	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9-_]*$`).MatchString(input)

		if !is_alphanumeric {
			return errors.New("enter only alphanumeric,-,_ for command name")
		}
		return nil
	}

	items := data.ExtractToolsList()
	index := -1
	var result string
	var err error
	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.Lable,
			Items:    items,
			AddLabel: "New command",
			Validate: validate,
		}
		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
			index = 0
		}
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func PromptGetInput(pc data.PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.Lable,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v \n", err)
		os.Exit(1)
	}

	return result
}

func PromptGetSelectFromExistingList(pc data.PromptContent, items []string) string {

	index := -1
	var result string
	var err error

	for index < 0 {

		promptTK := promptui.Select{
			Label: pc.Lable,
			Items: items,
		}
		index, result, err = promptTK.Run()
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func PromptGetSelectFromLike(pc data.PromptContent, items []string) string {

	index := -1
	var result string
	var err error
	for index < 0 {
		prompt := promptui.Select{
			Label: pc.Lable,
			Items: items,
		}
		index, result, err = prompt.Run()

	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
