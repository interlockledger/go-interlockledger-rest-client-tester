// BSD 3-Clause License
//
// Copyright (c) 2022, InterlockLedger
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package json

import (
	"fmt"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/flags"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/core"
	"github.com/spf13/cobra"
)

var flagChainKeyRefs *[]string

// Implements POST /jsonDocuments@{chain}/withChainKeys
var jsonAddWithChainKeysCmd = &cobra.Command{
	Use:   "add-with-chain-keys",
	Short: "Adds a JSON into the given chain using the specified key.",
	Long: `Adds a JSON into the given chain using the specified key.
	
Calls POST /jsonDocuments@{chain}/withChainKeys`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		if len(*flagChainKeyRefs) == 0 {
			return fmt.Errorf("At least one chain key id must be provided.")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Key references:")
		fmt.Println("=================")
		for _, s := range *flagPubKeyRefs {
			fmt.Println(s)
		}
		jsonDoc, err := loadJSON()
		if err != nil {
			return fmt.Errorf("Unable to load the JSON document: %w\n", err)
		}
		fmt.Println("JSON to be added:")
		fmt.Println("=================")
		core.PrintAsJSON(jsonDoc)

		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return fmt.Errorf("Unable to initialize the client: %w\n", err)
		}
		ret, _, err := apiClient.JsonDocumentApi.JsonDocumentsAddWithChainKeys(nil, flags.Flags.ChainId, *flagChainKeyRefs, jsonDoc)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		fmt.Println()
		fmt.Println("Result:")
		fmt.Println("=======")
		core.PrintAsJSON(ret)
		return nil
	},
}

func init() {
	flagChainKeyRefs = jsonAddWithChainKeysCmd.Flags().StringArray("chain-key-id", []string{}, "A chain key id.")
}