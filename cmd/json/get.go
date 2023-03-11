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
	"github.com/interlockledger/go-interlockledger-rest-client/client/jsondocs"
	"github.com/interlockledger/go-interlockledger-rest-client/crypto"
	"github.com/spf13/cobra"
)

var jsonGetCmdFlags = struct {
	privateKey string
	id         int64
}{}

// Implements GET /jsonDocuments@{chain}/{serial}
var jsonGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a JSON into the given chain",
	Long: `Get a JSON into the given chain. It requires the chain and id.
	
Call GET /jsonDocuments@{chain}/{serial}`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		if jsonGetCmdFlags.id == -1 {
			return fmt.Errorf("id is required.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return err
		}
		ret, _, err := apiClient.JsonDocumentApi.JsonDocumentsGet(nil, flags.Flags.ChainId, jsonGetCmdFlags.id)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		fmt.Println()
		fmt.Println("Encrypted JSON")
		fmt.Println("==============")
		core.PrintAsJSON(ret)
		if jsonGetCmdFlags.privateKey != "" {
			privKey, err := crypto.LoadPrivateKey(jsonGetCmdFlags.privateKey)
			if err != nil {
				return err
			}
			readerKey, err := crypto.NewReaderKeyFromPrivateKey(privKey)
			if err != nil {
				return err
			}
			dec, err := jsondocs.DecipherJSON(readerKey, &ret)
			if err != nil {
				return err
			}
			fmt.Println()
			fmt.Println("Decrypted JSON")
			fmt.Println("==============")
			fmt.Println(dec)
		}
		return nil
	},
}

func init() {
	jsonGetCmd.Flags().Int64VarP(&jsonGetCmdFlags.id, "id", "i", int64(-1), "The ID of the document. It may be required by some commands.")
	jsonGetCmd.Flags().StringVar(&jsonGetCmdFlags.privateKey, "private", "", "The reader's private key file. If set, the JSON document will be decrypted.")
}
