/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/secopsbear/cmdhelp/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a cmdhelp database and table",
	Long:  `Initialise a cmdhelp database and table`,
	Run: func(cmd *cobra.Command, args []string) {

		currentUser, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}
		folderPath := fmt.Sprintf("%s/.cmdhelp", currentUser.HomeDir)
		_, err = os.Stat(folderPath)
		if errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(folderPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err = os.Stat(folderPath + "/cmdhelp.db")
			if !errors.Is(err, os.ErrNotExist) {
				fmt.Print("db file exists, Do you want to overwrite [" + Color.RedColor("y") + "] : ")
				t := ""
				fmt.Scanf("%s", &t)
				if strings.ToLower(t) != "y" {
					fmt.Println("Did not overwrite the database")
					// cmd.Help()
					os.Exit(0)
				} else {
					log.Println(Color.RedColor("Overwriting the database!"))
				}
			}
		}
		file, err := os.Create(folderPath + "/cmdhelp.db")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		Conn, err = data.OpenDatabaseConn(folderPath + "/cmdhelp.db")
		if err != nil {
			log.Panic(err)
		}
		data.CreateTable()
	},
	//	PostRun: func(cmd *cobra.Command, args []string) {
	//		defer Conn.Close()
	//	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
