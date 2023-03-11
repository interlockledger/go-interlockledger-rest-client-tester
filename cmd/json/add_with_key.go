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

// Implements POST /jsonDocuments@{chain}/withKey
var jsonAddWithKeyCmd = &cobra.Command{
	Use:   "add-with-key",
	Short: "Adds a JSON into the given chain using the specified key.",
	Long: `Adds a JSON into the given chain using the specified key.

Calls POST /jsonDocuments@{chain}/withKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		if JSONRootCmdFlags.ReaderCertFile == "" {
			return fmt.Errorf("Reader certificate is required.")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := LoadReaderCertificate(JSONRootCmdFlags.ReaderCertFile)
		if err != nil {
			return err
		}
		encKey, keyId, err := key.EncodedPublicKey()
		if err != nil {
			return err
		}
		fmt.Println("ReaderKey:")
		fmt.Println("=================")
		fmt.Println(encKey)
		fmt.Println("KeyID:")
		fmt.Println("=================")
		fmt.Println(keyId)
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
		ret, _, err := apiClient.JsonDocumentApi.JsonDocumentsAddWithKey(nil, flags.Flags.ChainId,
			encKey, keyId,
			jsonDoc)
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
