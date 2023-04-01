/*
Copyright © 2023 Surya nanda K
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/secopsbear/cmdhelp/data"
	"github.com/secopsbear/cmdhelp/helpui"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the command from the commands list",
	Long:  `Delete the command from the commands list`,
	Run: func(cmd *cobra.Command, args []string) {
		toolsList := data.ExtractToolsList()
		if len(toolsList) == 0 {
			fmt.Println(Color.RedColor("There is no data!"))
			os.Exit(0)
		}
		deleteToolCommandTK(toolsList)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func promptGetSelectDeleteCommand(pc data.PromptContent, tool string) {

	itemsToolsData := data.DisplaySelectedCommand(tool)
	index := -1
	var result string
	var err error
	var items []string
	var itemsId []int
	for _, datas := range *itemsToolsData {
		items = append(items, datas.ToolCommand)
		itemsId = append(itemsId, datas.IdCommand)
	}

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
	fmt.Printf("Deleting %s from %s ", Color.YellowColor(result), Color.CyanColor(tool))
	data.DeleteCommand(itemsId[index])
}

func deleteToolCommandTK(toolsList []string) {
	wordPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a tool",
		Lable:    "Select a TOOL : ",
	}
	tool := helpui.PromptGetSelectFromExistingList(wordPromptContent, toolsList)

	toolExplainPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a delete command",
		Lable:    fmt.Sprintf("Select the command to Delete of %s commands : ", strings.ToUpper(tool)),
	}
	promptGetSelectDeleteCommand(toolExplainPromptContent, tool)
}
