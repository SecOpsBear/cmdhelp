/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/secopsbear/cmdhelp/data"
	"github.com/secopsbear/cmdhelp/helpui"
)

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "lookup commands for a tool",
	Long:  `lookup commands for a tool`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) < 3 {
			toolsList := data.ExtractToolsList()
			if len(toolsList) == 0 {
				fmt.Println(Color.RedColor("There is no data!"))
				os.Exit(0)
			}
			lookupToolCommand(toolsList)
		} else {
			findMe := os.Args[2]
			s := data.SearchForTool(findMe)
			if len(s) > 0 {
				for _, v := range s {
					fmt.Println("["+Color.CyanColor(v.Tool)+"] "+strings.TrimSpace(v.Gist), ": \n  "+Color.YellowColor(v.ToolCommand))
				}
			} else {
				d := data.SearchForToolLike(findMe)
				if len(d) == 0 {
					fmt.Println(Color.YellowColor("There is no data found for:"), Color.RedColor(findMe))
					os.Exit(0)
				}
				var toolsList []string
				for _, v := range d {
					toolsList = append(toolsList, v.Tool)
				}
				toolsList = append(toolsList, "** exit **")
				t := data.Unique(toolsList)
				lookupToolCommandFrom(t)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
}

func lookupToolCommand(toolsList []string) {
	mainToolPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a tool",
		Lable:    "Select a TOOL ",
	}
	tool := helpui.PromptGetSelectFromExistingList(mainToolPromptContent, toolsList)

	s := data.SearchForTool(tool)
	if len(s) > 0 {
		for _, v := range s {
			fmt.Println("["+Color.CyanColor(v.Tool)+"] "+strings.TrimSpace(v.Gist), ": \n  "+Color.YellowColor(v.ToolCommand))
		}
	} else {
		fmt.Println(Color.YellowColor("There is no data found for:"), Color.RedColor(tool))
	}

}

func lookupToolCommandFrom(toolsLike []string) {
	mainToolPromptContent := data.PromptContent{
		ErrorMsg: "Please provide a tool",
		Lable:    "Did you mean ",
	}
	tool := helpui.PromptGetSelectFromLike(mainToolPromptContent, toolsLike)
	if tool == "** exit **" {
		os.Exit(0)
	}

	s := data.SearchForTool(tool)
	if len(s) > 0 {
		for _, v := range s {
			fmt.Println("["+Color.CyanColor(v.Tool)+"] "+strings.TrimSpace(v.Gist), ": \n  "+Color.YellowColor(v.ToolCommand))
		}
	} else {
		fmt.Println(Color.YellowColor("There is no data found for:"), Color.RedColor(tool))
	}

}
