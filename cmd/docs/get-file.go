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
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/core"
)

// Implements GET /documents/{locator}/{index}
var docsGetFileCmd = &cobra.Command{
	Use:   "get-file",
	Short: "Downloads the specified file inside the documents pointed by the locator.",
	Long: `Downloads the specified file inside the documents pointed by the locator.
	
Calls GET /documents/{locator}/{index}
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := docsFlags.RequireLocatorAndIndex(); err != nil {
			return err
		}
		if err := docsFlags.RequireOutput(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return err
		}

		resp, err := apiClient.DocumentsApi.DocumentsGetSingleDocument(nil,
			docsFlags.Locator, docsFlags.Index)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		defer resp.Body.Close()

		// Open the output file
		fmt.Printf("Writing the file %s...\n", docsFlags.OutputFile)
		writer, err := os.OpenFile(docsFlags.OutputFile,
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer writer.Close()

		// Download it...
		n, err := io.Copy(writer, resp.Body)
		if err != nil {
			return err
		}
		fmt.Printf("File file with %d bytes written.\n", n)
		return nil
	},
}

func init() {
	docsFlags.RegisterLocatorParameter(docsGetFileCmd.Flags())
	docsFlags.RegisterIndexParameter(docsGetFileCmd.Flags())
	docsFlags.RegisterZipFileParameter(docsGetFileCmd.Flags())
}
