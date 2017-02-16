// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// START OMIT
var user string     // HL
var password string // HL

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Grant this app access to Moss",
	Long: `This command will allow the app to access Moss on behalf of you.
	
You must run this command to setup the app before executing any other
command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Trying to log user", user, "with password", password) // HL
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVar(&user, "user", "", "The username to login with") // HL
	loginCmd.Flags().StringVar(&password, "password", "", "The password")       // HL
}

// END OMIT
