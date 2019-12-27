/*
Copyright © 2019 Nik Ogura <nik@orionlabs.io>

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
	"encoding/json"
	"fmt"
	"github.com/onbeep/genkeyset/pkg/genkeyset"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var numKeys int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "genkeyset",
	Short: "A tool to generate JWK Key Sets for self-hosted Orionlabs PTT.",
	Long: `
A tool to generate JWK Key Sets for self-hosted Orionlabs PTT.

The JWK Key Set is described in https://tools.ietf.org/html/rfc7517.  
The RFC provides for additional members in the key per “Additional members” - 
https://tools.ietf.org/html/rfc7517#section-4.  Orion labs uses this to add the 'live' boolean attribute.

Orionlabs PTT requires a minimum of 3 active keys at all times, hence the default usage of this tool is to 
create 3 keys in the KeySet.  This value can be overwritten by the user if desired.

`,
	Run: func(cmd *cobra.Command, args []string) {
		keyset, err := genkeyset.GenerateKeySet(numKeys)
		if err != nil {
			log.Fatalf("Failed to generate keyset: %s", err)
		}

		jsonbuf, err := json.MarshalIndent(keyset, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshall JWK KeySet into json: %s", err)
		}

		fmt.Printf("%s\n", jsonbuf)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&numKeys, "keys", "k", 3, "Number of keys to generate in KeySet.")
}
