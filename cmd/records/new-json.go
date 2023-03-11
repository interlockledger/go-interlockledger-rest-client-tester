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

package records

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/interlockledger/go-interlockledger-rest-client-tester/cmd/flags"
	"github.com/interlockledger/go-interlockledger-rest-client-tester/core"
	"github.com/interlockledger/go-interlockledger-rest-client/client"
)

var recordNewJSONCmdFlags = struct {
	appId        int64
	payloadTagId int64
	recType      string
	payloadFile  string
}{}

// Implements POST /records@{chain}/asJson
var recordNewJSONCmd = &cobra.Command{
	Use:   "new-json",
	Short: "Creates a new record from a JSON.",
	Long: `Creates a new record. Use a param file like record-new.json to set the new chain parameters.

Calls POST /records@{chain}/asJson`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := flags.Flags.RequireChainId(); err != nil {
			return err
		}
		if recordNewJSONCmdFlags.payloadFile == "" {
			return fmt.Errorf("Payload file missing.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := core.AppCore.NewClient()
		if err != nil {
			return err
		}

		var payload map[string]any
		if err := core.LoadJSONFile(recordNewJSONCmdFlags.payloadFile, &payload); err != nil {
			return err
		}

		var options client.RecordApiRecordAddAsJsonOpts
		options.ApplicationId = flags.OptionalInt64(recordNewJSONCmdFlags.appId)
		options.PayloadTagId = flags.OptionalInt64(recordNewJSONCmdFlags.payloadTagId)
		options.Type_ = flags.OptionalString(recordNewJSONCmdFlags.recType)

		// Load the parameters
		ret, _, err := apiClient.RecordApi.RecordAddAsJson(nil, flags.Flags.ChainId, &options, &payload)
		if err != nil {
			return core.FormatRequestResponseCommandError(err)
		}
		core.PrintAsJSON(ret)
		return nil
	},
}

func init() {
	recordNewJSONCmd.Flags().Int64Var(&recordNewJSONCmdFlags.appId, "app-id", -1, "Application id.")
	recordNewJSONCmd.Flags().Int64Var(&recordNewJSONCmdFlags.payloadTagId, "payload-tag-id", -1, "Payload tag id.")
	recordNewJSONCmd.Flags().StringVar(&recordNewJSONCmdFlags.recType, "record-type", "", "Record type. Must be Root, Data or Closing.")
	recordNewJSONCmd.Flags().StringVar(&recordNewJSONCmdFlags.payloadFile, "payload", "", "Name of the payload file. It must be a valid JSON.")
}
