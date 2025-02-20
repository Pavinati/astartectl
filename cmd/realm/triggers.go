// Copyright © 2019 Ispirata Srl
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

package realm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// triggersCmd represents the triggers command
var triggersCmd = &cobra.Command{
	Use:   "triggers",
	Short: "Manage triggers",
	Long:  `List, show, install or delete triggers in your realm.`,
}

var triggersListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List triggers",
	Long:    `List the name of triggers installed in the realm.`,
	Example: `  astartectl realm-management triggers list`,
	RunE:    triggersListF,
}

var triggersShowCmd = &cobra.Command{
	Use:     "show <trigger_name>",
	Short:   "Show trigger",
	Long:    `Shows a trigger installed in the realm.`,
	Example: `  astartectl realm-management triggers show my_data_trigger`,
	Args:    cobra.ExactArgs(1),
	RunE:    triggersShowF,
}

var triggersInstallCmd = &cobra.Command{
	Use:   "install <trigger_file>",
	Short: "Install trigger",
	Long: `Install the given trigger in the realm.
<trigger_file> must be a path to a JSON file containing a valid Astarte trigger.`,
	Example: `  astartectl realm-management triggers install my_data_trigger.json`,
	Args:    cobra.ExactArgs(1),
	RunE:    triggersInstallF,
}

var triggersDeleteCmd = &cobra.Command{
	Use:     "delete <trigger_name>",
	Short:   "Delete a trigger",
	Long:    `Deletes the specified trigger from the realm.`,
	Example: `  astartectl realm-management triggers delete my_data_trigger`,
	Args:    cobra.ExactArgs(1),
	RunE:    triggersDeleteF,
}

func init() {
	RealmManagementCmd.AddCommand(triggersCmd)

	triggersCmd.AddCommand(
		triggersListCmd,
		triggersShowCmd,
		triggersInstallCmd,
		triggersDeleteCmd,
	)
}

func triggersListF(command *cobra.Command, args []string) error {
	realmTriggers, err := astarteAPIClient.RealmManagement.ListTriggers(realm, realmManagementJwt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(realmTriggers)
	return nil
}

func triggersShowF(command *cobra.Command, args []string) error {
	triggerName := args[0]

	triggerDefinition, err := astarteAPIClient.RealmManagement.GetTrigger(realm, triggerName, realmManagementJwt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respJSON, _ := json.MarshalIndent(triggerDefinition, "", "  ")
	fmt.Println(string(respJSON))
	return nil
}

func triggersInstallF(command *cobra.Command, args []string) error {
	triggerFile, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}

	var triggerBody map[string]interface{}
	err = json.Unmarshal(triggerFile, &triggerBody)
	if err != nil {
		return err
	}

	err = astarteAPIClient.RealmManagement.InstallTrigger(realm, triggerBody, realmManagementJwt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("ok")
	return nil
}

func triggersDeleteF(command *cobra.Command, args []string) error {
	triggerName := args[0]
	err := astarteAPIClient.RealmManagement.DeleteTrigger(realm, triggerName, realmManagementJwt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("ok")
	return nil
}
