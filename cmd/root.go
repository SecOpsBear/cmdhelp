/*
Copyright © 2023 Surya nanda K surya@secopsbear.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/secopsbear/cmdhelp/colors"
	"github.com/secopsbear/cmdhelp/data"
	"github.com/spf13/cobra"
)

var Conn *sql.DB
var Color colors.Color

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmdhelp",
	Short: "Use to store and retrieve example commands for various tools",
	Long:  `Use to store and retrieve example commands for various tools`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		Color = *colors.GetOSColors()

		currentUser, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}
		dbFilePath := filepath.Join(currentUser.HomeDir, ".cmdhelp", "cmdhelp.db")
		if cmd.Name() != "init" {

			if _, err := os.Stat(dbFilePath); err != nil {
				fmt.Println(Color.RedColor(fmt.Sprintf("%s database file doesnot exist", dbFilePath)))
				fmt.Println(Color.RedColor("Use ") + Color.CyanColor("init") + Color.RedColor(" subcommand to setup and initialize database"))
				cmd.Help()
				os.Exit(0)

			} else {
				Conn, err = data.OpenDatabaseConn(dbFilePath)
				if err != nil {
					log.Panic(err)
				}
			}

			_, err = data.CheckTableExist()
			if err != nil {
				fmt.Println(Color.RedColor(err.Error()))
				fmt.Println(Color.RedColor("Use ") + Color.WhiteColor("init") + Color.RedColor(" subcommand to setup and initialize database"))
				cmd.Help()
				os.Exit(0)
			}
		}

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		defer Conn.Close()
	},
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
