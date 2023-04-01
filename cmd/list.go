/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/secopsbear/cmdhelp/data"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list All the COMMANDS",
	Long:  `list All the COMMANDS`,
	Run: func(cmd *cobra.Command, args []string) {
		liststat := data.ExtractToolsList()
		if len(liststat) > 0 {
			fmt.Println("Tools : " + strings.Join(liststat, ","))
			td := data.DisplayAllNotes()
			for _, v := range td {
				fmt.Println("[" + Color.CyanColor(v.Tool) + "] " + Color.WhiteColor(v.Gist) + " : " + Color.YellowColor(v.ToolCommand))
			}
		} else {
			fmt.Println(Color.RedColor("There is no data!"))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
