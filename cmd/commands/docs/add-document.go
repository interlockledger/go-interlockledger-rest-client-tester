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

package docs

import (
	"bytes"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/core"
	"github.com/interlockledger/go-interlockledger-rest-client/client"
)

// Implements POST /documents/transaction/{transactionId}
var docsAddDocumentCmd = &cobra.Command{
	Use:   "add-document",
	Short: "Adds a document to the specified transaction.",
	Long: `Adds a document to the specified transaction.

Calls POST /documents/transaction/{transactionId}
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := docsFlags.RequireTransactionId(); err != nil {
			return err
		}
		if err := docsFlags.RequireDocument(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return err
		}

		// Load the parameters
		name := path.Base(docsFlags.DocumentFile)
		contents, err := os.ReadFile(docsFlags.DocumentFile)
		if err != nil {
			return err
		}
		var options client.DocumentsApiDocumentsAddDocumentParams
		options.Comment = docsFlags.Comment
		options.Path = docsFlags.Path
		options.Name = name
		options.ContentType = docsFlags.ContentType
		options.Contents = bytes.NewReader(contents)
		ret, _, err := apiClient.DocumentsApi.DocumentsAddDocument(nil,
			docsFlags.TransactionId,
			&options)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		core.PrintAsJSON(ret)
		return nil
	},
}

func init() {
	docsFlags.RegisterTransactionIDParameter(docsAddDocumentCmd.Flags())
	docsFlags.RegisterDocumentParameter(docsAddDocumentCmd.Flags())
}
