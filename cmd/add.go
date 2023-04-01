/*
Copyright © 2023 Surya nanda K
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/secopsbear/cmdhelp/data"
	"github.com/secopsbear/cmdhelp/helpui"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Enter a new command example for a COMMAND",
	Long:  `Enter a new command example for a COMMAND`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewToolCommand()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// createNewToolCommand creates a prompt for entering tool, its gist and the command
func createNewToolCommand() {
	mainToolPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a tool",
		Lable:    "Select a TOOL or add a new tool",
	}
	tool := strings.ToLower(helpui.PromptGetSelect(mainToolPromptContent))

	toolExplainPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a gist",
		Lable:    fmt.Sprintf("Short gist about %s command : ", strings.ToUpper(tool)),
	}

	toolExplainData := helpui.PromptGetInput(toolExplainPromptContent)

	commandPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a Command",
		Lable:    fmt.Sprintf("Enter the  %s command: ", strings.ToUpper(tool)),
	}
	commandData := helpui.PromptGetInput(commandPromptContent)

	data.InsertCommand(tool, toolExplainData, commandData)
}
