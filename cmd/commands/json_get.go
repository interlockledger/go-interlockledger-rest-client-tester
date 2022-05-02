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

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var jsonGetCmdParam = struct {
	Chain string
	Id    int64
}{}

// testCmd represents the test command
var jsonGetCmd = &cobra.Command{
	Use:   "json_get",
	Short: "Adds a dummy JSON into the given chain",
	RunE:  runJsonGet,
}

func init() {
	jsonGetCmd.Flags().StringVarP(&jsonGetCmdParam.Chain, "chain", "c", "", "The ID of the chain.")
	jsonGetCmd.Flags().Int64VarP(&jsonGetCmdParam.Id, "id", "i", int64(-1), "The ID of the document.")
}

func runJsonGet(cmd *cobra.Command, args []string) error {

	if jsonGetCmdParam.Chain == "" {
		return fmt.Errorf("chain is required.")
	}
	if jsonGetCmdParam.Id == -1 {
		return fmt.Errorf("document id is required.")
	}

	client, err := createAPIClient()
	ret, _, err := client.JsonDocumentApi.JsonDocumentsGet(nil, jsonGetCmdParam.Chain, jsonGetCmdParam.Id)
	if err != nil {
		return fmt.Errorf("Unable to get the json document: %w\n", err)
	}
	printAsJSON(ret)
	return nil
}
